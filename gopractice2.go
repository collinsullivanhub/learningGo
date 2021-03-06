
package main 

import (
"fmt"
"time"
)

func main() {

	//Anonymous functions
	nextInt := intSeq()
	fmt.Println(nextInt())	//prints 1
	fmt.Println(nextInt())	//prints 2
	fmt.Println(nextInt())  //prints 3

	nextInt2 := intSeq()
	fmt.Println(nextInt2()) //prints 1

	//Recursion
	fmt.Println(fact(7))

	//Pointers
	s := device{name:"IOS",age:23}
	fmt.Println(s.name) // IOS
	sp := &s
	fmt.Println(sp.age) //23
	sp.age = 24
	fmt.Println(sp.age) //24

	//Goroutine say function call
	go say("this call is as a goroutine: world")
	//normal say function call
	say("this is calling the function normally: hello")
	
	//Make a channel of the int type
	channelOne := make(chan int)
	go foo(channelOne, 5)
	go foo(channelOne, 3)
	v1 := <- channelOne
	v2 := <- channelOne
	fmt.Println(v1, v2)
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

//Channels
func foo(c chan int, someValue int){
	c <- someValue * 5
}




