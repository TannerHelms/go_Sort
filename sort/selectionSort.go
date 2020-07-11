package sort

func SelectionSort (arr interface{}) {


	var arrString []string
	var arrInt []int

	switch arr.(type) {
	case *[]string:
		for _, v := range *arr.(*[]string) {
			arrString = append(arrString, v)
		}
	case *[]int:
		for _, v := range *arr.(*[]int) {
			arrInt = append(arrInt, v)
		}
	default:
		break
	}

	if len(arrInt) > 0 {
		selInt(arrInt)
		*arr.(*[]int) = arrInt
		return
	} else if len(arrString) > 0 {
		selString(arrString)
		*arr.(*[]string) = arrString
		return
	}
}
func selInt(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		minInx := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minInx] {
				minInx = j
				arr[minInx], arr[i] = arr[i], arr[minInx]
			}
		}
	}
}
func selString(arr []string) {
	for i := 0; i < len(arr) - 1; i++ {
		minInx := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minInx] {
				minInx = j
				arr[minInx], arr[i] = arr[i], arr[minInx]
			}
		}
	}
}
