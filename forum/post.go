package forum

import (
	"time"

	"github.com/raggaer/respoe/client"
	"github.com/raggaer/respoe/util"
)

// PostList thread post list
type PostList struct {
	List       []*Post
	Title      string
	ForumName  string
	ForumURL   string
	Items      []*client.Item
	Pagination *util.Pagination
}

// Post thread post content
type Post struct {
	Staff        bool
	ValuedPoster bool
	Content      string
	ContentText  string
	Author       string
	Badges       []*client.Badge
	Avatar       string
	Achievement  PostAchievement
	CreatedAt    time.Time
}

// PostAchievement post league achievements
type PostAchievement struct {
	Alt string
	URL string
}
