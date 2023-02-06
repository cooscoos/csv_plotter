package main

import (
	"csv_plotter/src/plotread"

	"fmt"
	"log"
	"os"
)

// Read in some data from csv and plot it as a time series
func main() {

	filename := "1b14_traffic"

	file, err := os.Open(fmt.Sprintf("Data/%v.csv", filename))

	if err != nil {
		log.Panic(err)
	}

	table, err := plotread.ParseData(file)

	if err != nil {
		log.Panic(err)
	}

	// Plot data and save to ./Output/test with the same name as input csv
	if plotread.PlotData(table, filename) != nil {
		log.Panic(err)
	}

}
