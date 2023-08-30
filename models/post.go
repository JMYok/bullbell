package models

type Post struct {
	PostId      int64  `db:"post_id"`
	Title       string `db:"title"`
	Content     string `db:"content"`
	AuthorId    int64  `db:"author_id"`
	CommunityId int64  `db:"community_id"`
	Status      int8   `db:"status"`
	CreateTime  string `db:"create_time"`
	UpdateTime  string `db:"update_time"`
}
