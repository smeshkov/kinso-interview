package ctx

import "github.com/smeshkov/kinso-interview/app/storage"

var (
	DB *storage.Storage
)

// creates shared application context
func Setup() {
	DB = storage.New()
}
