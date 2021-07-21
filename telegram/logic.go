package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

type Service interface {
	ParseRequest(r *http.Request) (*Update, error)
	SendResponse(chatId int, text string) (string, error)
	HandleRequest(update *Update) error
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) ParseRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return nil, err
	}

	return &update, nil
}

func (s *service) SendResponse(chatId int, text string) (string, error) {
	telegramApi := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		os.Getenv("TELEGRAM_TOKEN"),
	)
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *service) HandleRequest(update *Update) error {
	isCmd := isCommand(update.Message.Text)

	if isCmd == false {
		s.SendResponse(update.Message.Chat.Id, "This is not a command")
		return nil
	}

	resp, err := handleCommand(update.Message.Text)
	if err != nil {
		return err
	}

	s.SendResponse(update.Message.Chat.Id, resp)

	return nil
}

func handleCommand(command string) (string, error) {
	cmd := TelegramCmd(command)

	switch cmd {
	case CreateTaskCmd:
		return "Create a new task", nil
	case CompleteTaskCmd:
		return "Complete a task", nil
	}

	return "", errors.New("This is not a valid command")
}

func isCommand(message string) bool {
	for _, cmd := range commands {
		if message == string(cmd) {
			return true
		}
	}

	return false
}

func whichCommand(message string) (*TelegramCmd, bool) {
	for _, cmd := range commands {
		if message == string(cmd) {
			return &cmd, true
		}
	}

	return nil, false
}
