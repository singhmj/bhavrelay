package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func readConfig(path string) (*SystemConfig, error) {
	configFilename := "service.json"
	path += "/" + configFilename
	config := SystemConfig{}
	err := parseJSON(path, &config)
	if err != nil {
		return nil, err
	}
	SetConfig(&config)
	return &config, nil
}

func startWebServer() error {
	config := GetConfig()
	address := config.Dumper.WebServer.Address + ":" + strconv.Itoa(config.Dumper.WebServer.Port)
	log.Printf("starting web-server at: %v\n", address)
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/bhav_copy/{symbol}/{series}", GetBhavData).Methods("GET")
	return http.ListenAndServe(address, router)
}

func start() error {
	cache := NewCache()
	SetCache(cache)

	log.Println("restoring bhav data...")
	// restore bhav
	if err := restoreCacheWithBhavData(time.Now()); err != nil {
		return err
	}

	// start bhav scheduler
	go func() {
		log.Println("starting scheduler...")
		if err := startScheduler(); err != nil {
			panic(fmt.Errorf("failed to start scheduler, err: %v", err))
		}
	}()

	log.Println("starting api-server...")
	// start web-server to serve requests
	if err := startWebServer(); err != nil {
		return err
	}
	return nil
}

func main() {

	// read config
	configPath := os.Args[1]
	if configPath == "" {
		panic("Config path is missing, usage:  ./dumper [config_path]")
	}

	if _, err := readConfig(configPath); err != nil {
		panic(fmt.Errorf("failed to read config: %v", err))
	}

	// initiate the system
	err := start()
	if err != nil {
		panic(fmt.Errorf("couldn't initiate the system, more: %v", err))
	}

	// TODO: add interrupt handler
	// TODO: handle data across multiple go-routines for performance and efficiency
	// TODO: gracefully shutdown
	// TODO: restructure the program into packages
	// TODO: update readme
	// TODO: add description to readme
	// TODO: improve cache implementation
	// TODO: improve csv parsing
}
