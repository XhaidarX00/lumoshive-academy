package main

import (
	"latihan/router"
	"log"
)

func main() {
	ch, err := InitializeService()
	if err != nil {
		panic(err)
	}

	// db close
	defer ch.DB.Close()

	// zap logger close
	defer ch.Logger.Sync()

	// init routes
	router.InitRoute(ch.Handler, ch.Logger)

	log.Println("Masuk Main")
}
