package client

import "time"

// PrivateMessage represents an account private message
type PrivateMessage struct {
	Unread     bool
	Subject    string
	Sender     string
	ReceivedAt time.Time
	URL        string
}
