package entity

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// TPlant represents a row from 'public.t_plant'.
type TPlant struct {
	PlantID         int       `json:"plant_id"`          // plant_id
	PlantName       string    `json:"plant_name"`        // plant_name
	PlantDetail     string    `json:"plant_detail"`      // plant_detail
	PlantCotegoryID int       `json:"plant_cotegory_id"` // plant_cotegory_id
	SeasonID        int       `json:"season_id"`         // season_id
	DeleteFlg       bool      `json:"delete_flg"`        // delete_flg
	CreateUser      string    `json:"create_user"`       // create_user
	CreateDatetime  time.Time `json:"create_datetime"`   // create_datetime
	UpdateUser      string    `json:"update_user"`       // update_user
	UpdateDatetime  time.Time `json:"update_datetime"`   // update_datetime
	Version         int       `json:"version"`           // version
}