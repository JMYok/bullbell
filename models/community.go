package models

import "time"

type Community struct {
	Id           uint64 `json:"id" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction,omitempty" db:"introduction"`
}

type CommunityDetail struct {
	Id           uint64    `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}
