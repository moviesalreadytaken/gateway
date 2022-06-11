package internal

import "github.com/gin-gonic/gin"

func AddRoutesV1(gin *gin.Engine, controller *GatewayController, dm *DurationMeter) {
	api := gin.Group("/api", dm.Middleware())
	{
		movies := api.Group("/movies-service")
		{
			movies.Any("/*path", controller.HandleMovies)
		}
		users := api.Group("/users-service")
		{
			users.Any("/*path", controller.HandleUsers)
		}
		recom := api.Group("/recomm-service")
		{
			recom.Any("/*path", controller.HandleRecommendations)
		}
	}
	stats := gin.Group("/stats")
	{
		stats.GET("/avg-time", controller.avgRoutesExecutionTime)
	}
}
