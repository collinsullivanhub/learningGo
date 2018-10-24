
package main 

import (
"fmt"
"time"
"sync"
)

//Concurrency

func main(){

	var waitYO sync.WaitGroup
	waitYO.Add(1)
	go func() {
		count("sheep")
		waitYO.Done()
	}()
	waitYO.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++{
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

