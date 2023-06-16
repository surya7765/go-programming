// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var (
// 	counter int32          // counter is a variable incremented by all goroutines.
// 	wg      sync.WaitGroup // wg is used to wait for the program to finish.
// 	mutex   sync.Mutex     // mutex is used to define a critical section of code.
// )

// func main() {
// 	wg.Add(3) // Add a count of two, one for each goroutine.

// 	go increment("Python")
// 	go increment("Go Programming Language")
// 	go increment("Java Programming Language")

// 	wg.Wait() // Wait for the goroutines to finish.
// 	fmt.Println("Counter:", counter)

// }

// func increment(lang string) {
// 	defer wg.Done() // Schedule the call to Done to tell main we are done.
// 	for i := 0; i < 3; i++ {
// 		mutex.Lock()
// 		{
// 			fmt.Println(lang)
// 			counter++
// 		}
// 		mutex.Unlock()
// 	}
// }

// package main

// import "fmt"

// func fib(n int, ch chan<- int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		ch <- x
// 		x, y = y, x+y
// 	}
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int, 10) // create a buffered channel with a capacity of 10
// 	go fib(15, ch)           // generate the first 10 Fibonacci numbers in a separate goroutine
// 	// read values from the channel until it's closed
// 	for x := range ch {
// 		fmt.Println(x)
// 	}
// }

package main

import (
	"fmt"
	"time"
)

// goroutine that sends the message
func SendMessage(msgChannel chan string, ackChannel chan string) {
	for {
		msgChannel <- "sending message @" + time.Now().String()
		time.Sleep(2 * time.Second)
		ack := <-ackChannel
		fmt.Println(ack)
	}
}

// goroutine that receives the message
func ReceiveMessage(msgChannel chan string, ackChannel chan string) {
	for {
		message := <-msgChannel
		fmt.Println(message)
		ackChannel <- "message received @" +
			time.Now().String()
	}
}

func main() {
	channel := make(chan string)
	// Channel to acknowledge message  receipt by receiver
	ackChnl := make(chan string)

	go SendMessage(channel, ackChnl)
	go ReceiveMessage(channel, ackChnl)
	fmt.Scanln()
}
