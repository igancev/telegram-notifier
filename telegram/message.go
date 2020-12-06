package telegram

type Message struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func (this Message) New(chatId int, text string) Message {
	this.ParseMode = "MarkdownV2"
	this.ChatID = chatId
	this.Text = text
	return this
}
