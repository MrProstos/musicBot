package repository

import (
	"github.com/MrProstos/musicBot/internal/models"
	"github.com/MrProstos/musicBot/pkg/psql"
	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
	"os/exec"
)

type Repository struct {
	psql.Psql
}

func New(db *psql.Psql) *Repository {
	return &Repository{*db}
}

func (repo *Repository) GetPlaylistByUserId(userId uint) *models.Playlist {
	var playlist *models.Playlist

	repo.First(playlist, "user_id = ?", userId)
	if playlist == nil {
		return nil
	}

	return playlist
}

func (repo *Repository) CreatePlaylist(userId uint) *models.Playlist {
	playlist := &models.Playlist{UserId: userId}
	repo.Create(playlist)

	return playlist
}

func (repo *Repository) StoreAudioFileFromYoutube(video *youtube.Video, playListId uint) *models.AudioStorage {
	formats := video.Formats.WithAudioChannels()
	filePath := models.AudioPath + uuid.New().String()

	cmd := exec.Command("ffmpeg", "-y", "-i", formats[0].URL, "-f", "mp3", filePath)
	cmd.Run()

	audioStorage := &models.AudioStorage{
		FilePath:   filePath,
		Title:      video.Title,
		Author:     video.Author,
		VideoId:    video.ID,
		PlaylistId: playListId,
	}

	repo.Create(audioStorage)

	return audioStorage
}
