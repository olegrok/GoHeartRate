//Package has only one external function named Calculate
//this function takes two parametrs: array of raw data
//and quantity of it. Returns calculated heart rate
//func Calculate(int, []float64) float64
package math

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
	"time"
)

type everything struct {
	x []float64
	y []float64
	n int
}

const blequ = 21

// Calculate : function to start all calculations
func Calculate(n int, data []float64) float64 {
	time.Sleep(1 * time.Second)
	return 42

	max := ftnCalc(n, filter(n, data))

	return max * 60 * math.Pi / 2
}

func filter(n int, data []float64) []float64 {
	bequ := [blequ]float64{-0, 0.000447346177473, 0.009905463720835, 0.01396310914665,
		-0.01762554215192, -0.06309569775754, -0.06381106741101, 0.004885452302831,
		0.1252345953647, 0.2400982483146, 0.2872795190548, 0.2400982483146,
		0.1252345953647, 0.004885452302831, -0.06381106741101, -0.06309569775754,
		-0.01762554215192, 0.01396310914665, 0.009905463720835, 0.000447346177473,
		-0}
	filtered := data[0:n]
	num := 0

	for i := blequ; i < n-blequ; i++ {
		filter := 0.0
		for j := 0; j < blequ; j++ {
			filter += data[i-(blequ/2)] * bequ[j]
		}
		filtered[num] = filter
		num++
	}
	return filtered
}

func ftnCalc(n int, data []float64) float64 {
	c := fft.FFTReal(data)
	max := 0.0
	for i := 0; i < n/2; i++ {
		a := cmplx.Abs(c[i])
		cord := i * (1000 / (n * 40))
		if (cord < 4) && (cord > 1) && (max < a) {
			max = a
		}
	}
	return max
}
