package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/MihaiBlebea/task-manager/db"
	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/MihaiBlebea/task-manager/domain/task"
	"github.com/MihaiBlebea/task-manager/domain/user"
	"github.com/MihaiBlebea/task-manager/nlp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	tb "gopkg.in/tucnak/telebot.v2"
)

func init() {
	rootCmd.AddCommand(telegramCmd)
}

var telegramCmd = &cobra.Command{
	Use:   "telegram",
	Short: "Start the telegram application.",
	Long:  "Start the telegram application.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		db, err := db.ConnectSQL()
		if err != nil {
			return err
		}
		db.AutoMigrate(&user.User{}, &task.Task{}, &project.Project{})

		nlp := nlp.New()

		b, err := tb.NewBot(tb.Settings{
			Token:  os.Getenv("TELEGRAM_TOKEN"),
			Poller: &tb.LongPoller{Timeout: 30 * time.Second},
		})

		if err != nil {
			return err
		}

		b.Handle("/hello", func(m *tb.Message) {
			b.Send(m.Sender, "Hello World!")
		})

		b.Handle("/task-create", func(m *tb.Message) {
			b.Send(m.Sender, "Let's create a new task")
		})

		b.Handle("/task-complete", func(m *tb.Message) {
			b.Send(m.Sender, "Let's complete a task")
		})

		b.Handle(tb.OnText, func(m *tb.Message) {
			b.Send(m.Sender, "This is a text message")
		})

		b.Handle(tb.OnAudio, func(m *tb.Message) {
			b.Send(m.Sender, "This is an audio file")
		})

		b.Handle(tb.OnVoice, func(m *tb.Message) {
			// b.Send(m.Sender, fmt.Sprintf("%+v", m.Voice.File))
			// b.Send(m.Sender, fmt.Sprintf("%+v", m.Voice.File.FileURL))

			path, err := b.FileURLByID(m.Voice.FileID)
			if err != nil {
				b.Send(m.Sender, fmt.Sprintf("%+v", err))
				return
			}

			r := downloadFile(path)

			msg, err := nlp.SpeechQuery(r)
			if err != nil {
				b.Send(m.Sender, fmt.Sprintf("%+v", err))
				return
			}

			b.Send(m.Sender, fmt.Sprintf("ceva ceva %+v", msg.Text))
		})

		b.Start()

		return nil
	},
}

func downloadFile(url string) io.Reader {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return bytes.NewReader(b)
}
