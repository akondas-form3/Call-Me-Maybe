package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/youshy/call-me-maybe/rest"
	"github.com/youshy/call-me-maybe/types"
)

func main() {
	humans := make(map[int]types.Human, 0)

	addr := "127.0.0.1:9999"

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5000000; i++ {
		human := types.Human{
			FirstName:  fmt.Sprint(rand.Intn(1000000)),
			LastName:   fmt.Sprint(rand.Intn(1000000)),
			Age:        rand.Intn(100),
			LikesPizza: true,
		}

		humans[i] = human
	}

	router := mux.NewRouter()

	router.Handle("/get", rest.GetHuman(humans)).Methods(http.MethodGet)
	router.Handle("/create", rest.CreateHuman(humans)).Methods(http.MethodPost)

	log.Printf("Server is listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
