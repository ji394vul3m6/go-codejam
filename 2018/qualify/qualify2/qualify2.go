package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		length, _ := strconv.Atoi(strings.TrimSpace(line))

		line, _ = reader.ReadString('\n')
		strNums := strings.Split(strings.TrimSpace(line), " ")
		nums := make([]int, length)
		for i, str := range strNums {
			n, _ := strconv.Atoi(str)
			nums[i] = n
		}

		ok, idx := qualify2(nums[0:length])
		if ok {
			fmt.Printf("Case #%d: OK\n", i+1)
		} else {
			fmt.Printf("Case #%d: %d\n", i+1, idx)
		}
	}
}

func qualify2(nums []int) (bool, int) {
	defer func() {
		// fmt.Println("=========================")
	}()

	// fmt.Println("=========================")
	// fmt.Printf("Get: %#v\n", nums)

	length := len(nums)

	oddNums := make([]int, (length+1)/2)
	evenNums := make([]int, length/2)

	for idx, n := range nums {
		if idx%2 == 0 {
			oddNums[idx/2] = n
		} else {
			evenNums[(idx-1)/2] = n
		}
	}
	// fmt.Printf("oddNums: %+v\n", oddNums)
	// fmt.Printf("evenNums: %+v\n", evenNums)
	sort.Sort(sort.IntSlice(oddNums))
	sort.Sort(sort.IntSlice(evenNums))
	// fmt.Printf("oddNums after sort: %+v\n", oddNums)
	// fmt.Printf("evenNums after sort: %+v\n", evenNums)

	i := 0
	j := 0
	for i < len(oddNums) && j < len(evenNums) {
		if evenNums[j] < oddNums[i] {
			return false, i * 2
		}
		if i+1 < len(oddNums) && oddNums[i+1] < evenNums[j] {
			return false, (j * 2) + 1
		}
		i++
		j++
	}

	return true, 0
}
