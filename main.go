package main

import (
	"os"

	"github.com/chrissgon/lowbot"
)

func main() {
	lowbot.DEBUG = true

	// make a channel. In this exemple is Telegram
	channel, _ := NewServerChannel()

	// make consumer
	consumer, _ := lowbot.NewChatGPTAssistantConsumer(os.Getenv("CHATGPT_TOKEN"), os.Getenv("CHATGPT_ASSISTANT"))

	// start bot
	lowbot.StartConsumer(consumer, []lowbot.IChannel{channel})
}
