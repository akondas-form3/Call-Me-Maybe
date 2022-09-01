package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	pb "github.com/youshy/call-me-maybe/grpc/humans"
	"github.com/youshy/call-me-maybe/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BenchmarkCreateHumansGRPC(b *testing.B) {
	serverAddr := "localhost:56565"

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		b.Fatal("Unable to dial grpc:", err)
	}
	defer conn.Close()

	client := pb.NewHumansClient(conn)

	b.ResetTimer()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		human := types.Human{
			FirstName:  fmt.Sprint(rand.Intn(1000000)),
			LastName:   fmt.Sprint(rand.Intn(1000000)),
			Age:        rand.Intn(100),
			LikesPizza: true,
		}

		_, err := createHuman(client, human)
		if err != nil {
			b.Fatal("Create:", err)
		}
	}
}

func BenchmarkGetHumansGRPC(b *testing.B) {
	serverAddr := "localhost:56565"

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		b.Fatal("Unable to dial grpc:", err)
	}
	defer conn.Close()

	client := pb.NewHumansClient(conn)

	b.ResetTimer()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		_, err := getHuman(client, i)
		if err != nil {
			b.Fatal("Create:", err)
		}
	}
}
