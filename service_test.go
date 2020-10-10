package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGetBhavURL(t *testing.T) {
	dir, _ := os.Getwd()
	_, err := readConfig(dir + "/config/service.json")
	if err != nil {
		t.Fatal(err)
	}
	url := getBhavURL(time.Now())
	fmt.Println(url)
}

func TestRestoreCacheWithBhavData(t *testing.T) {
	dir, _ := os.Getwd()
	_, err := readConfig(dir + "/config/service.json")
	if err != nil {
		t.Fatal(err)
	}
	SetCache(NewCache())
	err = restoreCacheWithBhavData(time.Now())
	if err != nil {
		t.Fatal(err)
	}
}

func TestParser(t *testing.T) {
	_, err := time.Parse("01-Jan-2006", "10-OCT-2020")
	if err != nil {
		t.Fatal(err)
	}
}
