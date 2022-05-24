package entity

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// TFertilizerHistory represents a row from 'public.t_fertilizer_history'.
type TFertilizerHistory struct {
	FertilizerHistoryID int       `json:"fertilizer_history_id"` // fertilizer_history_id
	PlantID             int       `json:"plant_id"`              // plant_id
	FertilizerID        int       `json:"fertilizer_id"`         // fertilizer_id
	UserID              int       `json:"user_id"`               // user_id
	DeleteFlg           bool      `json:"delete_flg"`            // delete_flg
	CreateUser          string    `json:"create_user"`           // create_user
	CreateDatetime      time.Time `json:"create_datetime"`       // create_datetime
	UpdateUser          string    `json:"update_user"`           // update_user
	UpdateDatetime      time.Time `json:"update_datetime"`       // update_datetime
	Version             int       `json:"version"`               // version
}
