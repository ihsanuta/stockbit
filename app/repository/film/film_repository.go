package film

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"stockbit/app/entity"
	"stockbit/app/global"

	"github.com/sirupsen/logrus"
)

type FilmRepository interface {
	GetListsFilm(params entity.ParamsFilm) (entity.FilmLists, error)
	GetFilmByID(filmID string) (entity.DataFilmDetail, error)
}

type filmRepository struct {
}

func IntiFilmRepository() FilmRepository {
	return &filmRepository{}
}

func (f *filmRepository) GetListsFilm(params entity.ParamsFilm) (entity.FilmLists, error) {
	var result entity.FilmLists

	query := params.ToParams().Encode()
	logrus.Print(query)

	filmClient := &http.Client{}
	logrus.Print("URL GET PARAMS : ", fmt.Sprintf("%s?apikey=%s%s", os.Getenv(global.EnvSourceURL), os.Getenv(global.EnvSourceKey), query))
	filmReq, err := http.NewRequest("GET", fmt.Sprintf("%s?apikey=%s%s", os.Getenv(global.EnvSourceURL), os.Getenv(global.EnvSourceKey), query), nil)
	if err != nil {
		return result, err
	}

	filmResp, err := filmClient.Do(filmReq)
	if err != nil {
		return result, err
	}
	defer filmResp.Body.Close()

	body, err := ioutil.ReadAll(filmResp.Body)
	if err != nil {
		return result, err
	}

	if filmResp.StatusCode != http.StatusOK {
		return result, errors.New("Get Data Film Failed")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (f *filmRepository) GetFilmByID(filmID string) (entity.DataFilmDetail, error) {
	var result entity.DataFilmDetail

	filmClient := &http.Client{}
	filmReq, err := http.NewRequest("GET", fmt.Sprintf("%s?apikey=%s&i=%s", os.Getenv(global.EnvSourceURL), os.Getenv(global.EnvSourceKey), filmID), nil)
	if err != nil {
		return result, err
	}

	filmResp, err := filmClient.Do(filmReq)
	if err != nil {
		return result, err
	}
	defer filmResp.Body.Close()

	body, err := ioutil.ReadAll(filmResp.Body)
	if err != nil {
		return result, err
	}

	if filmResp.StatusCode != http.StatusOK {
		return result, errors.New("Get Data Film Failed")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
