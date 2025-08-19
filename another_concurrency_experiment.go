package main

import (
	"fmt"
)

func square(nums []int, ch chan int){
	for _, n := range nums {
		ch <- n * n
	}
	close(ch)
}

//func say(s string){
//	for i := 0; i < 3; i++ {
//		fmt.Println(s)
//		time.Sleep(500 * time.Millisecond)
//	}
//}

func main(){
//	go say("world") // runs in a goroutine
//	say("hello")

	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	go square(numbers, ch)

	fmt.Println("This will be run right away even though square() is not done")
	
	// receive from channel
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("But this will only run after reading all the numbers from channel")
}
