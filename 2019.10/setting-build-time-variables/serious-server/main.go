package main

import (
	"log"
	"time"
)

// Declare some variables that will be assigned when compiling (go build)
// This ould be anything that would be useful to have in the program
// I use it to set some info that is useful to have in the logs
var (
	version        string
	buildTime      string // Useful when testing new code, to ensure that we're running a fresh build
	buildGitCommit string // More useful for later, when we need to know what code was running at the time
	buildHost      string // Mostly harmless
	buildUser      string // ditto
)

func main() {

	// Print out the info
	log.Printf(`
		Serious Server starting!!!!
		Version: %q
		Build on host %q by user %q from git commit %q on %q`,
		version,
		buildHost, buildUser, buildGitCommit, buildTime,
	)

	log.Println("Work work work")

	err := doSeriousShit()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Done, shutting down!")
}

func doSeriousShit() error {
	time.Sleep(time.Second * 5)
	return nil
}
