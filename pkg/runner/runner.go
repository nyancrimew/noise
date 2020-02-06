package runner

import (
	"github.com/deletescape/noise/pkg/entities"
	"log"
	"net/http"
	"sync"
)

var seeders []entities.Seeder

func RegisterSeeder(seeder entities.Seeder)  {
	seeders = append(seeders, seeder)
}

func Run(client *http.Client, debug bool) error  {
	log.Println("initializing", len(seeders), "seeders")
	eachSeeder(func(seeder entities.Seeder) {
		seeder.Init(&entities.SeederContext{
			Client: client,
			Debug:  debug,
		})
	})

	log.Println("starting up")
	for i := uint64(0); ;i++ {
		log.Println("iteration", i)

		if i % 1000 == 0 {
			eachSeeder(func(seeder entities.Seeder) {
				seeder.ReinitData(i)
			})
		}

		var wg sync.WaitGroup
		eachSeeder(func(seeder entities.Seeder) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				seeder.Loop(i)
			}()
		})
		wg.Wait()

		eachSeeder(func(seeder entities.Seeder) {
			err := seeder.Error()
			if err != nil {
				// TODO: have some sort of recovery logic
				log.Println("error in seeder:", err)
			}
		})
	}
}

func eachSeeder(f func(entities.Seeder)) {
	for _, seeder := range seeders {
		f(seeder)
	}
}
