package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseBhavFromCSV(data [][]byte) ([]Bhav, error) {
	bhavs := make([]Bhav, 0)
	for _, row := range data {
		tokens := strings.Split(string(row), ",")
		if tokens[0] == "SYMBOL" {
			// skip it
			continue
		}

		bhav, err := parseBhav(string(row))
		if err != nil {
			return nil, fmt.Errorf("failed to parse bhav from csv, err: %v", err)
		}
		bhavs = append(bhavs, bhav)
		// reader := csv.NewReader(bytes.NewReader(row))
		// records, err := reader.ReadAll()
		// if err != nil {
		// 	return nil, err
		// }

		// for _, record := range records {
		// 	if record == "SYMBOL" {
		// 		// skip this is header
		// 	}
		// 	bhav, err := parseBhav(record)
		// 	if err != nil {
		// 		return nil, fmt.Errorf("failed to parse bhav from csv, err: %v", err)
		// 	}

		// 	bhavs = append(bhavs, bhav)
		// }
	}

	return bhavs, nil
}

func parseBhav(data string) (Bhav, error) {
	bhav := Bhav{}

	tokens := strings.Split(data, ",")
	// NOTE: quick parser, since we know the values and their positions we can directly parse them
	// otherwise, please use reflection or use some third party library that can parse a csv into a structure
	bhav.Symbol = tokens[0]
	bhav.Series = tokens[1]

	// NOTE:
	// can tokens[2] == nil ?? if so, handle the case
	open, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return bhav, err
	}
	bhav.Open = open

	high, err := strconv.ParseFloat(tokens[3], 64)
	if err != nil {
		return bhav, err
	}
	bhav.High = high

	low, err := strconv.ParseFloat(tokens[4], 64)
	if err != nil {
		return bhav, err
	}
	bhav.Low = low

	close, err := strconv.ParseFloat(tokens[5], 64)
	if err != nil {
		return bhav, err
	}
	bhav.Close = close

	last, err := strconv.ParseFloat(tokens[6], 64)
	if err != nil {
		return bhav, err
	}
	bhav.Last = last

	prevClose, err := strconv.ParseFloat(tokens[7], 64)
	if err != nil {
		return bhav, err
	}
	bhav.PrevClose = prevClose

	totalRDQTY, err := strconv.ParseInt(tokens[8], 10, 64)
	if err != nil {
		return bhav, err
	}
	bhav.TotalRDQTY = int32(totalRDQTY)

	totalRDValue, err := strconv.ParseFloat(tokens[9], 64)
	if err != nil {
		return bhav, err
	}
	bhav.TotalRDVal = totalRDValue

	date, err := time.Parse("01-Jan-2006", tokens[10])
	if err != nil {
		return bhav, err
	}
	bhav.Timestamp = date

	totalTrades, err := strconv.ParseInt(tokens[11], 10, 64)
	if err != nil {
		return bhav, err
	}
	bhav.TotalTrades = int32(totalTrades)
	bhav.ISIN = tokens[12]

	return bhav, nil
}
