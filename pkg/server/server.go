//go:generate protoc -I ../worldbank --go_out=plugins=grpc:../worldbank ../worldbank/world_bank.proto

package server

import (
	"net/http"
	"sort"
	"time"

	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"github.com/go-logr/logr"
	"github.com/jkkitakita/wbdata-go"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
)

type WorldBankServer struct {
	pb.UnimplementedWorldBankServer

	client *wbdata.Client

	countries    map[string]*wbdata.Country
	regions      map[string]*wbdata.Region
	incomeLevels map[string]*wbdata.IncomeLevel
	lendingTypes map[string]*wbdata.LendingType

	log logr.Logger
}

func NewServer() *WorldBankServer {
	s := &WorldBankServer{
		client:       wbdata.NewClient(&http.Client{Timeout: 30 * time.Second}),
		countries:    make(map[string]*wbdata.Country),
		regions:      make(map[string]*wbdata.Region),
		incomeLevels: make(map[string]*wbdata.IncomeLevel),
		lendingTypes: make(map[string]*wbdata.LendingType),
		log:          klogr.New().WithValues("component", "server"),
	}

	s.log.Info("loading countries")
	s.loadCountries()
	s.log.Info("loaded countries")

	return s
}

func (s *WorldBankServer) loadCountries() {
	_, countries, err := s.client.Countries.ListCountries(wbdata.PageParams{
		Page:    1,
		PerPage: 400, // loads everything at once
	})
	if err != nil {
		klog.Fatalf("failed to load countries: %v", err)
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Name < countries[j].Name
	})

	for _, country := range countries {
		if country.Region.ID != "NA" {
			s.countries[country.ID] = &country
		}

		if _, ok := s.regions[country.Region.ID]; !ok {
			s.regions[country.Region.ID] = &country.Region
		}

		if _, ok := s.incomeLevels[country.IncomeLevels.ID]; !ok {
			s.incomeLevels[country.IncomeLevels.ID] = &country.IncomeLevels
		}

		if _, ok := s.lendingTypes[country.LendingType.ID]; !ok {
			s.lendingTypes[country.LendingType.ID] = &country.LendingType
		}
	}
}
