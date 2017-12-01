package forum

import "time"

// Post thread post content
type Post struct {
	Content   string
	Author    string
	CreatedAt time.Time
}
