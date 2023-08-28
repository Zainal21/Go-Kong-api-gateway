package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type User struct{
	Id	 string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type Response struct{
	Text string `json:"message"`
	Data []User `json:"data"`
}

func main() {
    http.HandleFunc("/users", userHandler)
    port := "8085"
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


func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	var id = randomString(10)

	response := Response{
        Text: "SUCCESS",
		Data: []User{
			{Id : id,Name : "Muhamnad", Address : "Kaliboto"},
			{Id : id,Name : "Zainal", Address: "Mojogedang"},
			{Id : id,Name : "Arifin", Address: "Karanganyar"},
		},
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
        return
    }

	w.Write(jsonResponse)
}
