package main

import (
	. "tzgolkin/controller"
	. "tzgolkin/disp"
	"math/rand"
	"flag"
	"os"
	"runtime/pprof"
	"log"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

	r := rand.New(rand.NewSource(78))

	ctrl := MakeController(r)
	disp := MakeDisplay(ctrl)

	disp.Run()
}