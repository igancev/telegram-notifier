package console

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Arguments struct {
	Token   string
	ChatId  int
	Message string
}

func ParseArguments() Arguments {
	token := flag.String("token", "", "Telegram bot token")
	chatId := flag.Int("chat_id", 0, "Telegram bot token")
	messageText := flag.String("message", "", "Telegram bot token")
	flag.Parse()

	arguments := Arguments{Token: *token, ChatId: *chatId, Message: *messageText}

	validate(&arguments)

	return arguments
}

func validate(arguments *Arguments) {
	if arguments.Token == "" {
		envToken := os.Getenv("TELEGRAM_NOTIFIER_TOKEN")
		if envToken == "" {
			fmt.Println("Please set token. Example: -token=\"1428033353:AAGg95hVqqW2qStmKlJHD0rXRG2vyMNVxlQ\"")
			os.Exit(1)
		}
		arguments.Token = envToken
	}

	if arguments.ChatId == 0 {
		envChatId := os.Getenv("TELEGRAM_NOTIFIER_CHAT_ID")
		if envChatId == "" {
			fmt.Println("Please set chat_id. Example: -chat_id=123456")
			os.Exit(1)
		}
		arguments.ChatId, _ = strconv.Atoi(envChatId)
	}

	if arguments.Message == "" {
		envMessage := os.Getenv("TELEGRAM_NOTIFIER_MESSAGE")
		if envMessage == "" {
			fmt.Println("Please set message. Example: -message=\"Hello, message!\"")
			os.Exit(1)
		}
		arguments.Message = envMessage
	}
}
