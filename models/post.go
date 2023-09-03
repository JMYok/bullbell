package models

import "time"

// 内存对齐

type Post struct {
	Id          uint64    `json:"id" db:"post_id"`
	AuthorId    uint64    `json:"author_id" db:"author_id"`
	CommunityId uint64    `json:"community_id" db:"community_id"`
	Status      int8      `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}
