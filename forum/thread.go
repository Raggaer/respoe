package forum

import (
	"time"

	"github.com/raggaer/respoe/util"
)

// Thread forum thread
type Thread struct {
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
