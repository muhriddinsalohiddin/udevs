package syntax

import "fmt"

func Fibonacci(count int) {
	var first, second int = 0, 1
	for count > 0 {
		fmt.Println(first)
		first, second = second, first+second
		count--
	}
}

func FizzBuzz(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("Fizz Buzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	}
}

func Palindrome(check string) {
	j := len(check)
	for i := range check {
		j--
		if check[i] != check[j] {
			fmt.Println("Not palindrome")
			return
		}
	}
	fmt.Println("Palindrome")

}
func OddEvenSum(num int, isEven string) {
	var sum int
	for i := 0; i <= num; i++ {
		if i%2 == 0 && isEven == "even" {
			sum += i
		} else if i%2 != 0 && isEven == "odd" {
			sum += i
		}
	}
	fmt.Println(sum)
}
func Duplicate(nums []int) {
	isDuplicate := make(map[int]bool)
	for i := range nums {
		if isDuplicate[nums[i]] {
			fmt.Println(true, nums[i])
			return
		} else {
			isDuplicate[nums[i]] = true
		}
	}
	fmt.Println(false)
}
