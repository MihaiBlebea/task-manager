package nlp

import witai "github.com/wit-ai/wit-go/v2"

func (s *service) Seed() error {
	s.client.CreateIntent("task_remove")

	s.client.CreateEntity(witai.Entity{
		Name: "task_slug",
	})

	s.client.TrainUtterances([]witai.Training{
		witai.Training{
			Text:   "Remove the task NEX-1234",
			Intent: "task_remove",
			Entities: []witai.TrainingEntity{
				witai.TrainingEntity{
					Entity:   "task_slug",
					Start:    15,
					End:      23,
					Body:     "NEX-1234",
					Entities: []witai.TrainingEntity{},
				},
			},
			Traits: []witai.TrainingTrait{},
		},
	})

	return nil
}
