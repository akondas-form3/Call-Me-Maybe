package main

import (
	"context"
	"time"

	pb "github.com/youshy/call-me-maybe/grpc/humans"
	"github.com/youshy/call-me-maybe/types"
)

func getHuman(client pb.HumansClient, id int) (types.Human, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	human, err := client.GetHuman(ctx, &pb.ID{
		Id: int32(id),
	})
	if err != nil {
		return types.Human{}, err
	}

	return mapPBToHuman(human), nil
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

func createHuman(client pb.HumansClient, human types.Human) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pb := mapHumanToPB(human)

	id, err := client.CreateHuman(ctx, pb)
	if err != nil {
		return 0, err
	}

	return int(id.Id), nil
}

func mapPBToHuman(human *pb.Human) types.Human {
	return types.Human{
		FirstName:  human.Firstname,
		LastName:   human.Lastname,
		Age:        int(human.Age),
		LikesPizza: human.Likespizza,
	}
}
