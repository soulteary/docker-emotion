package main

import (
	"fmt"
	"emotion/internal/config"
	"emotion/internal/web"
)

const (
	DEFAULT_API_URL     = "http://127.0.0.1:7860"
	DEFAULT_PORT        = "3000"
)

func main() {
	var apiUrl = config.SetDataFromEnv("GATE_PROXY_API", DEFAULT_API_URL)
	var appPort = config.SetDataFromEnv("GATE_APP_PORT", DEFAULT_PORT)

	fmt.Println("Proxy API Addr:", DEFAULT_API_URL)
	fmt.Println("Web Server port:", DEFAULT_PORT)

	web.API(apiUrl, appPort)
}