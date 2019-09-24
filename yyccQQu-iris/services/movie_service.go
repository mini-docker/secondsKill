package services

import (
	"fmt"
	"secondsKill/yyccQQu-iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManger(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	fmt.Println("我们获取到的视屏名称为：" + m.repo.GetMovieName())
	return m.repo.GetMovieName()
}
