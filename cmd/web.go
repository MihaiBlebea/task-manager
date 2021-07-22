package cmd

import (
	"os"

	"github.com/MihaiBlebea/task-manager/api"
	handler "github.com/MihaiBlebea/task-manager/api/handlers"
	"github.com/MihaiBlebea/task-manager/db"
	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/MihaiBlebea/task-manager/domain/task"
	"github.com/MihaiBlebea/task-manager/domain/user"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server.",
	Long:  "Start the web server.",
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

		hand := handler.New(db)

		serv := api.New(hand, l)
		serv.Server()

		return nil
	},
}
