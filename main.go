package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/typesense/typesense-go/v2/typesense"
	"github.com/typesense/typesense-go/v2/typesense/api"
	"github.com/typesense/typesense-go/v2/typesense/api/pointer"
)

func main() {
	client := typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("xyz"),
	)

	// Delete collection if exists
	_, err := client.Collection("companies").Delete(context.Background())
	if err != nil {
		log.Printf("Delete warning: %v", err) // Non-fatal if collection doesn't exist
	}

	// Create collection
	schema := &api.CollectionSchema{
		Name: "companies",
		Fields: []api.Field{
			{Name: "company_name", Type: "string"},
			{Name: "num_employees", Type: "int32"},
			{Name: "country", Type: "string"},
		},
		DefaultSortingField: pointer.String("num_employees"),
	}
	_, err = client.Collections().Create(context.Background(), schema)
	if err != nil {
		log.Fatalf("Error creating collection: %v", err)
	}

	// Create documents
	documents := []struct {
		ID           string `json:"id"`
		CompanyName  string `json:"company_name"`
		NumEmployees int    `json:"num_employees"`
		Country      string `json:"country"`
	}{
		{
			ID:           "123",
			CompanyName:  "Stark Industries",
			NumEmployees: 5215,
			Country:      "USA",
		},
		{
			ID:           "456",
			CompanyName:  "Wayne Enterprises",
			NumEmployees: 6000,
			Country:      "USA",
		},
	}
	log.Println("Inserting documents...")

	for _, doc := range documents {
		_, err := client.Collection("companies").Documents().Upsert(context.Background(), doc)
		if err != nil {
			log.Fatalf("Error inserting document: %v", err)
		}
		log.Printf("Inserted document: %s", doc.ID)
	}

	// Search
	searchParameters := &api.SearchCollectionParams{
		Q:        pointer.String("enterp"),
		QueryBy:  pointer.String("company_name"),
		FilterBy: pointer.String("num_employees:>100"),
		SortBy:   pointer.String("num_employees:desc"),
	}

	result, err := client.Collection("companies").Documents().Search(context.Background(), searchParameters)
	if err != nil {
		log.Fatalf("Search error: %v", err)
	}

	// Proper JSON marshaling
	jsonResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshal error: %v", err)
	}

	log.Printf("Search results:\n%s", jsonResult)
}
