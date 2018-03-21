package forum

import (
	"time"

	"github.com/raggaer/respoe/util"
)

// PostList thread post list
type PostList struct {
	List       []*Post
	Title      string
	ForumName  string
	ForumURL   string
	Pagination *util.Pagination
}

// Post thread post content
type Post struct {
	Staff        bool
	ValuedPoster bool
	Content      string
	Author       string
	Badges       []PostBadge
	Avatar       string
	Achievement  PostAchievement
	CreatedAt    time.Time
}

// PostBadge post supporter pack badge
type PostBadge struct {
	Name string
	URL  string
}

// PostAchievement post league achievements
type PostAchievement struct {
	Alt string
	URL string
}
