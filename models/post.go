package models

import "time"

// 内存对齐

type Post struct {
	Id          uint64    `json:"id,string" db:"post_id"`
	AuthorId    uint64    `json:"author_id,string" db:"author_id"`
	CommunityId uint64    `json:"community_id,string" db:"community_id"`
	Status      int8      `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*CommunityDetail `json:"community"`
	*Post            `json:"post"`
}
