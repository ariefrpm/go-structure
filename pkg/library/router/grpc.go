package router

import "context"

type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}

type grpcHandler struct {
	endpoint Endpoint
	decode DecodeGrpcRequestFunc
	encode EncodeGrpcResponseFunc
}

func NewGrpcHandler(endpoint Endpoint, decode DecodeGrpcRequestFunc, encode EncodeGrpcResponseFunc) *grpcHandler {
	return &grpcHandler{
		endpoint: endpoint,
		decode:   decode,
		encode:   encode,
	}
}

func (g *grpcHandler) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	request, err := g.decode(ctx, req)
	if err != nil {
		return ctx, nil, err
	}

	response, err := g.endpoint(ctx, request)
	if err != nil {
		return ctx, nil, err
	}

	grpcResp, err := g.encode(ctx, response)
	if err != nil {
		return ctx, nil, err
	}

	return ctx, grpcResp, nil
}

type DecodeGrpcRequestFunc func(context.Context, interface{}) (request interface{}, err error)
type EncodeGrpcResponseFunc func(context.Context, interface{}) (response interface{}, err error)
