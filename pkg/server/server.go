//go:generate protoc -I ../worldbank --go_out=plugins=grpc:../worldbank ../worldbank/world_bank.proto

package server

import (
	"context"
	"net/http"
	"sort"
	"time"

	"github.com/eloyekunle/world-bank-grpc/pkg/util"
	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"github.com/jkkitakita/wbdata-go"
	"k8s.io/klog/v2"
)

type WorldBankServer struct {
	pb.UnimplementedWorldBankServer

	client *wbdata.Client

	countries    map[string]*wbdata.Country
	regions      map[string]*wbdata.Region
	incomeLevels map[string]*wbdata.IncomeLevel
	lendingTypes map[string]*wbdata.LendingType
}

func NewServer() *WorldBankServer {
	s := &WorldBankServer{
		client:       wbdata.NewClient(&http.Client{Timeout: 30 * time.Second}),
		countries:    make(map[string]*wbdata.Country),
		regions:      make(map[string]*wbdata.Region),
		incomeLevels: make(map[string]*wbdata.IncomeLevel),
		lendingTypes: make(map[string]*wbdata.LendingType),
	}

	klog.Info("loading countries")
	s.loadCountries()
	klog.Info("loaded countries")

	return s
}

func (s *WorldBankServer) ListRegions(req *pb.Void, stream pb.WorldBank_ListRegionsServer) error {
	klog.Info("received request - ListRegions")
	for _, region := range s.regions {
		pbRegion := &pb.Region{
			Id:   region.ID,
			Name: region.Value,
		}

		if err := stream.Send(pbRegion); err != nil {
			return err
		}
	}

	return nil
}

func (s *WorldBankServer) ListIncomeLevels(req *pb.Void, stream pb.WorldBank_ListIncomeLevelsServer) error {
	klog.Info("received request - ListIncomeLevels")

	for _, incomeLevel := range s.incomeLevels {
		pbIncomeLevel := &pb.IncomeLevel{
			Id:   incomeLevel.ID,
			Name: incomeLevel.Value,
		}

		if err := stream.Send(pbIncomeLevel); err != nil {
			return err
		}
	}

	return nil
}

func (s *WorldBankServer) ListLendingTypes(req *pb.Void, stream pb.WorldBank_ListLendingTypesServer) error {
	klog.Info("received request - ListLendingTypes")

	for _, lendingType := range s.lendingTypes {
		pbLendingType := &pb.LendingType{
			Id:   lendingType.ID,
			Name: lendingType.Value,
		}

		if err := stream.Send(pbLendingType); err != nil {
			return err
		}
	}

	return nil
}

func (s *WorldBankServer) ListCountries(req *pb.CountryFilter, stream pb.WorldBank_ListCountriesServer) error {
	klog.Info("received request - ListCountries")

	var matches []*pb.Country

	for _, country := range s.countries {
		if (req.Region == util.All || req.Region == country.Region.ID) &&
			(req.LendingType == util.All || req.LendingType == country.LendingType.ID) &&
			(req.IncomeLevel == util.All || req.IncomeLevel == country.IncomeLevel.ID) {
			matches = append(matches, wbToPbCountry(*country))
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name < matches[j].Name
	})

	for _, country := range matches {
		if err := stream.Send(country); err != nil {
			return err
		}
	}

	return nil
}

func (s *WorldBankServer) GetCountry(ctx context.Context, req *pb.CountryID) (*pb.Country, error) {
	klog.Info("received request - GetCountry")

	country, ok := s.countries[req.Id]
	if !ok {
		return nil, nil
	}

	return wbToPbCountry(*country), nil
}

func (s *WorldBankServer) loadCountries() {
	_, countries, err := s.client.Countries.ListCountries(wbdata.PageParams{
		Page:    1,
		PerPage: 400, // loads everything at once
	})
	if err != nil {
		klog.Exitf("failed to load countries: %v", err)
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Name < countries[j].Name
	})

	for _, country := range countries {
		country := country

		if country.Region.ID == "NA" {
			continue
		}

		s.countries[country.ID] = &country

		if _, ok := s.regions[country.Region.ID]; !ok {
			s.regions[country.Region.ID] = &country.Region
		}

		if _, ok := s.incomeLevels[country.IncomeLevel.ID]; !ok {
			s.incomeLevels[country.IncomeLevel.ID] = &country.IncomeLevel
		}

		if _, ok := s.lendingTypes[country.LendingType.ID]; !ok {
			s.lendingTypes[country.LendingType.ID] = &country.LendingType
		}
	}
}

func wbToPbCountry(country wbdata.Country) *pb.Country {
	return &pb.Country{
		Id:          country.ID,
		Name:        country.Name,
		Capital:     country.CapitalCity,
		Region:      country.Region.Value,
		IncomeLevel: country.IncomeLevel.Value,
		LendingType: country.LendingType.Value,
	}
}
