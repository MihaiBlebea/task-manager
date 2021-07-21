package telegram

type TelegramCmd string

const (
	CreateTaskCmd   TelegramCmd = "/task-create"
	RemoveTaskCmd   TelegramCmd = "/task-remove"
	CompleteTaskCmd TelegramCmd = "/task-complete"
	TasksCmd        TelegramCmd = "/tasks"
)

var commands = []TelegramCmd{
	CreateTaskCmd,
	RemoveTaskCmd,
	CompleteTaskCmd,
	TasksCmd,
}
