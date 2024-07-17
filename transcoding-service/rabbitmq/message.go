package rabbitmq

type Message struct {
	VideoPath string `json:"video_path"`
	Title     string `json:"title"`
	Data      string `json:"data"`
}

type Response struct {
	Message    string `json:"message"`
	OutputPath string `json:"output_path"`
}
