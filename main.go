package main

import (
	"os"

	"github.com/chrissgon/lowbot"
)

func main() {
	// make a channel. In this exemple is Telegram
	channel, _ := NewServerChannel()

	// make consumer
	consumer, _ := lowbot.NewChatGPTAssistantConsumer(os.Getenv("CHATGPT_TOKEN"), "asst_29b3txBZzptV1xXAyHuIBm2L")

	// start bot
	lowbot.StartConsumer(consumer, []lowbot.IChannel{channel})
}
