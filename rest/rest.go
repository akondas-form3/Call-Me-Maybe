package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/youshy/call-me-maybe/types"
)

// Create One Human
func CreateHuman(humans map[int]types.Human) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var human types.Human

		err := json.NewDecoder(r.Body).Decode(&human)
		if err != nil {
			JSONResponse(w, http.StatusBadRequest, err)
			return
		}

		// Super smart way of generating IDs
		human.ID = len(humans) + 1

		humans[human.ID] = human

		JSONResponse(w, http.StatusCreated, nil)
	})
}

// Get a human
func GetHuman(humans map[int]types.Human) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		key, ok := r.URL.Query()["human"]
		if !ok {
			JSONResponse(w, http.StatusNotFound, nil)
			return
		}

		id, err := strconv.Atoi(key[0])
		if err != nil {
			JSONResponse(w, http.StatusBadRequest, err)
			return
		}

		human, ok := humans[id]
		if !ok {
			JSONResponse(w, http.StatusNotFound, nil)
			return
		}

		JSONResponse(w, http.StatusOK, human)
	})
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	response, _ := json.Marshal(output)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
