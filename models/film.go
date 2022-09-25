package models

import "time"

type Film struct {
	ID            int              `json:"id" gorm:"primary_key:auto_increment"`
	Title         string           `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string           `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type: varchar(255)"`
	LinkFilm      string           `json:"linkFilm" gorm:"type:text" form:"linkFilm"`
	Year          int              `json:"year" form:"year" gorm:"type: int"`
	CategoryID    int              `json:"category_id" form:"category_id"`
	Category      CategoryResponse `json:"category"`
	Desc          string           `json:"desc" gorm:"type:text" form:"desc"`
	CreatedAt     time.Time        `json:"-"`
	UpdatedAt     time.Time        `json:"-"`
}

type FilmResponse struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	ThumbnailFilm string           `json:"thumbnailFilm"`
	LinkFilm      string           `json:"linkFilm"`
	Year          int              `json:"year"`
	CategoryID    int              `json:"-"`
	Category      CategoryResponse `json:"category"`
	Desc          string           `json:"desc"`
}

type FilmResponseEps struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	ThumbnailFilm string           `json:"thumbnailFilm"`
	LinkFilm      string           `json:"linkFilm"`
	Year          int              `json:"year"`
	CategoryID    int              `json:"-"`
	Category      CategoryResponse `json:"category"`
	Desc          string           `json:"desc"`
}

func (FilmResponse) TableName() string {
	return "films"
}
func (FilmResponseEps) TableName() string {
	return "films"
}
