package plotread

import (
	"fmt"
	"regexp"

	"gonum.org/v1/plot/plotter"
)

func PlotCsv(filename string) error {

	var data plotter.XYs

	// Check the filename for the presence of the word "traffic"
	s, err := regexp.Match(`traffic`, []byte(filename))

	if err != nil {
		return err
	}

	// Handle traffic and air quality data differently
	if s {
		fmt.Println("Plotting traffic data")
		table, err := ParseData[TFlow](filename)

		if err != nil {
			return err
		}

		data = populateTraffic(table)
	} else {
		fmt.Println("Plotting air quality data")
		table, err := ParseData[AQual](filename)

		if err != nil {
			return err
		}
		data = populateAir(table)

	}

	// Plot data and save to ./Output/test with the same name as input csv
	if PlotData(data, filename) != nil {
		return err
	}
	return nil
}
