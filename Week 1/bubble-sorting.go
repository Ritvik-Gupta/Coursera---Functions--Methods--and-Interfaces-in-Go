package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bufioInput := bufio.NewReader(os.Stdin)

	lineBytes, _, _ := bufioInput.ReadLine()

	if len(lineBytes) == 0 {
		fmt.Println()
		return
	}

	numStrings := strings.Split(string(lineBytes), " ")

	if len(numStrings) > 10 {
		panic("Maximum Number of allowed integers in the array are 10")
	}

	numsArr := make([]int, 0, len(numStrings))
	for _, numStr := range numStrings {
		num, _ := strconv.Atoi(numStr)
		numsArr = append(numsArr, num)
	}

	BubbleSort(numsArr)

	for _, num := range numsArr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func BubbleSort(nums []int) {
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		for j := 0; j < numsLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func Swap(nums []int, index int) {
	temp := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = temp
}
