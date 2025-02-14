package services

import (
	"fmt"
	"log"
	"os"
	"github.com/nedpals/supabase-go"
)

type InstrumentService struct {
	client *supabase.Client
}

func NewInstrumentService() *InstrumentService {
	supabaseURL := os.Getenv("https://goqvhmgdefqpfmqusmlg.supabase.co")
  supabaseAnonKey := os.Getenv("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImdvcXZobWdkZWZxcGZtcXVzbWxnIiwicm9sZSI6ImFub24iLCJpYXQiOjE3Mzk1MTE3NDAsImV4cCI6MjA1NTA4Nzc0MH0.A8Fsx15b5COywRps30GLc9ECB8XqinmOEnKA4PadEEA")

	client := supabase.NewClient(supabaseURL, supabaseAnonKey)

	return &InstrumentService{client: client}
}

func (s *InstrumentService) GetInstruments() (interface{}, error) {
	resp, err := s.client.From("instruments").Select("*").Execute()
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return nil, err
	}

	fmt.Println("Fetched data:", resp.Data)
	return resp.Data, nil
}
