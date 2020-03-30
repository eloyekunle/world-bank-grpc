package mock_worldbank_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	wbmock "github.com/eloyekunle/world-bank-grpc/pkg/mock_worldbank"
	"github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestGetCountry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWorldBankClient := wbmock.NewMockWorldBankClient(ctrl)
	req := &worldbank.CountryID{Id: "NGA"}

	mockWorldBankClient.EXPECT().GetCountry(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&worldbank.Country{Id: "NGA", Name: "Nigeria", Capital: "Abuja"}, nil)

	testGetCountry(t, mockWorldBankClient)
}

func testGetCountry(t *testing.T, client worldbank.WorldBankClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetCountry(ctx, &worldbank.CountryID{Id: "NGA"})
	if err != nil || r.Capital != "Abuja" {
		t.Errorf("mocking failed")
	}

	t.Log("Reply : ", r.Capital)
}
