package main

import (
    "github.com/sev-2/raiden"
    "log"
)

func main() {
    app := raiden.New()
    
    // Tambahkan konfigurasi atau middleware jika diperlukan
    
    log.Println("Starting Raiden server on :8080")
    app.Start(":8080")
}
