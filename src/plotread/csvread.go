package plotread

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jszwec/csvutil"
)

// Parse data from filename and return a table of unmarshaled values.
// Generics should be constrained here and defined in schema.go.
func ParseData[T TFlow | AQual](filename string) ([]T, error) {

	file, err := os.Open(fmt.Sprintf("Data/%v.csv", filename))

	if err != nil {
		log.Panic(err)
	}

	var table []T

	csvReader := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(csvReader)

	if err != nil {
		return []T{}, err
	}

	// Read in header to get past it. We won't need it, so don't store.
	dec.Header()

	// Read in and unmarshal the remaining rows
	for {
		var u T
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			return []T{}, err
		}
		table = append(table, u)
	}

	return table, nil
}
