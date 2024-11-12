package main

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/router"
)

func main() {

	// init template
	pagehandler.InitTemplates()

	// init wire
	controllerHandler, db, mid, logger := InitializeService()
	if db == nil {
		return
	}

	// db close
	defer db.Close()

	// zap logger close
	defer logger.Sync()

	// init routes
	router.InitRoute(controllerHandler, mid)
}
