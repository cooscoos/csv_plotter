package main

import (
	"encoding/csv"
	"image/color"
	"io"
	"log"

	"fmt"
	"os"

	"github.com/jszwec/csvutil"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Define rules for unmarshalling Traffic Flow csv rows.
type TFlow struct {
	Hour   int     `csv:"hr"`
	Count0 float32 `csv:"M0_count"`
	Count1 float32 `csv:"M1_count"`
}

// Parse data from r and return a table of unmarshaled values.
func parseData(r io.Reader) ([]TFlow, error) {

	var table []TFlow

	csvReader := csv.NewReader(r)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return make([]TFlow, 0), err
	}

	// Read in header to get past it. We won't need it, so don't store.
	dec.Header()

	// Read in and unmarshal the remaining rows
	for {
		var u TFlow
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			return []TFlow{}, err
		}
		table = append(table, u)
	}

	return table, nil

}

// Take in a table of TFlow values and plot them to a png of given outname.
func plotData(table []TFlow, outname string) error {

	data := make(plotter.XYs, len(table))

	for i := range data {
		data[i].X = float64(table[i].Hour)
		data[i].Y = float64(table[i].Count0 + table[i].Count1)
	}

	p := plot.New()
	p.Title.Text = "Traffic flows"
	p.X.Label.Text = "Hour of week"
	p.Y.Label.Text = "Motorway traffic (vehicles)"
	p.Add(plotter.NewGrid())

	line, points, err := plotter.NewLinePoints(data)
	if err != nil {
		return err
	}

	line.Color = color.RGBA{G: 255, A: 255}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{R: 255, A: 255}

	p.Add(line, points)

	err = p.Save(10*vg.Centimeter, 5*vg.Centimeter, fmt.Sprintf("./Output/%v.png", outname))
	if err != nil {
		return err
	}

	return nil
}

// Read in some data from csv and plot it as a time series
func main() {

	filename := "1b14_traffic"

	file, err := os.Open(fmt.Sprintf("Data/%v.csv", filename))

	if err != nil {
		log.Panic(err)
	}

	table, err := parseData(file)

	if err != nil {
		log.Panic(err)
	}

	// Plot data and save to ./Output/test with the same name as input csv
	if plotData(table, filename) != nil {
		log.Panic(err)
	}

}
