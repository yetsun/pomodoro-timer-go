package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// RestMinutes rest 5 minutes
const RestMinutes = 5

//FocusMinutes focus 25 minutes
const FocusMinutes = 25

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("What do you plan to do: ")
		input, _ := reader.ReadString('\n')

		task := strings.TrimSpace(input)

		fmt.Println(task)

		if strings.EqualFold(task, "rest") {
			rest()
		} else {
			focus(task)
		}
	}

}

func rest() {
	for i := 0; i < RestMinutes; i++ {

		fmt.Printf("Rest"+". Minute: %d\n", i)

		time.Sleep(1 * time.Minute)
		// time.Sleep(1 * time.Second)
	}

	ring()
	fmt.Println("Rest done. Go back to work.")
}

func focus(task string) {
	fmt.Println("Focus on " + task)

	for i := 0; i < FocusMinutes; i++ {

		fmt.Printf("Focus on "+task+". Minute: %d\n", i)

		time.Sleep(1 * time.Minute)
		// time.Sleep(1 * time.Second)
	}

	ring()
	fmt.Println("Focus done. Rest now")
}

func ring() {
	fmt.Print("\a")
	time.Sleep(1 * time.Second)
	fmt.Print("\a")
	fmt.Print("\a")
	fmt.Print("\a")
	fmt.Print("\a")
}
