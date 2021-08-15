package moviedetail

import (
	"context"
	"encoding/json"
	"github.com/ariefrpm/movies2/pkg/library/router"
	"github.com/ariefrpm/movies2/pkg/proto"
	"net/http"
)

func restDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("i")
	return request{ID: id}, nil
}

func restEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func restEncodeError(_ context.Context, err error, w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func grpcEncodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(response)
	return &proto.MovieDetailResponse{
		Title:    res.Title,
		Year:     res.Year,
		Rated:    res.Rated,
		Released: res.Released,
		Runtime:  res.Runtime,
		Genre:    res.Genre,
		Director: res.Director,
		Writer:   res.Writer,
		Actors:   res.Actors,
		Plot:     res.Plot,
		Language: res.Language,
		Country:  res.Country,
		Awards:   res.Awards,
		Poster:   res.Poster,
		Response: res.Response,
	}, nil
}

func grpcDecodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.MovieDetailRequest)
	return request{ID:req.OmdbID}, nil
}

type grpcTransport struct {
	handler router.Handler
}

func (g *grpcTransport)  MovieDetail(ctx context.Context, request *proto.MovieDetailRequest) (*proto.MovieDetailResponse, error) {
	_, res, err := g.handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*proto.MovieDetailResponse), nil
}