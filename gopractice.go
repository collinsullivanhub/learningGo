//package 
package main

//imports
import (
"fmt"
"io/ioutil"
"log"
)

func main() {

	fmt.Println("Printing...")

	//variables
	var variableOne int = 7
	var variableTwo int = 13
	var sum int = variableOne + variableTwo
	fmt.Println(sum)

	//conditionals
	if variableOne > 10 {
		fmt.Println("variableOne is greater than 10...")
	} else if variableOne < 10 {
		fmt.Println("variableOne is less than 10...")
	} else {
		fmt.Println("Error...")
	}

	//arrays
	var arrayOne [5]int
	fmt.Println(arrayOne)
	arrayOne[2] = 7
	fmt.Println("Printing array one after assigning 7 to array...")
	fmt.Println(arrayOne)
	//now [0,0,7,0,0] 
	//can also do: 
	arrayTwo := [5]int{5,4,3,2,1}
	fmt.Println("Printing array two: ")
	fmt.Println(arrayTwo)

	//slices

	//maps, kind of like dictionaries
	myMap := make(map[string]int)
	myMap["ios"] = 256
	myMap["iosxr"] = 31
	myMap["iosxe"] = 42
	fmt.Println(myMap)

	//looping
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}

	myarr := []string{"a","b","c"}

	for index, value := range myarr{
		fmt.Println("index: ", index, "value: ", value)
	}

	//call sumFunc
	mathResult := sumFunc(2,3)
	fmt.Println(mathResult)

	//Use "device" struct data type
	d1 := device{name: "IOS", age: 23, ipAddress:"10.122.1.10", isOn:true}
	fmt.Println(d1.name)
	fmt.Println(d1.age)
	fmt.Println(d1.ipAddress)
	fmt.Println(d1.isOn)

	//logging
	log.Print("logging...") 

	//Open a file
	b, err := ioutil.ReadFile("file1.txt")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Below are file contents as bytes: ")
		fmt.Println(b) 	//print file contents as bytes
		srtringify := string(b)
		fmt.Println(srtringify) //print the file contents as a string
	}

	//Memory
	meep := 7
	fmt.Println(meep)
	meep2 := &meep //memory address
	fmt.Println(meep2)
	fmt.Println(*meep2)

	//object of car struct
	a_car := car{gas_pedal: 22341,
				 brake_pedal: 0,
				 steering_wheel: 12561,
				 top_speed_kmh: 225.0}
	fmt.Println(a_car.gas_pedal)

}

//func sum(name type, name type) type{}
func sumFunc(x int, y int) int{
	return x + y
}

//struct
type device struct{
	name string
	age int
	ipAddress string
	isOn bool
}

//methods
//value receivers = calculations on values 
//pointer receivers = modifying values from a struct

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 //0 to 65535
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}
//c can be whateveh, since this function is referencing a struct it is a method
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}



