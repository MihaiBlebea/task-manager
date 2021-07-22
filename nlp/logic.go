package nlp

import (
	"io"
	"os"

	witai "github.com/wit-ai/wit-go/v2"
)

type Service interface {
	TextQuery(message string) (*witai.MessageResponse, error)
	SpeechQuery(file io.Reader) (*witai.MessageResponse, error)
}

type service struct {
	client *witai.Client
}

func New() Service {
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	return &service{client}
}

func (s *service) TextQuery(message string) (*witai.MessageResponse, error) {
	return s.client.Parse(&witai.MessageRequest{
		Query: message,
	})
}

func (s *service) SpeechQuery(file io.Reader) (*witai.MessageResponse, error) {
	return s.client.Speech(&witai.MessageRequest{
		Speech: &witai.Speech{
			File:        file,
			ContentType: "audio/raw;encoding=unsigned-integer;bits=16;rate=8000;endian=big",
		},
	})
}
