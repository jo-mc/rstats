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

}
