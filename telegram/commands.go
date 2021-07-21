package telegram

type TelegramCmd string

const (
	ProjectsCmd      TelegramCmd = "/projects"
	CreateProjectCmd TelegramCmd = "/project-create"
	RemoveProjectCmd TelegramCmd = "/project-remove"
	CreateTaskCmd    TelegramCmd = "/task-create"
	RemoveTaskCmd    TelegramCmd = "/task-remove"
	CompleteTaskCmd  TelegramCmd = "/task-complete"
	TasksCmd         TelegramCmd = "/tasks"
)

var commands = []TelegramCmd{
	ProjectsCmd,
	CreateProjectCmd,
	RemoveProjectCmd,
	CreateTaskCmd,
	RemoveTaskCmd,
	CompleteTaskCmd,
	TasksCmd,
}
