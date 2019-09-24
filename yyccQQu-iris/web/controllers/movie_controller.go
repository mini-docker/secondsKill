package controllers

import (
	"github.com/kataras/iris/mvc"
	"secondsKill/yyccQQu-iris/repositories"
	"secondsKill/yyccQQu-iris/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {

	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieRepository)
	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}

}
