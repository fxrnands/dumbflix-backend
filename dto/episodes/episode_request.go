package episodesdto

type CreateEpisodeRequest struct {
	Title            string `json:"title" validate:"required"`
	ThumbnailEpisode string `json:"thumbnailEpisode" validate:"required"`
	LinkFilm         string `json:"linkFilm" validate:"required"`
	FilmID           int    `json:"film_id" form:"film_id" gorm:"type: int" validate:"required"`
}

type UpdateEpisodeRequest struct {
	Title            string `json:"title"`
	ThumbnailEpisode string `json:"thumbnailFilm"`
	LinkFilm         string `json:"linkFilm"`
	FilmID           int    `json:"film_id" form:"film_id" gorm:"type: int"`
}