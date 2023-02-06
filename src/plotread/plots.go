package plotread

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Take in a table of TFlow values and plot them to a png of given outname.
func PlotData(data plotter.XYs, outname string) error {

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

// Populate time series with traffic data
func populateTraffic(table []TFlow) plotter.XYs {
	data := make(plotter.XYs, len(table))

	for i := range data {
		data[i].X = float64(table[i].Hour)
		data[i].Y = float64(table[i].Count0 + table[i].Count1)
	}

	return data
}

// Populate time series with CO2 data
func populateAir(table []AQual) plotter.XYs {
	data := make(plotter.XYs, len(table))

	for i := range data {
		data[i].X = float64(table[i].Hour)
		data[i].Y = float64(table[i].CO2)
	}

	return data
}
