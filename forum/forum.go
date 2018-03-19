package forum

import "time"

// Forum represents a forum
type Forum struct {
	Name     string
	Posts    int
	Threads  int
	URL      string
	LastPost LastPost
}

// LastPost represents a forum last post
type LastPost struct {
	Author      string
	URL         string
	CreatedAt   time.Time
	Achievement PostAchievement
}
