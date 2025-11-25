package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Cleaned string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedBody := badWordsReplacer(params.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		Cleaned: cleanedBody,
	})
}

func badWordsReplacer(originalString string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(originalString, " ")
	for index, word := range words {
		wordInLower := strings.ToLower(word)
		if slices.Contains(badWords, wordInLower) {
			words[index] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}
