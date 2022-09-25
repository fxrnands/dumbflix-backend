package handlers

import (
	filmdto "be-dumbflix/dto/film"
	dto "be-dumbflix/dto/result"
	"be-dumbflix/models"
	"be-dumbflix/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

// `path_file` Global variable here ...
var PathFile = os.Getenv("PATH_FILE")

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := h.FilmRepository.FindFilms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image & Video property here ...
	for i, p := range films {
		films[i].ThumbnailFilm = os.Getenv("PATH_FILE") + p.ThumbnailFilm
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var film models.Film
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	film.ThumbnailFilm = os.Getenv("PATH_FILE") + film.ThumbnailFilm

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	year, _ := strconv.Atoi(r.FormValue("year"))
	category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := filmdto.FilmRequest{
		Title:         r.FormValue("title"),
		ThumbnailFilm: r.FormValue("thumbnailFilm"),
		LinkFilm:      r.FormValue("linkFilm"),
		Year:          year,
		CategoryID:    category_id,
		Desc:          r.FormValue("desc"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	film := models.Film{
		Title:         request.Title,
		ThumbnailFilm: filename,
		LinkFilm:      request.LinkFilm,
		Year:          request.Year,
		CategoryID:    category_id,
		Category:      models.CategoryResponse{},
		Desc:          request.Desc,
	}

	// err := mysql.DB.Create(&film).Error
	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	film, _ = h.FilmRepository.GetFilm(film.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(filmdto.UpdateFilmRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film, err := h.FilmRepository.GetFilm(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}

	if request.LinkFilm != "" {
		film.LinkFilm = request.LinkFilm
	}

	if request.Year != 0 {
		film.Year = request.Year
	}

	if request.Desc != "" {
		film.Desc = request.Desc
	}

	data, err := h.FilmRepository.UpdateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FilmRepository.DeleteFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDelFilm(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseFilm(u models.Film) models.FilmResponse {
	return models.FilmResponse{
		ID:            u.ID,
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		LinkFilm:      u.LinkFilm,
		Year:          u.Year,
		Category:      u.Category,
		Desc:          u.Desc,
	}
}

func convertResponseDelFilm(u models.Film) models.FilmResponse {
	return models.FilmResponse{
		ID:    u.ID,
		Title: u.Title,
		Year:  u.Year,
	}
}
