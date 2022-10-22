package main

import "fmt"

type Run interface {
	Running()
}

type Swim interface {
	Swimming()
}

type Sport interface {
	Run
	Swim
}

func GoSport(s Sport) {
	s.Running()
	s.Swimming()
}

type BOY struct {
	Name string
}

func (b *BOY) Running() {
	fmt.Println("runing1")
}

type GIRL struct {
	Name string
}

func (b *BOY) Swimming() {
	fmt.Println("swimming2")
}

/*
type Spo struct {
	BOY
	GIRL
}*/

/*function (s *Spo) Running() {
	fmt.Println("Spo 33")
}*/

func main() {

	//s := Spo{}
	b := BOY{}
	GoSport(&b)
}
