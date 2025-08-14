package concurrency

import (
	"fmt"
	"time"
)

func greet(name string, ch chan string) {
	time.Sleep(2 * time.Second) // simulate long running work
	msg := fmt.Sprintf("Hello %s!", name)
	ch <- msg
}

func Run() {
	ch := make(chan string)  // make a channel that works w strings
	go greet("Nnaemeka", ch) // send the channel into the function, remember to use 'go'
	fmt.Println(<-ch)        // print whatever u receive from the channel
}
