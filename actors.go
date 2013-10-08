package main

import "fmt"

type Actor interface {
	Character()
	Act()
}

type SpectrumActor interface {
	Actor
	FearAct()
}

func Introduce(a Actor) {
	a.Character()
}

func Perform(a Actor) {
	a.Character()
	fmt.Printf(": ")
	a.Act()
}

func BeAfraid(a SpectrumActor) {
	a.Character()
	fmt.Printf(": ")
	a.FearAct()
}

type SpeechActor struct {
	characterName string
	monologue     string
	afraid        string
}

func (s SpeechActor) Character() {
	fmt.Printf(s.characterName)
}

func (s SpeechActor) Act() {
	fmt.Println(s.monologue)
}

func (s SpeechActor) FearAct() {
	fmt.Println(s.afraid)
}

type DieingActor struct {
	characterName string
	deathScene    string
}

func (d DieingActor) Character() {
	fmt.Printf(d.characterName)
}

func (d DieingActor) Act() {
	fmt.Println(d.deathScene)
}

func play() {
	Friend := DieingActor{"Friend", "Blergh"}
	Hamlet := SpeechActor{"Hamlet", "To be or not to be... That is the question", "Oh no!"}
	Perform(Hamlet)
	Perform(Friend)
	BeAfraid(Hamlet)
}