package main

import "log"

func main() {

	cnfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cnfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
