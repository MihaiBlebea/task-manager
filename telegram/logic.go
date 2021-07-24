package telegram

import (
	"fmt"
	"os"
	"time"

	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/MihaiBlebea/task-manager/nlp"
	"github.com/MihaiBlebea/task-manager/telegram/context"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Service interface {
	Start()
}

type service struct {
	bot *tb.Bot
	nlp nlp.Service
	tm  domain.TaskManager
}

func New(nlp nlp.Service, tm domain.TaskManager) (Service, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 30 * time.Second},
	})
	if err != nil {
		return &service{}, err
	}

	// Add the commands
	serv := &service{
		bot: bot,
		nlp: nlp,
		tm:  tm,
	}

	return serv, nil
}

func (s *service) Start() {
	s.bot.Handle("/hello", func(m *tb.Message) {
		fmt.Println(m.Text)
		s.bot.Send(m.Sender, fmt.Sprintf("%+v", m))
	})

	s.bot.Handle(onTask, func(m *tb.Message) {
		context.Cache.Cancel(int(m.Chat.ID))

		ctx, err := context.TaskCreateContext(s.tm)
		if err != nil {
			s.bot.Send(m.Sender, err.Error())
		}
		context.Cache.AddContext(int(m.Chat.ID), ctx)
		s.bot.Send(m.Sender, ctx.GetCurrentQuestion())
	})

	s.bot.Handle(onSkip, func(m *tb.Message) {
		resp := context.Cache.SkipStep(int(m.Chat.ID))
		s.bot.Send(m.Sender, resp)
	})

	s.bot.Handle(onCancel, func(m *tb.Message) {
		context.Cache.Cancel(int(m.Chat.ID))
		s.bot.Send(m.Sender, "Context was cancelled")
	})

	s.bot.Handle(onTasks, func(m *tb.Message) {
		context.Cache.Cancel(int(m.Chat.ID))

		ctx, err := context.TasksContext(s.tm)
		if err != nil {
			s.bot.Send(m.Sender, err.Error())
		}
		context.Cache.AddContext(int(m.Chat.ID), ctx)
		s.bot.Send(m.Sender, ctx.GetCurrentQuestion())
	})

	s.bot.Handle(tb.OnText, func(m *tb.Message) {
		chatId := int(m.Chat.ID)
		if context.Cache.HasPendingContext(chatId) == false {
			s.bot.Send(m.Sender, "This is a text message")
			return
		}
		resp := context.Cache.ResolveStep(chatId, m.Text)

		s.bot.Send(m.Sender, resp)
	})

	s.bot.Handle(tb.OnAudio, func(m *tb.Message) {
		s.bot.Send(m.Sender, "This is an audio file")
	})

	s.bot.Handle(tb.OnVoice, func(m *tb.Message) {
		path, err := s.bot.FileURLByID(m.Voice.FileID)
		if err != nil {
			s.bot.Send(m.Sender, fmt.Sprintf("%+v", err))
			return
		}

		tmpPath, err := downloadFileTmp(path)
		if err != nil {
			s.bot.Send(m.Sender, fmt.Sprintf("%+v", err))
			return
		}

		mp3Path, err := oggToMp3(tmpPath)
		if err != nil {
			s.bot.Send(m.Sender, fmt.Sprintf("%+v", err))
			return
		}

		f, err := os.Open(mp3Path)
		if err != nil {
			s.bot.Send(m.Sender, fmt.Sprintf("%+v", err))
			return
		}

		msg, err := s.nlp.SpeechQuery(f)
		if err != nil {
			s.bot.Send(m.Sender, fmt.Sprintf("%+v", err))
			return
		}

		os.RemoveAll("/tmp")

		s.bot.Send(m.Sender, msg.Text)
		s.bot.Send(m.Sender, fmt.Sprintf("%+v", msg.Entities))
	})

	s.bot.Start()
}
