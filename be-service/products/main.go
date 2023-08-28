package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Item struct{
	Id	 string `json:"id"`
	Name string `json:"name"`
}

type Response struct{
	Text string `json:"message"`
	Data []Item `json:"data"`
}

func main() {
    http.HandleFunc("/products", productHander)
    port := "8084"
    fmt.Printf("Server is running on port %s...\n", port)
    http.ListenAndServe(":"+port, nil)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func productHander(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	var id = randomString(10)

	items := []Item{
		{Id : id, Name : "Clean Code",},
		{Id : id,Name : "Refactoring"},
		{Id : id,Name : "The Pragmatic Programming"},
	}

	response := Response{
        Text: "SUCCESS",
		Data: items,
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
        return
    }

	w.Write(jsonResponse)
}
