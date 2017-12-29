package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

// RestMinutes rest 5 minutes
const RestMinutes = 5

//FocusMinutes focus 25 minutes
const FocusMinutes = 25

func progress(secondCount int, startWord string, endWord string) {

	fmt.Println("------------------------ " + startWord + " ------------------------")

	bar := pb.StartNew(secondCount)
	for i := 0; i < secondCount; i++ {
		bar.Increment()
		time.Sleep(time.Second)
	}

	bar.FinishPrint("------------------------ " + endWord + " ------------------------")
	ring()
	fmt.Println()
	fmt.Println()
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("What do you plan to do: ")
		input, _ := reader.ReadString('\n')
		task := strings.TrimSpace(input)

		if task != "" {
			if strings.EqualFold(task, "rest") {
				rest()
			} else {
				focus(task)
			}
		}
	}

}

func rest() {
	progress(RestMinutes*60, "Rest", "Rest done. Go back to work.")
}

func focus(task string) {
	progress(FocusMinutes*60, "Focus on "+task, "Focus done. Rest now.")
}

func ring() {
	fmt.Print("\a")
	time.Sleep(1 * time.Second)
	fmt.Print("\a")
	fmt.Print("\a")
	fmt.Print("\a")
	fmt.Print("\a")
}
