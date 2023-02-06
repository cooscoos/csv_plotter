package plotread

// Define rules for unmarshalling Traffic Flow csv rows.
type TFlow struct {
	Hour   int     `csv:"hr"`
	Count0 float32 `csv:"M0_count"`
	Count1 float32 `csv:"M1_count"`
}
