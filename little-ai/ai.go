package ai

import (
	"fmt"
	"math/rand"
	"bufio"
	"os"
	"strings"
)

var responses = []string{
	"pull yourself together.",
	"it is what is it.",
	"maybe.",
	"ask again later.",
	"without a doubt.",
	"very doubtful.",
	"i know but u cant understand.",
	"r u killer?",
	"im sure is that you are not going to be able to do that right now so shut up.",
	"u r very smart?",
	"im good enough.",
	"why.",
	"you",
	"lol.",
	"<3",
	"maybe you",
	"have",
	"had",
	"pls go own way.",
	"don't think about anything",
	"so be it",
	"i wantn't saying nothing",
	"i love unconditionally u",
	"Te amo",
}

func getRandomResponse() string {
	
	randomIndex := rand.Intn(len(responses))
	return responses[randomIndex]
}

func Xai() {
	fmt.Println("ask me anything. if u want exit in me u should write 'exit'")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("ask: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}
		question := strings.ToLower(strings.TrimSpace(input))
		if question == "exit" {
			fmt.Println("see u later.")
			break
		}
		response := getRandomResponse()
		fmt.Printf("answer: %s\n", response)
	}
}
