package forum

import (
	"time"

	"github.com/raggaer/respoe/util"
)

// ThreadList forum thread list
type ThreadList struct {
	ForumName  string
	Pagination *util.Pagination
	List       []*Thread
}

// Thread forum thread
type Thread struct {
	ID         int64
	Staff      bool
	Support    bool
	Sticky     bool
	Locked     bool
	Title      string
	Replies    int
	Views      int
	URL        string
	Author     string
	CreatedAt  time.Time
	Pagination *util.Pagination
}
