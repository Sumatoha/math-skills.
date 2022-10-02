package apply

import (
	"math"
	"strconv"
	"strings"
)

func Convert(s string) []float64 {
	if len(s) == 0 {
		return nil
	}
	answer := []float64{}
	temp := strings.Split(s, "\n")
	for i := 0; i < len(temp); i++ {
		if temp[i] != "" {
			num, _ := strconv.Atoi(temp[i])
			answer = append(answer, float64(num))
		}
	}
	return answer
}

func Overflow(num []float64) []float64 {
	for i := 0; i < len(num); i++ {
		if num[i] > math.MaxInt64 {
			return []float64{}
		}
	}
	return num
}
