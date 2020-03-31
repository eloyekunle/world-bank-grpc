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

	req := &worldbank.CountryID{Id: "NGA"}

	mockWorldBankClient := wbmock.NewMockWorldBankClient(ctrl)
	mockWorldBankClient.EXPECT().GetCountry(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&worldbank.Country{Id: "NGA", Name: "Nigeria", Capital: "Abuja"}, nil)

	testGetCountry(t, mockWorldBankClient)
}

func testGetCountry(t *testing.T, client worldbank.WorldBankClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := client.GetCountry(ctx, &worldbank.CountryID{Id: "NGA"})
	if err != nil || r.Capital != "Abuja" {
		t.Fatalf("mocking failed")
	}

	t.Log("Reply:", r.Capital)
}

var msg = &worldbank.Region{
	Id:   "SSF",
	Name: "Sub-Saharan Africa",
}

func TestListRegions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := wbmock.NewMockWorldBank_ListRegionsClient(ctrl)
	stream.EXPECT().Recv().Return(msg, nil)

	mockWorldBankClient := wbmock.NewMockWorldBankClient(ctrl)
	mockWorldBankClient.EXPECT().ListRegions(
		gomock.Any(),
		gomock.Any(),
	).Return(stream, nil)

	if err := testListRegions(mockWorldBankClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testListRegions(client worldbank.WorldBankClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListRegions(ctx, &worldbank.Void{})
	if err != nil {
		return err
	}

	got, err := stream.Recv()
	if err != nil {
		return err
	}
	if !proto.Equal(got, msg) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, msg)
	}

	return nil
}
