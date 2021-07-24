package context

type Context struct {
	Label        string
	CurrentStep  int
	Steps        []Step
	Process      func(c *Context) (string, error)
	Confirmation string
}

type Step struct {
	Question func() string
	Response string
}

func (c *Context) GetCurrentStep() *Step {
	return &c.Steps[c.CurrentStep]
}

func (c *Context) GetStep(index int) *Step {
	return &c.Steps[index]
}

func (c *Context) IncrementStep() {
	c.CurrentStep++
}

func (c *Context) GetCurrentQuestion() string {
	return c.GetCurrentStep().Question()
}

func (c *Context) IsComplete() bool {
	return c.CurrentStep == len(c.Steps)
}
