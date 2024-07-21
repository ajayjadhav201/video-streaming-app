package transcoder

import (
	"log"
	"os/exec"
	"transcoding-service/rabbitmq"
	"transcoding-service/utils"
)

var (
	nfsBaseDir string = "/mnt/nfs/hls"
)

func TranscodeVideo(message *rabbitmq.Message) error {
	//
	if message.VideoPath == "" {
		return utils.Error("Video path not found")
	}
	err := convertToHLS(message.VideoPath, message.Title)
	return err
}

func convertToHLS(videoPath string, title string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", nfsBaseDir+"/"+title+".m3u8")
	err := cmd.Run()
	if err != nil {
		// log.Fatalf("Failed to convert video: %v", err)
		return err
	}
	log.Printf("Successfully converted %s to HLS", videoPath)
	return nil
}
