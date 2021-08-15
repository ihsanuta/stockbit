package film

import (
	"stockbit/app/entity"
	"stockbit/app/repository/film"
)

type FilmService interface {
	GetFilmByParams(params entity.ParamsFilm) (entity.FilmLists, error)
	GetByFilmByID(filmID string) (entity.DataFilmDetail, error)
}

type filmService struct {
	filmRepository film.FilmRepository
}

func InitFilmService(film film.FilmRepository) FilmService {
	return &filmService{
		filmRepository: film,
	}
}

func (f *filmService) GetFilmByParams(params entity.ParamsFilm) (entity.FilmLists, error) {
	return f.filmRepository.GetListsFilm(params)
}

func (f *filmService) GetByFilmByID(filmID string) (entity.DataFilmDetail, error) {
	return f.filmRepository.GetFilmByID(filmID)
}
