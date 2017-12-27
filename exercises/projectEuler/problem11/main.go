package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func combo(nums []int) []int {
	var maxSize int
	var maxValue int

	maxSize = 20
	maxValue = 0
	var total4d, total4r, total4diagr, total4diagl int
	var calcs []int
	indexNum := 0
	fmt.Printf("row, col, indexNum, nums[indexNum], total4d, total4r, total4diagl, total4diagr\n")
	for row := 0; row < maxSize; row++ {
		for col := 0; col < maxSize; col++ {
			total4d = 0
			total4r = 0
			total4diagr = 0
			total4diagl = 0
			indexNum = (20 * row) + col

			if row < 16 {
				total4d = nums[indexNum] * nums[indexNum+20] * nums[indexNum+40] * nums[indexNum+60]
				if total4d > maxValue {
					maxValue = total4d
				}

			}
			if col < 16 {
				total4r = nums[indexNum] * nums[indexNum+1] * nums[indexNum+2] * nums[indexNum+3]
				if total4r > maxValue {
					maxValue = total4r
				}

			}

			if col > 3 {
				if row < 16 {
					total4diagl = nums[indexNum] * nums[indexNum+19] * nums[indexNum+38] * nums[indexNum+57]
					if total4diagl > maxValue {
						maxValue = total4diagl
					}

				}

			}
			if col < 16 {
				if row < 16 {
					total4diagr = nums[indexNum] * nums[indexNum+21] * nums[indexNum+42] * nums[indexNum+63]
					if total4diagr > maxValue {
						maxValue = total4diagr
					}
				}

			}
			fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v\n", row, col, indexNum, nums[indexNum], total4d, total4r, total4diagl, total4diagr)

		}

		fmt.Println(maxValue)
		fmt.Println()

	}

	return calcs
}

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.S
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		ss := strings.Split(l, " ")
		/* 		fmt.Println(len(ss))
		   		fmt.Println(ss) */
		for i := 0; i < len(ss); i++ {
			n, err := strconv.Atoi(ss[i])
			if err != nil {
				return nil, err
			}
			nums = append(nums, n)
		}

	}

	return nums, nil
}

func main() {
	lenNums := 0
	nums, err := readFile("data.txt")
	if err != nil {
		panic(err)
	}
	lenNums = len(nums)
	fmt.Println(lenNums)
	fmt.Println(nums)

	/* 	for i := 0; i < lenNums; i++ {
		fmt.Printf("%v,%v\n", i, nums[i])
	} */
	combo(nums)

}
