package sort

import (
	"fmt"
)

func QuickSort (arr interface{}) {


	var arrString []string
	var arrInt []int
	var arrByte []uint8

	switch arr.(type) {
	case *[]string:
		for _, v := range *arr.(*[]string) {
			arrString = append(arrString, v)
		}
	case *[]int:
		for _, v := range *arr.(*[]int) {
			arrInt = append(arrInt, v)
		}
	case *[]interface{}:
		sort(arr)
	case *[][]uint8:
		for _, v := range *arr.(*[]uint8) {
			arrByte = append(arrByte, v)
		}
	default:
		break
	}

	if len(arrInt) > 0 {
		 sortInt(arrInt, 0, len(arrInt) - 1)
		 *arr.(*[]int) = arrInt
		 return
	} else if len(arrString) > 0 {
		sortString(arrString, 0, len(arrString) - 1)
		*arr.(*[]string) = arrString
		return
	}
}

func sortInt(arr []int, start int, end int) {
	if end-start < 1 {
		return
	}
	pivot := arr[end]
	splitIndex := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[splitIndex] = arr[splitIndex], arr[i]
			splitIndex++
		}
	}
	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot
	sortInt(arr, 0, splitIndex-1)
	sortInt(arr, splitIndex+1, end)
}
func sortString(arr []string, start int, end int) {
	if end-start < 1 {
		return
	}
	pivot := arr[end]
	splitIndex := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[splitIndex] = arr[splitIndex], arr[i]
			splitIndex++
		}
	}
	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot
	sortString(arr, 0, splitIndex-1)
	sortString(arr, splitIndex+1, end)

}
func sort (arr interface{}) {
	var arrByte []string
	for _,v := range *arr.(*[]interface{}) {
		string := fmt.Sprint(v)
		arrByte = append(arrByte, string)
	}
	QuickSort(&arrByte)

	var arrInterface []interface{}

	for _, v := range arrByte {
		data := fmt.Sprintf("%s", v)
		arrInterface = append(arrInterface, data)
	}
	*arr.(*[]interface{}) = arrInterface
}










