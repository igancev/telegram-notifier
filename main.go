package main

import (
	"./console"
	"./telegram"
)

func main() {
	args := console.ParseArguments()

	telegramApi := telegram.Api{}.New(args.Token)

	telegramApi.SendMessage(telegram.Message{}.New(args.ChatId, args.Message))
}
