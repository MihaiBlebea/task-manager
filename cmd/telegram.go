package cmd

import (
	"os"

	"github.com/MihaiBlebea/task-manager/db"
	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/MihaiBlebea/task-manager/domain/user"
	"github.com/MihaiBlebea/task-manager/nlp"
	"github.com/MihaiBlebea/task-manager/telegram"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
		db.AutoMigrate(&user.User{})

		tm := domain.New(db)

		nlp := nlp.New()

		tel, err := telegram.New(nlp, tm)
		if err != nil {
			return err
		}

		tel.Start()

		return nil
	},
}
