package entities

import "github.com/valyala/fasthttp"

type Blaster interface {
	// Name is called exactly once, to get a name for logging
	Name() string
	// Init is called once, while the runner is starting up.
	// It is intended to be used to initialize any blaster context.
	Init(context *BlasterContext)
	// ReinitData is called every 1000 iterations, it allows you to keep some values
	// (like 50 users for example) for a number of requests to make the data seem more organic.
	ReinitData(iteration uint64)
	// Blast is called in every iteration and should usually do one request
	Blast(iteration uint64)
	// Sleep is called after every Blast and allows you to sleep for any time and do some teardown
	// This is especially helpful to combat rate limiting, make requests more organic
	Sleep(iteration uint64)
	// Error is checked after every iteration and should return any error that may have occurred
	// during the iteration which should be logged.
	// You may also handle recovery here for now.
	Error() error
}

type BlasterContext struct {
	// The client you may use for any requests
	Client *fasthttp.Client
	// Do not log *anything* (unless critical) if Debug is set to false
	Debug  bool
}
