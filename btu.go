package btu

import (
	"sort"
)

func Len(arr interface{}) int { //Len
	switch arr := arr.(type) {
	case []int:
		return len(arr)
	case []string:
		return len(arr)
	default:
		return 0
	}
}

func Reverse(arr interface{}) { //Reverse
	switch arr := arr.(type) {
	case []int:
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	case []string:
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func Sort(arr interface{}) { //Sort
	switch arr := arr.(type) {
	case []int:
		sort.Ints(arr)
	case []string:
		sort.Strings(arr)
	}
}

func Sum(arr interface{}) (sum int) { //Sum
	switch arr := arr.(type) {
	case []int:
		for _, num := range arr {
			sum += num
		}
	}
	return sum
}

func Max(arr interface{}) (max interface{}) { //Max
	switch arr := arr.(type) {
	case []int:
		if len(arr) > 0 {
			max = arr[0]
			for _, num := range arr[1:] {
				if num > max.(int) {
					max = num
				}
			}
		}
	case []string:
		if len(arr) > 0 {
			max = arr[0]
			for _, str := range arr[1:] {
				if str > max.(string) {
					max = str
				}
			}
		}
	}
	return max
}

func Min(arr interface{}) (min interface{}) { //Min
	switch arr := arr.(type) {
	case []int:
		if len(arr) > 0 {
			min = arr[0]
			for _, num := range arr[1:] {
				if num < min.(int) {
					min = num
				}
			}
		}
	case []string:
		if len(arr) > 0 {
			min = arr[0]
			for _, str := range arr[1:] {
				if str < min.(string) {
					min = str
				}
			}
		}
	}
	return min
}

func Copy(arr interface{}) interface{} { //Copy
	switch arr := arr.(type) {
	case []int:
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		return copyArr
	case []float64:
		copyArr := make([]float64, len(arr))
		copy(copyArr, arr)
		return copyArr
	case []string:
		copyArr := make([]string, len(arr))
		copy(copyArr, arr)
		return copyArr
	case []bool:
		copyArr := make([]bool, len(arr))
		copy(copyArr, arr)
		return copyArr
	default:
		return nil
	}
}
