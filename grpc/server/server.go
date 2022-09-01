package main

import (
	pb "github.com/youshy/call-me-maybe/grpc/humans"
)

type server struct {
	pb.UnimplementedHumansServer
}
