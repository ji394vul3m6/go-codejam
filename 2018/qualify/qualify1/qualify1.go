package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	strNum, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))
	// fmt.Printf("get nums: %d\n", num)

	for i := 0; i < num; i++ {
		line, _ := reader.ReadString('\n')
		params := strings.Split(strings.TrimSpace(line), " ")
		d, _ := strconv.Atoi(params[0])
		pattern := params[1]

		hasSol, count := qualify1(d, pattern)
		if !hasSol {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
		} else {
			fmt.Printf("Case #%d: %d\n", i+1, count)
		}
	}
}

func qualify1(d int, pattern string) (hasSolution bool, changeCount int) {
	defer func() {
		// fmt.Println("=============================")
	}()
	shootCount := 0
	chargeCount := 0
	nowDamage := 1
	nowTotal := 0

	// fmt.Println("=============================")
	for _, char := range pattern {
		if char == 'C' {
			nowDamage *= 2
			chargeCount++
		} else if char == 'S' {
			shootCount++
			nowTotal += nowDamage
		}
	}
	// fmt.Printf("Input %d: %s\n", d, pattern)
	// fmt.Printf("Init total shoot damage: %d\n", nowTotal)

	if shootCount > d {
		hasSolution = false
		return
	}
	hasSolution = true

	for calculateDamage(pattern) > d {
		changeCount++
		rightestShoot := 0
		for i := len(pattern) - 1; i >= 0; i-- {
			if pattern[i] == 'S' {
				rightestShoot = i
				break
			}
		}
		target := 0
		for i := rightestShoot - 1; i >= 0; i-- {
			if pattern[i] == 'C' {
				target = i
				break
			}
		}
		pattern = swap(pattern, target, target+1)
		// fmt.Printf("Change %d: pattern: %s, damage: %d\n", changeCount, pattern, calculateDamage(pattern))
	}

	return
}

func calculateDamage(pattern string) int {
	nowDamage := 1
	nowTotal := 0
	for _, char := range pattern {
		if char == 'C' {
			nowDamage *= 2
		} else if char == 'S' {
			nowTotal += nowDamage
		}
	}
	return nowTotal
}

func swap(in string, i int, j int) string {
	out := []rune(in)
	out[i], out[j] = out[j], out[i]
	return string(out)
}
