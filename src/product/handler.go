package product

import (
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) PostProduct(ctx context.Context, in *ProductRequest) (*ProductReply, error) {
	return &ProductReply{Message: "OK"}, nil
}
