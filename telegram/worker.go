package telegram

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (s *service) RunNotificationWorker() {
	for {
		users, err := s.tm.AllUser()
		if err != nil {
			// Handle errors
		}

		for _, user := range users {
			s.bot.Send(&tb.User{ID: user.ChatID}, "This is a notification")
		}

		time.Sleep(time.Second * 5)
	}
}
