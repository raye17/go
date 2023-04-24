package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	l := zap.NewExample()
	l.WithOptions(zap.Development(), zap.Fields(zap.Time("time:", time.Now()))).Warn("msg:warn")
}
