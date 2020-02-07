package runner

import (
	"github.com/deletescape/noise/pkg/entities"
	"github.com/valyala/fasthttp"
	"log"
	"sync"
)

var blasters []*entities.Blaster

func RegisterBlaster(blaster entities.Blaster) {
	blasters = append(blasters, &blaster)
}

func Run(debug bool) {
	log.Println("initializing", len(blasters), "blasters")
	eachBlaster(func(blaster *entities.Blaster) {
		(*blaster).Init(&entities.BlasterContext{
			Client: &fasthttp.Client{},
			Debug:  debug,
		})
	})

	log.Println("starting up")

	var loopGroup sync.WaitGroup
	// every blaster loop runs in its own goroutine
	eachBlaster(func(blaster *entities.Blaster) {
		b := *blaster
		name := b.Name()
		loopGroup.Add(1)
		go func() {
			defer loopGroup.Done()
			for i := uint64(0); ; i++ {
				if i % 100 == 0 {
					log.Println(name, "iter", i)
				}

				if i%1000 == 0 {
					b.ReinitData(i)
				}

				b.Blast(i)

				b.Sleep(i)

				if err := b.Error(); err != nil {
					log.Println("error in blaster:", err)
				}
			}
		}()
	})
	loopGroup.Wait()
}

func eachBlaster(f func(*entities.Blaster)) {
	for _, blaster := range blasters {
		f(blaster)
	}
}
