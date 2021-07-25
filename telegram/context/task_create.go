package context

import (
	"github.com/MihaiBlebea/task-manager/domain"
)

func TaskCreateContext(tm domain.TaskManager) (*Context, error) {
	// projects, err := tm.GetUserProjects(1)
	// if err != nil {
	// 	return &Context{}, err
	// }

	// ctx := Context{
	// 	Label: "task-create",
	// 	Steps: []Step{
	// 		{
	// 			Question: func() string { return "What is the title of the task?" },
	// 		},
	// 		{
	// 			Question: func() string {
	// 				question := "What project should this sit under?\n"
	// 				for _, project := range projects {
	// 					question += fmt.Sprintf("%d %s\n", project.ID, project.Title)
	// 				}

	// 				return question
	// 			},
	// 		},
	// 		{
	// 			Question: func() string { return "What is the expire datetime?" },
	// 		},
	// 		{
	// 			Question: func() string { return "Do you want to add some notes?" },
	// 		},
	// 		{
	// 			Question: func() string { return "What is the priority of this task?" },
	// 		},
	// 	},
	// }

	// ctx.Process = func(c *Context) (string, error) {
	// 	title := c.GetStep(0).Response
	// 	projectName := c.GetStep(1).Response
	// 	expireAt, err := time.Parse("2006-01-02 16:00", c.GetStep(2).Response)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	notes := c.GetStep(3).Response
	// 	priority, err := strconv.Atoi(c.GetStep(4).Response)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	for _, proj := range projects {
	// 		if strings.ToLower(proj.Title) == strings.ToLower(projectName) {
	// 			id, err := tm.CreateTask(1, 0, proj.ID, title, notes, expireAt.Format("2006-01-02T15:04:05.000Z"), false, 1, "", priority)
	// 			if err != nil {
	// 				return "", err
	// 			}

	// 			return fmt.Sprintf("Task with id %d was created", id), nil
	// 		}
	// 	}

	// 	return "", errors.New("Could not find project by name")
	// }

	return &Context{}, nil
}
