package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func getURL() string {
	return ""
}

func fetchBhavData(url string) ([]Bhav, error) {
	// hit with an http request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bhav data, err: %v", err)
	}
	resp, err := sendHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := readHTTPResponseBody(resp)
	if err != nil {
		return nil, err
	}

	// extract the csv file from the zip
	fileVsdata, err := ReadZip(body)
	if err != nil {
		return nil, err
	}

	// NOTE:
	// this is a quick and dirty hack
	// since we know that there'll be only file inside the zip, so we can skip the check
	// please append a check if you think that more than one file could be present in the zip
	var csvData [][]byte
	for _, data := range fileVsdata {
		csvData = parseCSV(data)
		break
	}

	// parse the csv
	bhavs, err := parseBhavFromCSV(csvData)
	return bhavs, err
}

func parseCSV(data []byte) [][]byte {
	csv := make([][]byte, 0)
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		val := scanner.Text()
		// if val == "SYMBOL" then skip it
		// fmt.Println(string(val))
		csv = append(csv, []byte(val))
	}

	return csv
}

func getBhavFromCache(symbol, series string) *Bhav {
	key := getCacheKey(symbol, series)
	value, exists := GetCache().Get(key)
	if !exists {
		return nil
	}

	return value.(*Bhav)
}

func restoreCacheWithBhavData(dateTime time.Time) error {
	// fetch data from http endpoint
	for {
		url := getBhavURL(dateTime)
		log.Printf("fetching data from nse with url: %v\n", url)
		bhavs, err := fetchBhavData(url)
		if err != nil {
			log.Printf("failed to fetch data from nse with url: %v, will try for the previous day, err: %v\n", url, err)
			time.Sleep(time.Second * 1)
			dateTime = dateTime.AddDate(0, 0, -1)
			continue
		}

		if len(bhavs) == 0 {
			// let us go one day back
			// and retry it after a sleep of 1 second,
			// we don't want to bombard the nse servers with requests,
			time.Sleep(time.Second * 1)
			dateTime = dateTime.AddDate(0, 0, -1)
			continue
		}

		// we have fetched data successfully, now we can populate our cache
		// and then quit from our loop
		for _, bhav := range bhavs {
			b := bhav
			key := getCacheKey(bhav.Symbol, bhav.Series)
			GetCache().Add(key, &b)
		}
		break
	}

	return nil
}

func getCacheKey(symbol, series string) string {
	return "sym-" + symbol + "::" + "series-" + series
}

func getBhavURL(dateTime time.Time) string {
	url := GetConfig().Bhav.URL
	year := fmt.Sprintf("%d", dateTime.Year())
	month := fmt.Sprintf("%s", strings.ToUpper(dateTime.Month().String()[:3]))
	day := fmt.Sprintf("%02d", dateTime.Day())

	newDate := fmt.Sprintf("%s%s%s", day, month, year)

	url = strings.Replace(url, "{YEAR}", year, 1)
	url = strings.Replace(url, "{MONTH}", month, 1)
	url = strings.Replace(url, "{DATE}", newDate, 1)

	return url
}
