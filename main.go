package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run main.go <username>")
		os.Exit(1)
	}

	username := os.Args[1]
	fmt.Println("Looking for activities for:", username)

	url := "https://api.github.com/users/" + username + "/events"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Connection successful!")
	fmt.Println("Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	var events []map[string]interface{}

	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

    fmt.Printf("\n--- Atividade Recente de %s ---\n", username)

	for _, event := range events {
		if event["type"] == "PushEvent" {

			repoMap := event["repo"].(map[string]interface{})
			repoName := repoMap["name"]
			fmt.Println()

			payloadMap := event["payload"].(map[string]interface{})
			numCommits := int(payloadMap["size"].(float64))

			createdAtString := event["created_at"].(string)

			layout := "2006-01-02T15:04:05Z"
			eventTime, err := time.Parse(layout, createdAtString)
			if err != nil {
				fmt.Println("Erro ao ler data do evento:", err)
			}

			formattedTime := eventTime.Format("02/01/2006 às 15:04")
			fmt.Printf("[%s] - Pushed %d commits to %s\n", formattedTime, numCommits, repoName)

		}

	}
}
