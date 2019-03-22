package main

import "fmt"

func quickSort(values []int, left, right int) int  {
	temp := values[left]
	for left < right {
		// 当队尾指针大于基准数据时，向左移动靠右指针
		for left < right && values[right] >= temp {
			right--
		}
		// 指向队尾的时候最后一个值进行变更
		values[left] = values[right]
		for left < right && values[left] <= temp {
			left++
		}
		values[right] = values[left]
	}
	values[left] = temp
	return left
}


func getIndex(values []int, left, right int) {

	if left < right {
		// 分治, 取得index值
		index := quickSort(values, left, right)
		quickSort(values, 0, index-1)
		quickSort(values, index, right)
	}
}



func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	getIndex(values, 0, len(values)-1)
}


func main(){
	numArray := []int{2,5,1,3,4}
	QuickSort(numArray)
	for _, n := range(numArray) {
		fmt.Print(n)
	}

}