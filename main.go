package main

import (
	"bufio"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Question: ")
		scanner.Scan()
		question := scanner.Text()
		if question == "exit" {
			break
		}
	}
}
