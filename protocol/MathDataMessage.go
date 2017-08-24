package protocol

// MathData contains data from measurements
type MathData struct {
	DataArray []float64 `json:"data_array"`
	TimeArray []float64 `json:"time_array"`
}
