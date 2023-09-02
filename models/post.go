package models

type Post struct {
	Id          uint64 `json:"id" db:"post_id"`
	Title       string `json:"title" db:"title"`
	Content     string `json:"content" db:"content"`
	AuthorId    int64  `json:"author_id" db:"author_id"`
	CommunityId int64  `json:"community_id" db:"community_id"`
	Status      int8   `json:"status" db:"status"`
	CreateTime  string `json:"create_time" db:"create_time"`
	UpdateTime  string `json:"update_time" db:"update_time"`
}
