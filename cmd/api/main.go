package main

import "log"

func main() {

	cnfg := config{
		addr: ":3000",
	}

	app := &application{
		config: cnfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
