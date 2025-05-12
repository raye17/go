package main

type Fruit int

const (
	Apple Fruit = iota
	Banana
	Cherry
)
const (
	Oring Fruit = iota
	Grape
	Watermelon = iota + 1
)
