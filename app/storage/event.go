package storage

type Event struct {
	ID        string
	UserID    string
	CreatedAt string
	Source    string
	Weight    float64
	RawData   string
}
