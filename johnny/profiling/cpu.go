package profiling

import (
	"log"
	"os"
	"runtime/pprof"
)

func ProfileCpuUsage(cpuprofile string) func() {

	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	return func() {
		defer f.Close() // error handling omitted for example
		pprof.StopCPUProfile()
	}

}

