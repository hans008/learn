package main

import "fmt"

//裴伯纳切数列
//1 1 2 3 5 8 13.....


func fib() func() int {
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}

}


func main(){
	f := fib()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}