package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mozilla-ai/any-llm-go/providers/openai"
)

func main() {
	ctx := context.Background()

	provider, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	models, err := provider.ListModels(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, model := range models.Data {
		fmt.Println(model.ID)
	}
}
