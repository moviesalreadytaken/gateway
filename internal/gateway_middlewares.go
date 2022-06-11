package internal

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RedirectDetector() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Next()
		status := g.Writer.Status()
		if status >= 300 && status < 400 {
			log.Println("HTTP REDIRECT DETECTED, request =", g.Request) //fixme
		}
	}
}

type DurationMeter struct {
	AvgServiceResponseDuration map[string]time.Duration
}

func NewDurationMeter() *DurationMeter {
	return &DurationMeter{
		AvgServiceResponseDuration: make(map[string]time.Duration),
	}
}

func (dm *DurationMeter) Middleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		t1 := time.Now()
		g.Next()
		duration := time.Since(t1)
		path := g.Request.URL.Path
		val, ok := dm.AvgServiceResponseDuration[path]
		if ok {
			dm.AvgServiceResponseDuration[path] = (val + duration) / 2
		} else {
			dm.AvgServiceResponseDuration[path] = duration
		}
	}
}
