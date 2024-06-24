package main

import (
	"fmt"
	"math/rand"
	"os"
	"bufio"
	"strings"
)

func Game() string {
	options := [3]string{"r", "p", "s"}
	randomInt := rand.Intn(len(options))
	bot := options[randomInt]
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Let's Play the Game! It's Your Turn [r, p or s]\nYou: ")
	user, _ := reader.ReadString('\n')
	user = strings.ToLower(strings.TrimSpace(user))
	
	fmt.Printf("Bot: %v\n", bot)
	switch {
	case (user == "r" && bot == "r"), (user == "p" && bot == "p"), (user == "s" && bot == "s"):
		fmt.Println("Tie!")
	case (user == "r" && bot == "p"), (user == "p" && bot == "s"), (user == "s" && bot == "r"):
		fmt.Println("Bot Wins!")
	case (user == "r" && bot == "s"), (user == "p" && bot == "r"), (user == "s" && bot == "p"):
		fmt.Println("You Win! Great Job!!")
	default:
		return "exit"
	}
	return ""
}

func mainer() {
	// Goal: Rock, Paper, Scissor
	for {
		if Game() == "exit" {
			break
		}
		Game()
	}
}