package saldodetail

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ariefrpm/movies2/gen/go/proto/v1"
	"github.com/ariefrpm/movies2/pkg/library/router"
)

func restDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("i")

	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return request{kasId: i}, nil
}

func restEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	log.Print(response)
	return json.NewEncoder(w).Encode(response)
}

func restEncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func grpcEncodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(response)
	log.Print("SINI RESPONSE", res)

	return &proto.SaldoDetailResponse{
		TotalSaldo: res.TotalSaldo,
	}, nil
}

func grpcDecodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.SaldoDetailRequest)
	return request{kasId: req.KasId}, nil
}

type grpcTransport struct {
	proto.UnimplementedSaldoDetailServiceServer
	handler router.Handler
}

func (g *grpcTransport) SaldoDetail(ctx context.Context, request *proto.SaldoDetailRequest) (*proto.SaldoDetailResponse, error) {
	_, res, err := g.handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*proto.SaldoDetailResponse), nil
}
