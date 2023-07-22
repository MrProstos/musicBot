package TelegramBot

import (
	"MrProstos/download_utils/App/Controllers"
	"MrProstos/download_utils/App/Utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

const (
	HelpCommand     string = "help"
	PlaylistCommand string = "playlist"
	DownloadCommand string = "download"
)

// getBot get instance tgbotapi.BotAPI
func getBot() *tgbotapi.BotAPI {
	if tgBot == nil {
		initBot()
	}

	return tgBot
}

var tgBot *tgbotapi.BotAPI

// initBot Initialize telegram bot
func initBot() {
	bot, err := tgbotapi.NewBotAPI(Utils.GetEnv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug, _ = strconv.ParseBool(Utils.GetEnv("TELEGRAM_DEBUG"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	tgBot = bot
}

// Listening start listening updates
func Listening() {
	setCommands()

	u := tgbotapi.NewUpdate(-1)
	u.Timeout = 60

	for update := range getBot().GetUpdatesChan(u) {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ParseMode = tgbotapi.ModeHTML

			msg.Text = commandHandler(*update.Message).Text

			if _, err := getBot().Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

// setCommands Get telegram commands
func setCommands() {
	commands := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     HelpCommand,
			Description: "Помощь",
		},
		tgbotapi.BotCommand{
			Command:     PlaylistCommand,
			Description: "Показать список треков",
		},
		tgbotapi.BotCommand{
			Command:     DownloadCommand,
			Description: "Скачать трек",
		},
	)

	_, err := getBot().Request(commands)
	if err != nil {
		log.Fatalln(err)
	}
}

func commandHandler(message tgbotapi.Message) tgbotapi.Message {
	switch message.Command() {
	case HelpCommand:
	case DownloadCommand:
		Controllers.DownloadMusic(message.CommandArguments())
	case PlaylistCommand:
	default:
		message.Text = "Я не знаю такой команды"
	}

	return message
}
