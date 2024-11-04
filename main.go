package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenAI API key not set in environment")
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessageParts(
				openai.TextPart("What's in this image?"),
				openai.ImagePart("https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"),
			),
		}),
		Model:       openai.String("gpt-4o"),
		Temperature: openai.Float(0),
	})

	fmt.Println(chatCompletion)

	if err != nil {
		panic(err.Error())
	}
}
