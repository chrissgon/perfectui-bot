package main

import (
	"log"
	"os"

	"github.com/chrissgon/lowbot"
)

func main() {
	lowbot.DEBUG = true

	// make a channel. In this exemple is Telegram
	channel, err := NewServerChannel()

	if err != nil {
		log.Println(err)
		return
	}

	// make consumer
	consumer, err := lowbot.NewChatGPTAssistantConsumer(os.Getenv("CHATGPT_TOKEN"), os.Getenv("CHATGPT_ASSISTANT"))

	if err != nil {
		log.Println(err)
		return
	}

	// start bot
	lowbot.StartConsumer(consumer, []lowbot.IChannel{channel})
}
