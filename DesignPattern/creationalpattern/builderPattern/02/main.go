package main

import "fmt"

type Leek struct {
}
type PeasantInter interface {
	Water() error
	Fertilization() error
	Reap() *Leek
}
type Peasant struct {
}

func (p *Peasant) Water() error {
	fmt.Println("浇水")
	return nil
}
func (p *Peasant) Fertilization() error {
	fmt.Println("施肥")
	return nil
}
func (p *Peasant) Reap() *Leek {
	fmt.Println("收割")
	return &Leek{}
}

type Village struct {
	peasant Peasant
}

func (v *Village) SetF(p Peasant) {
	v.peasant = p
}
func (v *Village) GetP() *Leek {
	return v.peasant.Reap()
}
