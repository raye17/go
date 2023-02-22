package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
}

func NewShip() *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./images/ship.png")
	if err != nil {
		log.Fatalf("load image failed,err:%v\n", err)
	}
	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
	}
	return ship
}
func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	screen.DrawImage(ship.image, op)
}
