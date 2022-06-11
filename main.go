package main

import (
	"log"

	ginpkg "github.com/gin-gonic/gin"

	"github.com/moviesalreadytaken/gateway/internal"
)

func main() {
	gin := ginpkg.New()
	cnf := internal.LoadCnfFromEnv()
	dm := internal.NewDurationMeter()
	controller, err := internal.NewGatewayController(cnf, dm)
	if err != nil {
		log.Fatalf("error while gateway controller initialization = %s", err.Error())
	}
	gin.Use(internal.RedirectDetector())
	internal.AddRoutesV1(gin, controller, dm)
	log.Fatal(gin.Run(":10000"))
}
