package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg string
}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("<-<-<-<-<-<-<-<-<-<-<-<-")
		i.msg = "left Pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("->->->->->->->->->->->->")
		i.msg = "right Pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("-----------------------")
		i.msg = "space Pressed"
	}
}
