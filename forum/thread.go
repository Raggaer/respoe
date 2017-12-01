package forum

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
	Pagination *Pagination
}
