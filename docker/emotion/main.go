package main

import (
	"emotion/internal/config"
	"emotion/internal/web"
	"fmt"
)

const (
	DEFAULT_API_URL = "http://0.0.0.0:7860"
	DEFAULT_PORT    = "3000"
)

func main() {
	var apiUrl = config.SetDataFromEnv("GATE_PROXY_API", DEFAULT_API_URL)
	var appPort = config.SetDataFromEnv("GATE_APP_PORT", DEFAULT_PORT)

	fmt.Println("Proxy API Addr:", DEFAULT_API_URL)
	fmt.Println("Web Server port:", DEFAULT_PORT)

	web.API(apiUrl, appPort)
}
