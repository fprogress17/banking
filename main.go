package main

//import "github.com/ashishjuyal/banking/app"
import (
	"github.com/fprogress17/banking/app"
	"github.com/fprogress17/banking/logger"
)

func main() {
	// log.Println("starting our application")
	logger.Info("starting the application")
	app.Start()
	//Start()
}
