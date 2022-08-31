package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func BenchmarkCreateHumans(b *testing.B) {
	humans := make(map[int]Human, 0)

	ts := httptest.NewServer(CreateHuman(humans))
	defer ts.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	b.ResetTimer()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		human := Human{
			FirstName:  fmt.Sprint(rand.Intn(1000000)),
			LastName:   fmt.Sprint(rand.Intn(1000000)),
			Age:        rand.Intn(100),
			LikesPizza: true,
		}

		payload, err := json.Marshal(human)
		if err != nil {
			b.Fatal("Get:", err)
		}

		res, err := cl.Post(ts.URL, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			b.Fatal("Post:", err)
		}

		if http.StatusCreated != res.StatusCode {
			b.Fatal("Unable to create a human!")
		}
	}
}
