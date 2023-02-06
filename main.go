package main

import (
	"csv_plotter/src/plotread"
)

// Read in some data from csv and plot it as a time series
func main() {

	// Plot and save traffic data
	plotread.PlotCsv("1b14_traffic")

	// Plot and save air quality data
	plotread.PlotCsv("1b14_outair")
}
