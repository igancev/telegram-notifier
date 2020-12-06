# Telegram notifier

Simple console utility for send message to telegram chat from telegram bot.

## Create telegram bot

For create and manage Telegram bots you will need to communicate with [@botfather](https://t.me/botfather). Details [here](https://core.telegram.org/bots#3-how-do-i-create-a-bot).

## Install

```bash
git clone https://github.com/igancev/telegram-notifier
cd telegram-notifier
go build -o telegram-notifier .
sudo mv telegram-notifier /usr/local/bin
telegram-notifier --help
```

or use docker image

```
docker pull igancev/telegram-notifier:latest
```

## Usage

To send a message, you must provide 3 required arguments:

- `token`. Authorization api token for bot. Received from the [@botfather](https://t.me/botfather) when creating the bot

- `chat_id`. Unique telegram identifier for the target chat. If you want sending message yourself, you may detect his for example in [@userinfobot](https://t.me/userinfobot).

- `message`. Any message what you want to send. Supports telegram markdown syntax.

This parameters must be pass as arguments or define as environment variables:

- `TELEGRAM_NOTIFIER_TOKEN`
- `TELEGRAM_NOTIFIER_CHAT_ID`
- `TELEGRAM_NOTIFIER_MESSAGE`

Inline arguments is a higher priority than environment variables.

### Bin

```bash
telegram-notifier -token='telegramBotApiToken' -chat_id=123456 -message='*Any* _message_ with `markdown`'
```

### Docker

```bash
docker run --rm igancev/telegram-notifier telegram-notifier -token='telegramBotApiToken' -chat_id=123456 -message='*Any* _message_ with `markdown`'
```

### Docker compose

```yaml
version: '3'
services:
    telegram-notifier:
        image: igancev/telegram-notifier:latest
        environment:
          - TELEGRAM_NOTIFIER_TOKEN=1428033353:AAGg95hVqqW2qStmKlJHD0rXRG2vyMNVxlQ
          - TELEGRAM_NOTIFIER_CHAT_ID=123456
          - TELEGRAM_NOTIFIER_MESSAGE=*Any* _message_ with `markdown`
```

```bash
docker-compose up
```
