package repository

import (
	"github.com/MrProstos/musicBot/internal/models"
	"github.com/MrProstos/musicBot/pkg/psql"
	"github.com/google/uuid"
	"log"
	"os/exec"
)

type Repository struct {
	*psql.Psql
}

func New(db *psql.Psql) *Repository {
	return &Repository{db}
}

func (repo *Repository) GetPlaylistByUserId(userId uint) *models.Playlist {
	playlist := &models.Playlist{}

	result := repo.Where("user_id = ?", userId).First(playlist)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	return playlist
}

func (repo *Repository) CreatePlaylist(userId uint) *models.Playlist {
	playlist := &models.Playlist{UserId: userId}
	repo.Create(playlist)

	return playlist
}

func (repo *Repository) GetAudioFileById(fileId string) *models.AudioStorage {
	audioStorage := &models.AudioStorage{}

	result := repo.Where("file_id = ?", fileId).First(audioStorage)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	return audioStorage
}

func (repo *Repository) StoreAudioFileFromYoutube(audioStorage *models.AudioStorage) *models.AudioStorage {
	filePath := models.AudioPath + uuid.New().String()
	cmd := exec.Command("ffmpeg", "-y", "-i", audioStorage.FilePath, "-f", "mp3", filePath)
	cmd.Run()

	audioStorage.FilePath = filePath

	repo.Create(audioStorage)

	return audioStorage
}
