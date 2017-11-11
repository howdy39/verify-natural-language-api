package main

import (
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1"
	"encoding/json"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	text := "Natural Language API を使うのは楽しいです"

	analyzeEntities(ctx, client, text)
	analyzeSentiment(ctx, client, text)
	analyzeSyntax(ctx, client, text)
}

func analyzeEntities(ctx context.Context, client *language.Client, text string) {
	resp, err := client.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}
	rawJson, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to Marshal json: %v", err)
	}
	jsonStr := string(rawJson)
	fmt.Println("analyzeEntities\n" + jsonStr + "\n")
}

func analyzeSentiment(ctx context.Context, client *language.Client, text string) {
	resp, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}
	rawJson, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to Marshal json: %v", err)
	}
	jsonStr := string(rawJson)
	fmt.Println("analyzeSentiment\n" + jsonStr + "\n")
}

func analyzeSyntax(ctx context.Context, client *language.Client, text string) {
	resp, err := client.AnnotateText(ctx, &languagepb.AnnotateTextRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		Features: &languagepb.AnnotateTextRequest_Features{
			ExtractSyntax: true,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}
	rawJson, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to Marshal json: %v", err)
	}
	jsonStr := string(rawJson)
	fmt.Println("analyzeSyntax\n" + jsonStr + "\n")
}
