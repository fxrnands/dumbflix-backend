package models

import "time"

type Episode struct {
  ID              int                  `json:"id" gorm:"primary_key:auto_increment"`
  Title           string               `json:"title" form:"title" gorm:"type: varchar(255)"`
  ThumbnailEpisode   string               `json:"thumbnailEpisode" form:"thumbnailEpisode" gorm:"type: varchar(255)"`
  LinkFilm        string               `json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)"`
  FilmID          int                  `json:"film_id" form:"id"`
  Film			      FilmResponseEps		   `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  CreatedAt       time.Time            `json:"-"`
  UpdatedAt       time.Time            `json:"-"`
}