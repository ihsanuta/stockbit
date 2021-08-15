package repository

import "stockbit/app/repository/film"

type Repository struct {
	Film film.FilmRepository
}

func Init() *Repository {
	return &Repository{
		Film: film.IntiFilmRepository(),
	}
}
