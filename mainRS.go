package main

import (
	"fmt"
	"math"
	"rstat/rs"
	//	"github.com/jo-mc/rstat/rs"
)

func main() {
	rS := new(rs.RStats)

	rs.RollingStat(17.2, rS)
	//fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(18.1, rS)
	//fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(16.5, rS)
	//fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(18.3, rS)
	//fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(12.6, rS)
	fmt.Println("The sample Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	fmt.Println("The estimated variance :", (rS.M2/(float64(rS.N) - 1)))
	fmt.Println("Largest Value :",rS.Max)
	fmt.Println("Smallest Value :",rS.Min)
	fmt.Println("Median : NA")
	fmt.Println("Standard Deviation : ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	if rS.N > 0 {
	fmt.Println("Skew :",((math.Pow(float64(rS.N)-1.0, 1.5)/float64(rS.N))*rS.M3) / (math.Pow(rS.M2, 1.5)))
	} else {
		fmt.Println("Skew : 0")
	}
	if rS.N > 0 {
	fmt.Println("Kurtosis :",(((float64(rS.N)-1.0)/float64(rS.N))*(float64(rS.N) - 1.0))*rS.M4 / ((rS.M2 * rS.M2)) - 3.0) 
	} else {
		fmt.Println("Kurtosis : 0")
	}
	fmt.Println("Number of items :",rS.N)


	// datajain := []float64{0.02, 0.15, 0.74, 3.39, 0.83,
	// 	22.37, 10.15, 15.43, 38.62, 15.92,
	// 	34.60, 10.28, 1.47, 0.40, 0.05,
	// 	11.39, 0.27, 0.42, 0.09, 11.37}
	datajain := []float64{1, 3, 5, 6, 9, 11, 12, 13, 19, 21, 
		22, 32, 35, 36, 45, 44, 55, 68, 79, 80, 81, 88, 90,
		91, 92, 100, 112, 113, 114, 120, 121, 132, 145, 146, 
		149, 150, 155, 180, 189, 190}		
	// test_quantile(0.5, data_jain, n_jain, expected_jain, tol1, "jain");
	rQ := new(rs.RQuant)
	rs.Reinit(rQ, 0.5)
	for i := 0; i < 40; i++ {
		rs.QuantRoller(datajain[i], rQ)
	}

	fmt.Println("Quant for 0.5 : ", rs.RQuantResult(rQ))
}
