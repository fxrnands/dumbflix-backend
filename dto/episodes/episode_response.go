package episodesdto

type EpisodeResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type: varchar(255)"`
	LinkFilm      string `json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)"`
	FilmID        int    `json:"film_id" form:"film_id"`
}