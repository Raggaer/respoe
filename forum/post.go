package forum

import "time"

// Post thread post content
type Post struct {
	Staff        bool
	ValuedPoster bool
	Content      string
	Author       string
	Badges       []PostBadge
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
