
package main 

import (
"fmt"
"time"
)

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

	//Pointers
	s := device{name:"IOS",age:23}
	fmt.Println(s.name) // IOS
	sp := &s
	fmt.Println(sp.age) //23
	sp.age = 24
	fmt.Println(sp.age) //24

	//Goroutine function calls
	go say("world")
	say("hello")
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

type device struct {
	name string
	age int
}

//Goroutines 
func say(s string) {
	for i:=0; i < 1; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

