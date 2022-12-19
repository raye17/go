package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input *Input
	cfg   *Config
	ship  *Ship
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	return &Game{
		input: &Input{},
		ship:  NewShip(),
		cfg:   cfg,
	}
}
func (g *Game) Update() error {
	g.input.Update()
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2
}
