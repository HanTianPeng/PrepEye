package main

import "fmt"

func modifyArray(nums [3]int) {
	nums[0] = 100
	fmt.Println(nums) // [100, 2, 3]
}

func modifySlice(nums []int) {
	nums[0] = 100
	fmt.Println(nums) // [100, 2, 3]
}

func handle1() {
	nums := [3]int{1, 2, 3}
	modifyArray(nums)
	fmt.Println(nums) // [1, 2, 3]
}

func handle2() {
	nums := []int{1, 2, 3}
	modifySlice(nums)
	fmt.Println(nums) // [100, 2, 3]
}

// handle3: 如果扩容发生后,切片的底层数组会被重新分配,原切片和新切片底层数组不再共享
func handle3() {
	slices := make([]int, 0, 2)
	slices = append(slices, 1)
	slices = append(slices, 2)

	slices2 := slices

	fmt.Println(slices)  // [1, 2]
	fmt.Println(slices2) // [1, 2]

	slices2 = append(slices2, 3)

	fmt.Println(slices)  // [1, 2]
	fmt.Println(slices2) // [1, 2]
}

func modifySlices4(nums []int) {
	nums = append(nums, 1)
	fmt.Println(nums) // [1, 2, 3]
}

// handle4: 如果扩容发生后,切片的底层数组会被重新分配,原切片和新切片底层数组不再共享
func handle4() {
	slices := make([]int, 0, 2)
	slices = append(slices, 1)
	slices = append(slices, 2)
	modifySlices4(slices)
	fmt.Println(slices) // [1, 2]
}

func main() {
	handle1()
	fmt.Println("===1===")
	handle2()
	fmt.Println("===2===")
	handle3()
	fmt.Println("===3===")
	handle4()
	fmt.Println("===4===")
}
