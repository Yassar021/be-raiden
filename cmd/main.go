package main

import (
	"log"
	"github.com/sev-2/raiden"
	"be-raiden/controller"  // Make sure this matches your module name
	"be-raiden/services"    // Make sure this matches your module name
)

func main() {
	// Initialize the InstrumentService
	instrumentService := services.NewInstrumentService()

	// Initialize the InstrumentController
	instrumentController := controller.NewInstrumentController(instrumentService)

	// Create the Raiden app
	app := raiden.New()

	// Set up routes
	app.GET("/instruments", instrumentController.GetInstruments)

	// Start the server
	log.Println("Server running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
