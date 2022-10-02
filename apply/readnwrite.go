package apply

import (
	"fmt"
	"os"
)

func Check(i []string) string {
	if len(i) != 2 {
		return ""
	}
	return i[1]
}

func Read(s string) string {
	if len(s) == 0 {
		return ""
	}
	if s[len(s)-4:] != ".txt" {
		return ""
	}
	content, err := os.ReadFile(s)
	if err != nil {
		return ""
	}
	return string(content)
}

func PrintAll(nums []float64) {
	if len(nums) == 0 {
		fmt.Println("Empty file or unsuitable content/filename/overflow")
		return
	}
	fmt.Println((Avg(nums)))
	fmt.Println((Median(nums)))
	fmt.Println((Variance(nums)))
	fmt.Println(SD(nums))
}
