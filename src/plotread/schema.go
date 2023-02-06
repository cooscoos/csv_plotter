package plotread

// Struct for unmarshalling traffic flow csv rows.
type TFlow struct {
	Hour   int     `csv:"hr"`
	Count0 float32 `csv:"M0_count"`
	Count1 float32 `csv:"M1_count"`
}

// Struct for unmarshalling outdoor air quality csv rows.
type AQual struct {
	Hour int     `csv:"hr"`
	CO2  float32 `csv:"out_CO2"`
	PM   float32 `csv:"out_PM"`
}
