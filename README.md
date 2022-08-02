# chatbot

Chatbot built with Wit.at and Wolfram Alpha API.

It basically answer any common questions.

## Example

```sh
Question: What is the capital of Algeria?
Answer: The capital city of Algeria is Algiers, Algeria
Question: exit
```

## Setup

1. Get an API key from [Wit.ai](https://wit.ai/).
2. Get an API key from [Wolfram Alpha](https://developer.wolframalpha.com/).
3. Create a `.env` file with the following content:

```sh
WIT_AI_TOKEN=<WIT_TOKEN>
WOLFRAM_APP_ID=<WOLFRAM_TOKEN>
```

## Usage

```sh
go run main.go
```
