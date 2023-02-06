package plotread

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

// Parse data from r and return a table of unmarshaled values.
func ParseData(r io.Reader) ([]TFlow, error) {

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
