package entity

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// MFertilizer represents a row from 'public.m_fertilizer'.
type MFertilizer struct {
	FertilizerID   int       `json:"fertilizer_id"`   // fertilizer_id
	FertilizerName string    `json:"fertilizer_name"` // fertilizer_name
	DeleteFlg      bool      `json:"delete_flg"`      // delete_flg
	CreateUser     string    `json:"create_user"`     // create_user
	CreateDatetime time.Time `json:"create_datetime"` // create_datetime
	UpdateUser     string    `json:"update_user"`     // update_user
	UpdateDatetime time.Time `json:"update_datetime"` // update_datetime
	Version        int       `json:"version"`         // version
}
