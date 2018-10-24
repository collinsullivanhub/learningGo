package main 

import ("fmt")


func main() {

	//Anonymous functions
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//prints 1,2,3

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	//prints 1

	//Recursion
	fmt.Println(fact(7))
}

//Anonymous function 
func intSeq() func() int {
	i := 0
	return func() int{
		i++
		return i		
	}
}

//Recursive function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
