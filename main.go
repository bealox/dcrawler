package main

import (
	"log"
	"os"
	"time"

	"github.com/bealox/dcrawler/matcher"
)

func main() {

	/**
	TODO : <LOG>See if I can find elegant way to store log
	**/

	layout := "02-Jan-06"
	f, err := os.OpenFile("/var/log/projects/codedog/crawler"+time.Now().Format(layout)+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	log.SetOutput(f)
	t0 := time.Now()
	matcher.Run()
	t1 := time.Now()
	log.Printf("The call took %v to run.\n", t1.Sub(t0).String())

}
