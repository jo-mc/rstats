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
	fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(18.1, rS)
	fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(16.5, rS)
	fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(18.3, rS)
	fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))
	rs.RollingStat(12.6, rS)
	fmt.Println("Mean :", rS.M1, "Std Dev: ", math.Sqrt(rS.M2/((float64(rS.N))-1.0)))

	datajain := []float64{0.02, 0.15, 0.74, 3.39, 0.83,
		22.37, 10.15, 15.43, 38.62, 15.92,
		34.60, 10.28, 1.47, 0.40, 0.05,
		11.39, 0.27, 0.42, 0.09, 11.37}
	// test_quantile(0.5, data_jain, n_jain, expected_jain, tol1, "jain");
	rQ := new(rs.RQuant)
	rs.Reinit(rQ, 0.5)
	for i := 0; i < 20; i++ {
		rs.QuantRoller(datajain[i], rQ)
	}

	fmt.Println("Quant for 0.5 : ", rs.RQuantResult(rQ))
}
