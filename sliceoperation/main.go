package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	// s1 := slice[0:2]
	// s2 := slice[3:]
	// slice = append(s1, s2...)
	// fmt.Println(slice)
	var res []int
	for i := range slice {
		if slice[i] != 3 {
			res = append(res, slice[i])
		}
	}
	slice = res
	fmt.Println(slice)

	//{1,2,3,4,5,6}
	//2
	//{[1,3],[1,5],[2,4],[2,6],[3,5],[4,6]}

	//var slice [][]int

	//loop 1 i:=0;
	// loop 2 j=i+1,j<
	//  sum:= sli[i]+slice[j]
	// if(sum%2==0){
	//[1,3]
	//}
	sl := []int{1, 2, 3, 4, 5, 6}
	val := getPairs(sl)
	fmt.Printf("value %v \n", val)

	/// [2,10,1,3,5,8]
	// sort in ascending order
	// search the element 3

	slll := []int{2, 10, 1, 3, 5, 8}
	sortArray(slll)
	ele := searchElement(slll, 6)
	fmt.Println(ele)
}

func getPairs(nums []int) [][]int {

	var resultSlice [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum%2 == 0 {
				slice := []int{nums[i], nums[j]}
				resultSlice = append(resultSlice, slice)
			}
		}

	}
	return resultSlice

}

func sortArray(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func searchElement(arr []int, elementSearch int) int {

	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == elementSearch {
			return arr[mid]
		} else if arr[mid] < elementSearch {
			low = mid + 1
		} else if arr[mid] > elementSearch {
			high = mid - 1
		}

	}
	return -1

}
