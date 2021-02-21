package fsm

import (
	"fmt"
	"log"
	"testing"
)

func TestFSM(t *testing.T) {
	//  模拟转门
	const (
		// 状态
		Locked   = "Locked"
		Unlocked = "Unlocked"

		//  事件
		Push = "push"
		Coin = "Coin"
	)

	f := FSM{
		Transitions: []Transition{
			{From: Locked, Event: Coin, To: Unlocked},
			{From: Unlocked, Event: Push, To: Locked},
		},
		Callbacks: []Callback{
			OnEntry(Unlocked, func(t *Transition, _ ...interface{}) {
				log.Println("进入了Unlocked状态")
			}),
			OnExit(Locked, func(t *Transition, _ ...interface{}) {
				log.Println("离开了locked状态")
			}),
			OnEntry(Locked, func(t *Transition, _ ...interface{}) {
				log.Println("进入了Locked状态")
			}),
			OnExit(Unlocked, func(t *Transition, _ ...interface{}) {
				log.Println("离开了Unlocked状态")
			}),
			OnXXXEvent(Push, func(t *Transition, _ ...interface{}) {
				log.Println("发生了Push事件")
			}),
			OnXXXEvent(Coin, func(t *Transition, _ ...interface{}) {
				log.Println("发生了Coin事件")
			}),
		},
	}
	state := Locked
	state = f.Trigger(state, Coin)
	state = f.Trigger(state, Push)
	log.Panicln(state)

}

func TestPtr(t *testing.T) {
	s1 := struct {
		Name string
		Age  int
	}(struct {
		Name string
		Age  int
	}{Name: "xm", Age: 11})
	fmt.Println(s1.Name)
}
