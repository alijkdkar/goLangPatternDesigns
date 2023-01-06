package main

import "fmt"


type excutionFunc func() error

type Step struct {
	nextStep *Step
	myFunc   excutionFunc
}

func NewStep(e excutionFunc) *Step {
	return &Step{
		nextStep: nil,
		myFunc:   e,
	}
}

func (st *Step) setNextStep(newNextStep *Step) {
	st.nextStep = newNextStep
}
func (stp *Step) Execute() {

	if stp == nil {
		panic("Nothing to exceute")
	}

	if stp.myFunc == nil {
		fmt.Println("no func")
	}
	if err := stp.myFunc(); err != nil {
		fmt.Println("execution Function occerd error")
	}
	if stp.nextStep != nil {
		stp.nextStep.Execute()
	}

}

type ChainBuilder struct {
	FirstStep, LastStep *Step
}

func newChainBuilder() *ChainBuilder {
	return &ChainBuilder{
		FirstStep: nil,
		LastStep:  nil,
	}
}

func (c *ChainBuilder) SetNextStep(stp *Step) *ChainBuilder {
	if c.FirstStep == nil {
		c.FirstStep = stp
	}

	if c.LastStep != nil {
		c.LastStep.setNextStep(stp)
	}
	c.LastStep = stp

	return c
}

func (c *ChainBuilder) build() *Step {
	return c.FirstStep
}

func f1() error {
	fmt.Println("f1")
	return nil
}
func f2() error {
	fmt.Println("f2")
	return nil
}
func f3() error {
	fmt.Println("f3")
	return nil
}

func main() {

	c := newChainBuilder()
	firstStep := c.SetNextStep(NewStep(f1)).SetNextStep(NewStep(f2)).SetNextStep(NewStep(f3)).build()
	firstStep.Execute()
}
