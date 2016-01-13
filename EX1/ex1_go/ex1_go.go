// Go 1.2
// go run helloworld_go.go

package main



import (
    . "fmt"
    "runtime"
    "time"
)

var counter int = 0 // filthy global scum

func threadTheFirst() {
	for i:= 1; i <= 1000000; i++ {
		counter++
	}

}

func threadTheSecond() {
	for i:= 1; i <= 1000000; i++ {
		counter--
	}

}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go threadTheFirst()
	go threadTheSecond()

    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println(counter)
}