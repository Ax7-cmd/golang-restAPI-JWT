package Grpc

import (
	"context"
	"sync"

	pb "golang-restAPI-JWT/Core/Grpc/Services"
)

// RouteGuideServer is route object
type RouteGuideServer struct {
	mu sync.Mutex // protects routeNotes
}

// Testing grpc demo function
func (s *RouteGuideServer) Testing(ctx context.Context, agentInfo *pb.Request) (*pb.Response, error) {
	return &pb.Response{Id: 1, Msg: "first grpc testing successful"}, nil
}

//NewServer is new routeGuideServer object
func NewServer() *RouteGuideServer {
	s := &RouteGuideServer{}
	return s
}
