package context

import (
	"errors"
)

var Cache Service

func init() {
	Cache = New()
}

type Service interface {
	AddContext(userId int, ctx *Context)
	GetUserContext(userId int) (*Context, error)
	ResolveStep(userId int, response string) string
	HasPendingContext(userId int) bool
	SkipStep(userId int) string
	Cancel(userId int)
}

type service struct {
	cache map[int]Context
}

func New() Service {
	return &service{
		cache: make(map[int]Context),
	}
}

func (s *service) AddContext(userId int, ctx *Context) {
	s.cache[userId] = *ctx
}

func (s *service) GetUserContext(userId int) (*Context, error) {
	if val, ok := s.cache[userId]; ok {
		return &val, nil
	}

	return &Context{}, errors.New("Could not find user context")
}

func (s *service) HasPendingContext(userId int) bool {
	if _, ok := s.cache[userId]; ok {
		return true
	}

	return false
}

func (s *service) ResolveStep(userId int, response string) string {
	ctx := s.cache[userId]

	step := ctx.GetCurrentStep()
	step.Response = response
	ctx.IncrementStep()

	s.cache[userId] = ctx

	if ctx.IsComplete() {
		resp, err := ctx.Process(&ctx)
		if err != nil {
			return err.Error()
		}
		defer s.Cancel(userId)

		return resp
	}

	return ctx.GetCurrentQuestion()
}

func (s *service) SkipStep(userId int) string {
	ctx := s.cache[userId]
	ctx.IncrementStep()

	s.cache[userId] = ctx

	return ctx.GetCurrentQuestion()
}

func (s *service) Cancel(userId int) {
	delete(s.cache, userId)
}

// {
// 	"1234": {
// 		"context": "task-create",
// 		"confirmation": "I added the task with id NEX-1234"
// 		"steps": [
// 			{
// 				"response": "What is the title of the task?"
// 				"payload": "Set up a catch up with Ric"
// 			},
// 			{
// 				"response": "When do you want to schedule this?"
// 				"payload": "20-07-2020"
// 			},
// 			{
// 				"response": "When time?"
// 				"payload": "20-07-2020"
// 			},
// 			{
// 				"response": "Do you want to add any notes?"
// 				"payload": "tell Ric that he must complete his task for this sprint"
// 			},
// 			{
// 				"response": "what is the priority for this task?"
// 				"payload": "1"
// 			},
// 		]
// 	}
// }
