package Patterns

import "fmt"

type State interface {
	FirstOperation() error
	SecondOperation() error
	ThirdOperation() error
}

type Context struct {
	FirstState   State
	SecondState  State
	CurrentState State
}

func NewContext() *Context {
	c := &Context{}

	fs := &FirstState{
		Context: c,
	}

	ss := &SecondState{
		Context: c,
	}

	c.SetState(fs)
	c.FirstState = fs
	c.SecondState = ss
	return c
}

func (c *Context) FirstOperation() error {
	return c.CurrentState.FirstOperation()
}

func (c *Context) SecondOperation() error {
	return c.CurrentState.SecondOperation()
}

func (c *Context) ThirdOperation() error {
	return c.CurrentState.ThirdOperation()
}

func (c *Context) SetState(s State) {
	c.CurrentState = s
}

type FirstState struct {
	Context *Context
}

func (firstS *FirstState) FirstOperation() error {
	fmt.Println("acces: First command ")
	return nil
}

func (firstS *FirstState) SecondOperation() error {
	fmt.Println("block: Second command ")

	return nil
}

func (firstS *FirstState) ThirdOperation() error {
	fmt.Println("State 1 -> State 2")
	firstS.Context.SetState(firstS.Context.SecondState)
	return nil
}

type SecondState struct {
	Context *Context
}

func (secondS *SecondState) FirstOperation() error {
	fmt.Println("block: First command ")
	return nil
}

func (secondS *SecondState) SecondOperation() error {
	fmt.Println("acces: Second command ")

	return nil
}

func (secondS *SecondState) ThirdOperation() error {

	return nil
}
