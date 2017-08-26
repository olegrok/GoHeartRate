package protocol

// MathData contains data from measurements
type MathData struct {
	DataArray []uint32  `json:"data_array"`
	TimeArray []float64 `json:"time_array"`
}
