package filmdto

type FilmRequest struct {
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type: varchar(255)" `
	LinkFilm      string `json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)" `
	Year          int    `json:"year" form:"year" gorm:"type: int" validate:"required"`
	CategoryID    int    `json:"category_id" form:"category_id" gorm:"type: int"`
	Desc          string `json:"desc" gorm:"type:text" form:"desc" validate:"required"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type: varchar(255)" `
	LinkFilm      string `json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)" `
	Year          int    `json:"year" form:"year" gorm:"type: int"`
	CategoryID    int    `json:"category_id" form:"category_id" gorm:"type: int"`
	Desc          string `json:"desc" gorm:"type:text" form:"desc"`
}
