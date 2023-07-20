package main

import "MrProstos/download_utils/models/database"

//const (
//	HELP_COMMAND     string = "/help"
//	PLAYLIST_COMMAND string = "/playlist"
//	SEARCH_COMMAND   string = "/search"
//)

//var commands = tgbotapi.NewSetMyCommands(
//	tgbotapi.BotCommand{
//		Command:     "/help",
//		Description: "Помощь",
//	},
//	tgbotapi.BotCommand{
//		Command:     "/playlist",
//		Description: "Показать списко треков",
//	},
//	tgbotapi.BotCommand{
//		Command:     "/search",
//		Description: "Поиск трека",
//	},
//)

func main() {
	database.GetDB()
	//bot, err := tgbotapi.NewBotAPI(utils.GetEnv("TELEGRAM_TOKEN"))
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//bot.Debug = true
	//
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//u := tgbotapi.NewUpdate(-1)
	//u.Timeout = 60
	//
	//updates := bot.GetUpdatesChan(u)
	//tgbotapi.NewRemoveKeyboard(true)
	//_, err = bot.Request(commands)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//for update := range updates {
	//	if update.Message != nil {
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//
	//		switch update.Message.Command() {
	//		case HELP_COMMAND:
	//			msg.Text = "I understand /sayhi and /status."
	//		case SEARCH_COMMAND:
	//			msg.Text = "playlist"
	//		case PLAYLIST_COMMAND:
	//			msg.Text = "search"
	//		default:
	//			msg.Text = "I don't know that command"
	//		}
	//
	//		if _, err := bot.Send(msg); err != nil {
	//			log.Panic(err)
	//		}
	//	}
	//}
}
