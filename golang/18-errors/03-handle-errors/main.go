package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("defered call in main")
	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("error happened:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	f, err := os.Open("not-existing.txt")
	if err != nil {
		//log.Println("Error happened:", err) // just prints...
		//log.Fatalln("Fatal error happened (exiting):", err) //immedita exit, no defffered
		log.Panicln("Error happened, PANIC! :", err) //exits (if not recivered) but still calls deferred
	}
	defer f.Close()
	fmt.Println("check the logfile")
}
