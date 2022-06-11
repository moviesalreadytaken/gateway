package internal

import (
	"net/http"
	"net/http/httputil"
	urlpkg "net/url"

	"github.com/gin-gonic/gin"
)

type GatewayController struct {
	dm                       *DurationMeter
	userServiceUrl           *urlpkg.URL
	movieServiceUrl          *urlpkg.URL
	recommendationServiceUrl *urlpkg.URL
}

func NewGatewayController(cnf *AppConfig, dm *DurationMeter) (*GatewayController, error) {
	usersUrl, err := urlpkg.Parse(cnf.UsersServiceUrl)
	if err != nil {
		return nil, err
	}
	moviesUrl, err := urlpkg.Parse(cnf.MoviesServiceUrl)
	if err != nil {
		return nil, err
	}
	recomUrl, err := urlpkg.Parse(cnf.RecommendationServiceUrl)
	if err != nil {
		return nil, err
	}
	return &GatewayController{
		dm:                       dm,
		userServiceUrl:           usersUrl,
		movieServiceUrl:          moviesUrl,
		recommendationServiceUrl: recomUrl,
	}, nil
}

func (c *GatewayController) HandleMovies(g *gin.Context) {
	newProxy(c.movieServiceUrl, g)
}

func (c *GatewayController) HandleUsers(g *gin.Context) {
	newProxy(c.userServiceUrl, g)
}

func (c *GatewayController) HandleRecommendations(g *gin.Context) {
	newProxy(c.recommendationServiceUrl, g)
}

func (c *GatewayController) avgRoutesExecutionTime(g *gin.Context) {
	res := make(map[string]float64)
	for k, v := range c.dm.AvgServiceResponseDuration {
		res[k] = v.Seconds()
	}
	g.JSON(http.StatusOK, res)
}

func newProxy(url *urlpkg.URL, g *gin.Context) {
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		req.Header = g.Request.Header
		req.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = g.Param("path")
		req.Body = g.Request.Body
	}
	proxy.ServeHTTP(g.Writer, g.Request)
}
