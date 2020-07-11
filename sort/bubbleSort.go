package sort


func BubbleSort (arr interface{}) {
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
	case []interface{}:
		sort(arr)
	default:
		break

	}

	if len(arrInt) > 0 {
		bubbleInt(arrInt)
		*arr.(*[]int) = arrInt
		return
	} else if len(arrString) > 0 {
		bubbleString(arrString)
		*arr.(*[]string) = arrString
		return
	}
}

func bubbleString(arr []string) {
	n := len(arr)

	for i := 0; i < n - 1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleInt(arr []int) {
	n := len(arr)

	for i := 0; i < n - 1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

