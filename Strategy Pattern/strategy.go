package main

import "fmt"

// --- Strategy Interfaces ---

type WalkableRobot interface {
	Walk()
}

type TalkableRobot interface {
	Talk()
}

type FlyableRobot interface {
	Fly()
}

// --- Concrete Strategies for Walk ---

type NormalWalk struct{}

func (n NormalWalk) Walk() {
	fmt.Println("Walking normally...")
}

type NoWalk struct{}

func (n NoWalk) Walk() {
	fmt.Println("Cannot walk.")
}

// --- Concrete Strategies for Talk ---

type NormalTalk struct{}

func (n NormalTalk) Talk() {
	fmt.Println("Talking normally...")
}

type NoTalk struct{}

func (n NoTalk) Talk() {
	fmt.Println("Cannot talk.")
}

// --- Concrete Strategies for Fly ---

type NormalFly struct{}

func (n NormalFly) Fly() {
	fmt.Println("Flying normally...")
}

type NoFly struct{}

func (n NoFly) Fly() {
	fmt.Println("Cannot fly.")
}

// --- Robot Base Struct (Composition of strategies) ---

type Robot struct {
	walkBehavior WalkableRobot
	talkBehavior TalkableRobot
	flyBehavior  FlyableRobot
}

func (r *Robot) Walk() {
	r.walkBehavior.Walk()
}

func (r *Robot) Talk() {
	r.talkBehavior.Talk()
}

func (r *Robot) Fly() {
	r.flyBehavior.Fly()
}

// --- Abstract-like behavior (using interface in Go) ---

type Projectable interface {
	Projection()
}

// --- Concrete Robots ---

type CompanionRobot struct {
	Robot
}

func NewCompanionRobot(w WalkableRobot, t TalkableRobot, f FlyableRobot) *CompanionRobot {
	return &CompanionRobot{
		Robot{
			walkBehavior: w,
			talkBehavior: t,
			flyBehavior:  f,
		},
	}
}

func (c *CompanionRobot) Projection() {
	fmt.Println("Displaying friendly companion features...")
}

type WorkerRobot struct {
	Robot
}

func NewWorkerRobot(w WalkableRobot, t TalkableRobot, f FlyableRobot) *WorkerRobot {
	return &WorkerRobot{
		Robot{
			walkBehavior: w,
			talkBehavior: t,
			flyBehavior:  f,
		},
	}
}

func (w *WorkerRobot) Projection() {
	fmt.Println("Displaying worker efficiency stats...")
}

func main() {
	robot1 := NewCompanionRobot(NormalWalk{}, NormalTalk{}, NoFly{})
	robot1.Walk()
	robot1.Talk()
	robot1.Fly()
	robot1.Projection()

	fmt.Println("--------------------")

	robot2 := NewWorkerRobot(NoWalk{}, NoTalk{}, NormalFly{})
	robot2.Walk()
	robot2.Talk()
	robot2.Fly()
	robot2.Projection()
}
