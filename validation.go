package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type parameters struct {
	body string
}

type returnErr struct {
	Error string `json:"error"`
}

type resp struct {
	Valid bool `json:"valid"`
}

func validate_chirp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s\n", err)
		w.WriteHeader(500)
		return
	}

	if len(params.body) > 140 {
		error := returnErr{
			Error: "Chirp is too long",
		}

		dat, err := json.Marshal(error)
		if err != nil {
			log.Printf("Error unmarshalling json: %s", err)
			return
		}

		w.WriteHeader(400)
		w.Write(dat)
	}
	resp := resp{
		Valid: true,
	}
	dat, err := json.Marshal(resp)
	w.WriteHeader(200)
	w.Write(dat)
}
