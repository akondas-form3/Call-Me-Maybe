package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/youshy/call-me-maybe/grpc/humans"
	"github.com/youshy/call-me-maybe/types"
	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 56565, "The server port")
	howmany = flag.Int("howmany", 500000, "How many humans should be created in a map")
)

type humanServer struct {
	pb.UnimplementedHumansServer
	humans map[int]types.Human
}

func (s *humanServer) GetHuman(ctx context.Context, in *pb.ID) (*pb.Human, error) {
	human, ok := s.humans[int(in.Id)]
	if !ok {
		return &pb.Human{}, errors.New("Unable to find a human!")
	}

	return mapHumanToPB(human), nil
}

func (s *humanServer) CreateHuman(ctx context.Context, in *pb.Human) (*pb.ID, error) {
	h := mapPBToHuman(in)

	h.ID = len(s.humans) + 1

	s.humans[h.ID] = h

	return &pb.ID{Id: int32(h.ID)}, nil
}

func makeHumans(howmany int) map[int]types.Human {
	hm := make(map[int]types.Human, howmany)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < howmany; i++ {
		human := types.Human{
			FirstName:  fmt.Sprint(rand.Intn(1000000)),
			LastName:   fmt.Sprint(rand.Intn(1000000)),
			Age:        rand.Intn(100),
			LikesPizza: true,
		}
		hm[i] = human
	}

	return hm
}

func mapHumanToPB(human types.Human) *pb.Human {
	return &pb.Human{
		Id:         int32(human.ID),
		Firstname:  human.FirstName,
		Lastname:   human.LastName,
		Age:        int32(human.Age),
		Likespizza: human.LikesPizza,
	}
}

func mapPBToHuman(human *pb.Human) types.Human {
	return types.Human{
		FirstName:  human.Firstname,
		LastName:   human.Lastname,
		Age:        int(human.Age),
		LikesPizza: human.Likespizza,
	}
}

func newServer(amount int) *humanServer {
	s := &humanServer{
		humans: makeHumans(amount),
	}
	return s
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterHumansServer(grpcServer, newServer(*howmany))

	grpcServer.Serve(lis)
}
