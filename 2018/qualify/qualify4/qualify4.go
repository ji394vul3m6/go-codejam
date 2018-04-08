package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	strNum, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))

	for i := 0; i < num; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		A, _ := strconv.ParseFloat(line, 64)
		fmt.Printf("Case #%d:\n", i+1)
		qualify4(A)
	}
}

func qualify4(area float64) {
	if area <= 1.414213 {
		a := (area - math.Sqrt(2-area*area)) / 2
		b := (area + math.Sqrt(2-area*area)) / 2
		fmt.Printf("%.12f %.12f 0\n", (area-a)/2, (area-b)/2)
		fmt.Printf("%.12f %.12f 0\n", (b-area)/2, (area-a)/2)
		fmt.Println("0 0 0.5")
	} else {
		sin := (area - math.Sqrt(6-2*area*area)) / 3
		cos := math.Sqrt(1 - sin*sin)

		p1 := []interface{}{math.Sqrt(2) / 4, -1 * math.Sqrt(2) / 4 * cos, -1 * math.Sqrt(2) / 4 * sin}
		p2 := []interface{}{math.Sqrt(2) / 4, math.Sqrt(2) / 4 * cos, math.Sqrt(2) / 4 * sin}
		p3 := []interface{}{0.0, 0.5 * sin, 0.5 * cos}

		fmt.Printf("%.12f, %.12f, %.12f\n", p1...)
		fmt.Printf("%.12f, %.12f, %.12f\n", p2...)
		fmt.Printf("%.12f, %.12f, %.12f\n", p3...)
	}
}
