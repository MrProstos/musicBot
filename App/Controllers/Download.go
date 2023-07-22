package Controllers

import (
	"MrProstos/download_utils/App/Models"
	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
	"os/exec"
)

func DownloadMusic(url string) (string, error) {
	client := new(youtube.Client)

	video, err := client.GetVideo(url)
	if err != nil {
		return "", nil
	}

	formats := video.Formats.WithAudioChannels()
	url = formats[0].URL

	fileName := uuid.New().String()
	exec.Command("ffmpeg", "-i", url, Models.AudioPath+fileName)

	return fileName, nil
}
