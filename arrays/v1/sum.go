package main

func Sum(numbers [5]int) int {
	sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}
	//用range 让函数更简洁
	for _, v := range numbers {
		sum += v
	}
	return sum
}
