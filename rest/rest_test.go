package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/youshy/call-me-maybe/types"
)

func BenchmarkCreateHumansREST(b *testing.B) {
	// Create the database
	humans := make(map[int]types.Human, 0)

	l, err := net.Listen("tcp", "127.0.0.1:9998")
	if err != nil {
		b.Fatal("error spinning up server", err)
	}

	// Start the server
	ts := httptest.NewUnstartedServer(CreateHuman(humans))

	// Add custom listener
	ts.Listener.Close()
	ts.Listener = l

	ts.Start()
	defer ts.Close()

	// Create transport for the client
	// and create the client
	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	// Reset the benchmark timer
	b.ResetTimer()

	// Get 'em seeds in!
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		// We're living in the future, henceforth, names have been
		// abandonned and we use only numbers.
		human := types.Human{
			FirstName:  fmt.Sprint(rand.Intn(1000000)),
			LastName:   fmt.Sprint(rand.Intn(1000000)),
			Age:        rand.Intn(100),
			LikesPizza: true,
		}

		// Build the request payload
		payload, err := json.Marshal(human)
		if err != nil {
			b.Fatal("Get:", err)
		}

		// Send the payload
		res, err := cl.Post(ts.URL, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			b.Fatal("Post:", err)
		}
		defer res.Body.Close()

		if http.StatusCreated != res.StatusCode {
			b.Fatal("Unable to create a human!")
		}
	}
}

// Again, create the database map
// The reason for that is that we want to have
// a completely separate state for both
var gethumans = make(map[int]types.Human, 0)

func init() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 500000; i++ {
		human := types.Human{
			FirstName:  fmt.Sprint(rand.Intn(1000000)),
			LastName:   fmt.Sprint(rand.Intn(1000000)),
			Age:        rand.Intn(100),
			LikesPizza: true,
		}
		gethumans[i] = human
	}
}

func BenchmarkGetHumansREST(b *testing.B) {
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		b.Fatal("error spinning up server", err)
	}

	ts := httptest.NewUnstartedServer(GetHuman(gethumans))

	ts.Listener.Close()
	ts.Listener = l

	ts.Start()
	defer ts.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		url := fmt.Sprintf("%s?human=%v", ts.URL, j)

		// Try to get a human with a random ID
		res, err := cl.Get(url)
		if err != nil {
			b.Fatal("Get:", err)
		}
		defer res.Body.Close()

		if http.StatusOK != res.StatusCode {
			b.Fatalf("Unable to get a human with id of %v!", j)
		}
	}
}
