package apply

import (
	"math"
	"sort"
	"strconv"
)

func Round(num float64) int {
	if int(num) > int(num-0.5) {
		return int(num)
	}
	return int(num + 1)
}

func Avg(nums []float64) string {
	temp := 0.0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
	}
	return "Average: " + strconv.Itoa(Round(temp/float64(len(nums))))
}

func Median(nums []float64) string {
	sort.Float64s(nums)
	if len(nums)%2 == 0 {
		return "Median: " + strconv.Itoa(Round((nums[len(nums)/2]+nums[len(nums)/2-1])/2))
	}
	return "Median: " + strconv.Itoa(int(nums[len(nums)/2]))
}

func SD(nums []float64) string {
	if (sVariance(nums)) > math.MaxInt-1 || (sVariance(nums)) <= math.MinInt {
		return "Standart Deviation: out of range"
	}
	return "Standart Deviation: " + strconv.Itoa(int(Round(math.Sqrt(float64(sVariance(nums))))))
}

func Variance(nums []float64) string {
	temp := []float64{}
	temp1 := (sAvg(nums))
	asd := nums
	for i := 0; i < len(asd); i++ {
		temp = append(temp, (asd[i]-temp1)*(asd[i]-temp1))
	}
	if Round(sAvg(temp)) > math.MaxInt-1 || Round(sAvg(temp)) <= math.MinInt {
		return "Variance: out of range"
	}
	return "Variance: " + strconv.Itoa(Round(sAvg(temp)))
}

func sVariance(nums []float64) float64 {
	temp := []float64{}
	temp1 := (sAvg(nums))
	for i := 0; i < len(nums); i++ {
		temp = append(temp, (nums[i]-temp1)*(nums[i]-temp1))
	}
	if int(sAvg(temp)) >= math.MaxInt || int(sAvg(temp)) < math.MinInt {
		return math.MaxInt
	}
	return (sAvg(temp))
}

func sAvg(nums []float64) float64 {
	temp := 0.0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
	}
	return (temp / float64(len(nums)))
}
