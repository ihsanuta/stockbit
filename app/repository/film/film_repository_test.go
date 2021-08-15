package film

import (
	"stockbit/app/entity"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	repository FilmRepository
)

func prepareTest(t *testing.T) {
	_ = godotenv.Load("../../../.env")
	repository = IntiFilmRepository()
}

func TestGetDataFilm(t *testing.T) {
	prepareTest(t)

	Convey("Test For Get Film Search", t, func() {
		Convey("Success to get film search", func() {
			params := entity.ParamsFilm{
				Title:      "Batman",
				Pagination: 2,
			}
			film, err := repository.GetListsFilm(params)
			if err != nil {
				t.Errorf("ERROR GET DATA FILM : %s", err.Error())
				return
			}

			if strings.ToLower(film.Response) != "true" {
				t.Errorf("ERROR GET DATA FILM Response %s", film.Response)
				return
			}

			t.Log("SUCCES GET DATA FILM")
			return
		})

		Convey("Get film search to many response", func() {
			params := entity.ParamsFilm{
				Title: "B",
			}
			film, err := repository.GetListsFilm(params)
			if err != nil {
				t.Errorf("ERROR GET DATA FILM : %s", err.Error())
				return
			}

			if strings.ToLower(film.Response) == "true" {
				t.Errorf("ERROR Test To Many Response %s", film.Response)
				return
			}

			t.Log("SUCCES Test To Many Response")
			return
		})

		Convey("Get film search without title", func() {
			params := entity.ParamsFilm{
				Pagination: 2,
			}
			film, err := repository.GetListsFilm(params)
			if err != nil {
				t.Errorf("ERROR GET DATA FILM : %s", err.Error())
				return
			}

			if strings.ToLower(film.Response) == "true" {
				t.Errorf("ERROR Test To Many Response %s", film.Response)
				return
			}

			t.Log("SUCCES Get film search without title")
			return
		})
	})
}

func TestGetDataFilmID(t *testing.T) {
	prepareTest(t)

	Convey("Test For Get Film Search By ID", t, func() {
		Convey("Success to get film By ID", func() {
			idFilm := "tt4853102"
			film, err := repository.GetFilmByID(idFilm)
			if err != nil {
				t.Errorf("ERROR GET DATA FILM : %s", err.Error())
				return
			}

			if strings.ToLower(film.Response) != "true" {
				t.Errorf("ERROR GET DATA FILM Response %s", film.Response)
				return
			}

			t.Log("SUCCES GET DATA FILM BY ID")
			return
		})

		Convey("Get film incorrect id", func() {
			idFilm := ""
			film, err := repository.GetFilmByID(idFilm)
			if err != nil {
				t.Errorf("ERROR GET DATA FILM : %s", err.Error())
				return
			}

			if strings.ToLower(film.Response) == "true" {
				t.Errorf("ERROR Test incorrect id %s", film.Response)
				return
			}

			t.Log("SUCCES Test incorrect id")
			return
		})
	})
}
