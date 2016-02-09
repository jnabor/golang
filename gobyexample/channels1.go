// Channels are the pipes that connect concurrent

// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.

// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// values. This specificity increases the type-safety of
// the program.

package main

import (
	"fmt"
	"time"
)

// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.package main

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channels1() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages = make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
