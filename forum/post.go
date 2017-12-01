package forum

import "time"

// Post thread post content
type Post struct {
	Staff     bool
	Content   string
	Author    string
	CreatedAt time.Time
}
