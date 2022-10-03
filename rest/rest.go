package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/youshy/call-me-maybe/types"
)

// Create One Human takes a map of humans which represents the database.
// Normally there would be a call to the DB, but for the sake of simplicity,
// there isn't one.
func CreateHuman(humans map[int]types.Human) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create an empty struct
		var human types.Human

		// Build the decoder and decode struct into JSON
		err := json.NewDecoder(r.Body).Decode(&human)
		if err != nil {
			// Return the request with the error in JSON payload
			JSONResponse(w, http.StatusBadRequest, err)
			return
		}

		// Super smart way of generating IDs
		human.ID = len(humans) + 1

		// Assign the ID
		humans[human.ID] = human

		// Return the request with JSON payload
		JSONResponse(w, http.StatusCreated, nil)
	})
}

// Get a human takes a map of humans which represents the database.
// Takes a Query param to get the human's ID.
func GetHuman(humans map[int]types.Human) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Remember to close your body!
		defer r.Body.Close()

		// Check which human to fetch
		key, ok := r.URL.Query()["human"]
		if !ok {
			JSONResponse(w, http.StatusNotFound, nil)
			return
		}

		// Convert the strign to an integer
		id, err := strconv.Atoi(key[0])
		if err != nil {
			JSONResponse(w, http.StatusBadRequest, err)
			return
		}

		// Find our human in the map...database
		human, ok := humans[id]
		if !ok {
			JSONResponse(w, http.StatusNotFound, nil)
			return
		}

		JSONResponse(w, http.StatusOK, human)
	})
}

// JSONResponse is a utility method to generate a valid JSON response.
func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	// Marshal json
	response, _ := json.Marshal(output)

	// Write header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Write the response
	w.Write(response)
}
