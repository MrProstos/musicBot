package usecase

import (
	"github.com/MrProstos/musicBot/config"
	"github.com/MrProstos/musicBot/internal/models"
	"github.com/MrProstos/musicBot/internal/usecase/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kkdai/youtube/v2"
	"log"
	"regexp"
)

const (
	HelpCommand     string = "help"
	PlaylistCommand string = "playlist"
)

// MusicBot telegram bot realization
type MusicBot struct {
	*tgbotapi.BotAPI

	DB *repository.Repository
}

// NewMusicBot create new MusicBot instance
func NewMusicBot(cfg config.App) (*MusicBot, error) {
	newBot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	newBot.Debug = cfg.Debug

	return &MusicBot{BotAPI: newBot}, nil
}

// StartListeningUpdates start listening for updates
func (mBot *MusicBot) StartListeningUpdates() {
	mBot.setCommands()

	updateConfig := tgbotapi.NewUpdate(-1)
	updateConfig.Timeout = 60

	updates := mBot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		var msg tgbotapi.Chattable

		if update.Message.IsCommand() {
			mBot.CommandController(update, &msg)
		} else {
			mBot.TextController(update, &msg)
		}

		if msg != nil {
			_, err := mBot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// setCommands set telegram commands
func (mBot *MusicBot) setCommands() {
	mBot.Request(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     HelpCommand,
			Description: "Помощь",
		},
		tgbotapi.BotCommand{
			Command:     PlaylistCommand,
			Description: "Показать список треков",
		},
	))
}

// TextController text controller
func (mBot *MusicBot) TextController(update tgbotapi.Update, msg *tgbotapi.Chattable) {

	if mBot.isYoutubeUrl(update.Message.Text) {

		playlist := mBot.DB.GetPlaylistByUserId(uint(update.Message.From.ID))
		if playlist == nil {
			playlist = mBot.DB.CreatePlaylist(uint(update.Message.From.ID))
		}

		audioStorage := mBot.getAudioStorageFromYoutube(update.Message.Text)
		audioStorage.PlaylistId = playlist.Id

		// if the audio file is not in the storage, then download it
		existAudioStorage := mBot.DB.GetAudioFileById(audioStorage.FileId)
		if existAudioStorage == nil {
			audioStorage = mBot.DB.StoreAudioFileFromYoutube(audioStorage)
		} else {
			audioStorage.FilePath = existAudioStorage.FilePath
		}

		msgAudio := tgbotapi.NewAudio(update.Message.From.ID, tgbotapi.FilePath(audioStorage.FilePath))
		msgAudio.Title = audioStorage.Title
		msgAudio.Performer = audioStorage.Author

		*msg = msgAudio
	}
}

// getAudioStorageFromYoutube get models.AudioStorage by youtube url
func (mBot *MusicBot) getAudioStorageFromYoutube(videoUrl string) *models.AudioStorage {
	video, err := new(youtube.Client).GetVideo(videoUrl)
	if err != nil {
		return nil
	}

	formats := video.Formats.WithAudioChannels()
	return &models.AudioStorage{
		FileId:   video.ID,
		Title:    video.Title,
		Author:   video.Author,
		FilePath: formats[0].URL,
	}
}

// isYoutubeUrl is youtube url
func (mBot *MusicBot) isYoutubeUrl(url string) bool {
	regex := regexp.MustCompile("https:\\/\\/www\\.youtube\\.com\\/watch.+v=\\w+")

	result := regex.Find([]byte(url))
	if result == nil {
		return false
	}

	return true
}

func (mBot *MusicBot) CommandController(update tgbotapi.Update, msg *tgbotapi.Chattable) {

}
