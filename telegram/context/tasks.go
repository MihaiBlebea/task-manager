package context

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MihaiBlebea/task-manager/domain"
)

func TasksContext(tm domain.TaskManager) (*Context, error) {
	projects, err := tm.GetUserProjects(1)
	if err != nil {
		return &Context{}, err
	}

	ctx := Context{
		Label: "tasks",
		Steps: []Step{
			{
				Question: func() string {
					question := "Chose a project:\n"
					for _, project := range projects {
						question += fmt.Sprintf("%d %s\n", project.ID, project.Title)
					}

					return question
				},
			},
		},
	}

	ctx.Process = func(c *Context) (string, error) {
		projectName := c.GetStep(0).Response
		if err != nil {
			return "", err
		}

		for _, proj := range projects {
			if strings.ToLower(proj.Title) == strings.ToLower(projectName) {
				response := "Here are your tasks:\n"
				for _, task := range proj.Tasks {
					response += fmt.Sprintf("%d %s\n", task.ID, task.Title)
				}

				return response, nil
			}
		}

		return "", errors.New("Sorry could not find any project with that name")
	}

	return &ctx, nil
}
