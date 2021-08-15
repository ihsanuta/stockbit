package service

import (
	"stockbit/app/repository"
	"stockbit/app/service/film"
)

type Service struct {
	Film film.FilmService
}

func Init(repository *repository.Repository) *Service {
	return &Service{
		Film: film.InitFilmService(
			repository.Film,
		),
	}
}
