package internal

import (
	"fmt"
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
		methodPath := fmt.Sprintf("[%s]=%s",g.Request.Method,g.Request.URL.Path)
		oldDuration, ok := dm.AvgServiceResponseDuration[methodPath]
		if ok {
			dm.AvgServiceResponseDuration[methodPath] = (oldDuration + duration) / 2
		} else {
			dm.AvgServiceResponseDuration[methodPath] = duration
		}
	}
}
