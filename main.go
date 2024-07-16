package main

import (
	"log"
	"net/http"

	"banking-app/api"
	"banking-app/config"
)

func main() {
	// Konfigürasyonu yükle
	config.LoadConfig()

	// Veritabanı bağlantısını başlat
	config.InitDB()

	// Router'ı başlat
	router := api.Init()

	// Sunucuyu başlat
	log.Fatal(http.ListenAndServe(":8080", router))
}
