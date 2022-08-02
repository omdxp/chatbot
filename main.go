package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/Krognol/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

func main() {
	godotenv.Load(".env")

	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient := &wolfram.Client{
		AppID: os.Getenv("WOLFRAM_APP_ID"),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Question: ")
		scanner.Scan()
		question := scanner.Text()
		if question == "exit" {
			break
		}

		response, err := client.Parse(&witai.MessageRequest{
			Query: question,
		})
		if err != nil {
			log.Fatal(err)
		}
		data, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		value := gjson.Get(string(data), "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
		if value.Exists() {
			answer, err := wolframClient.GetSpokentAnswerQuery(value.String(), wolfram.Metric, 1000)
			if err != nil {
				log.Fatal(err)
			}
			println("Answer: " + answer)
		} else {
			log.Println("No Wolfram Search Query")
		}
	}
}
