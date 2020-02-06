package entities

import "net/http"

type Seeder interface {
	Init(context *SeederContext)
	ReinitData(iteration uint64)
	Loop(iteration uint64)
	Error() error
}

type SeederContext struct {
	Client *http.Client
	Debug  bool
}
