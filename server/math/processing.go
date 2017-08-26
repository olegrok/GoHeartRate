//Package math has only one external function named Calculate
//this function takes two parametrs: array of raw data
//and quantity of it. Returns calculated heart rate
//func Calculate(int, []float64) float64
package math

import (
	"fmt"
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
)

const blequ = 21

func aver(times []float64) float64 {
	aver := 0.0
	for i := 0; i < len(times); i++ {
		aver = aver + times[i]
	}
	return aver / float64(len(times))
}

// Calculate : function to start all calculations
func Calculate(data []float64, times []float64) float64 {
	//fmt.Println(n, data)
	n := len(data)
	avgsmpltm := aver(times)
	fmt.Println("avgsmpltm =", avgsmpltm)
	max := ftncalc(avgsmpltm, n, filter(n, data, generateSoMeSHIT(2.0*math.Pi/avgsmpltm, 10.0, blequ)))
	fmt.Println("max =", max)
	return max * 60.0 / 2.0
}

func filter(n int, data []float64, bequ []float64) []float64 {

	filtered := data[0 : n-blequ]
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
func generateSoMeSHIT(OmegaC float64, bW float64, lNum int) []float64 {
	OmegaLow := OmegaC - bW/2.0
	OmegaHigh := OmegaC + bW/2.0
	var FirCoeff []float64
	for j := 0; j < lNum; j++ {
		Arg := float64(j) - float64(lNum-1)/2.0

		if Arg == 0.0 {
			FirCoeff = append(FirCoeff, 0.0)
		} else {

			FirCoeff = append(FirCoeff, (math.Cos(OmegaLow*Arg*math.Pi)-math.Cos(OmegaHigh*Arg*math.Pi))/math.Pi/Arg)
		}
	}

	fmt.Println(FirCoeff)
	return FirCoeff
}

//func generate(n int, alpha float64) [blequ]float64 {
//	var w [blequ]float64
//	two := 2.0
//
//	for i := 0; i < n; i++ {
//		if ((alpha / 2) * float64(n-1)) > float64(i) {
//
//			w[i] = (0.5) * (1 + math.Cos(math.Pi*((two*float64(i))/(alpha*(float64(n-1)))-1)))
//		}
//		if ((alpha/2)*float64(n-1)) <= float64(i) && (float64(n-1)*(1-(alpha)/2)) >= float64(i) {
//			w[i] = 1
//		}
//		if ((float64(n-1) * (1 - (alpha)/2)) < float64(i)) && ((n - 1) >= i) {
//			w[i] = (0.5) * (1 + math.Cos(math.Pi*((two*float64(i))/(alpha*(float64(n-1)))+1-2/(alpha))))
//		}
//
//	}
//	fmt.Println("FILTER= ", w)
//	return w
//
//}

func ftncalc(avgsmpltm float64, n int, data []float64) float64 {
	c := fft.FFTReal(data)
	max := 0.0
	cordOur := 0.0
	for i := 0; i < n/2; i++ {
		a := cmplx.Abs(c[i])
		cord := float64(i) / (float64(n) * avgsmpltm)
		fmt.Println(a, cord)
		if (cord < 4) && (cord > 1) && (max < a) {
			max = a
			cordOur = cord
		}
	}
	return cordOur
}
