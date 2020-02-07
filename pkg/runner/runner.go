package runner

import (
	"github.com/deletescape/noise/pkg/entities"
	"log"
	"net/http"
	"sync"
)

var seeders []entities.Seeder

func RegisterSeeder(seeder entities.Seeder) {
	seeders = append(seeders, seeder)
}

func Run(client *http.Client, debug bool) {
	log.Println("initializing", len(seeders), "seeders")
	eachSeeder(func(seeder entities.Seeder) {
		seeder.Init(&entities.SeederContext{
			Client: client,
			Debug:  debug,
		})
	})

	log.Println("starting up")

	var loopGroup sync.WaitGroup
	// every seeder loop runs in its own goroutine
	eachSeeder(func(seeder entities.Seeder) {
		loopGroup.Add(1)
		go func() {
			defer loopGroup.Done()
			for i := uint64(0); ; i++ {
				if i%1000 == 0 {
					seeder.ReinitData(i)
				}

				seeder.Loop(i)

				seeder.Sleep(i)

				if err := seeder.Error(); err != nil {
					log.Println("error in seeder:", err)
				}
			}
		}()
	})
	loopGroup.Wait()
}

func eachSeeder(f func(entities.Seeder)) {
	for _, seeder := range seeders {
		f(seeder)
	}
}
