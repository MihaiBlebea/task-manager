package cmd

import (
	"fmt"
	"os"

	"github.com/MihaiBlebea/task-manager/api"
	handler "github.com/MihaiBlebea/task-manager/api/handlers"
	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/MihaiBlebea/task-manager/domain/task"
	"github.com/MihaiBlebea/task-manager/domain/user"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application server.",
	Long:  "Start the application server.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		// dsn := fmt.Sprintf(
		// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
		// 	os.Getenv("POSTGRES_HOST"),
		// 	os.Getenv("POSTGRES_USER"),
		// 	os.Getenv("POSTGRES_PASSWORD"),
		// 	os.Getenv("POSTGRES_DB"),
		// 	os.Getenv("POSTGRES_PORT"),
		// )
		db, err := connectSQL()
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

func connectSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
