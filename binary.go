package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	stop := make(chan bool)

	// Start time
	startTime := time.Now()
	go func(startTime time.Time) {
		defer fmt.Println("Execution stopped.") // Print "Execution stopped" message when the goroutine exits
		for {
			select {
			case <-stop:
				// Calculate and display the time it took to stop
				fmt.Printf("Time to stop: %s\n", time.Since(startTime))
				return
			default:
				// Task to be executed
				time.Sleep(time.Second) // Simulating some work
			}
		}
	}(startTime)
	reader := bufio.NewReader(os.Stdin)
	list := []int{0, 35, 33, 22, 98, 77, 88, 23, 56, 74, 45, 44, 34, 43, 63, 54}
	sort.Ints(list)

	left := 0
	right := len(list) - 1
	middle := left + ((right - left) / 2)

	fmt.Print("[0, 35, 33, 22, 98, 77, 88, 23, 56, 74, 45, 44, 34, 43, 63, 54]\nPick a Number: ")
	user, _ := reader.ReadString('\n')
	usr, _ := strconv.Atoi(strings.ToLower(strings.TrimSpace(user)))
	fmt.Println(list)
	for {
		if list[0] == usr {
			fmt.Printf("\nNumber %v found at location %v at the first instance\n", usr, 0)
			stop <- true
			break
		}
		if list[right] == usr {
			fmt.Printf("\nNumber %v found at location %v at the last instance\n", usr, right)
			stop <- true
			break
		}
		if right < left {
			fmt.Println("\nNot Found Anywhere Bro!")
			break
		} else {
			middle = left + ((right - left) / 2)

			if list[middle] > usr {
				fmt.Printf("\n%v is greater than chosen %v", list[middle], usr)
				right = middle - 1
			} else if list[middle] < usr {
				fmt.Printf("\n%v is less than chosen %v", list[middle], usr)
				left = middle + 1
			} else {
				fmt.Printf("\nNumber %v found at location %v between %v and %v\n", usr, middle, list[middle-1], list[middle+1])
				stop <- true
				break
			}
		}
	}

}
