package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "math"
    "os"
    "runtime"
    "strconv"
    "strings"
    "time"
)


var (
    numBurn        int
    updateInterval int
    //targetLoad     float64
)

func cpuBurn(targetLoad float64, loops int, cpuUsage_channel chan float64) {

    var cpuUsage float64 = 0

    for {
        select {
        case cpuUsage = <- cpuUsage_channel: {
            }
        default:

        }
        if cpuUsage < targetLoad {
             for i := 0; i < loops; i++ {
                }
        } else {
            time.Sleep(time.Millisecond * 100)
        }
        runtime.Gosched()
    }
    _ = cpuUsage

}

func init() {
    flag.IntVar(&numBurn, "n", 0, "number of cores to burn (0 = all)")
    flag.IntVar(&updateInterval, "u", 100, "milliseconds between updates (0 = don't update)")
    //flag.Float64Var(&targetLoad, "c", 10, "Target CPU Load in %")
    flag.Parse()
    if numBurn <= 0 {
        numBurn = runtime.NumCPU()
    }
}


func getCPUSample() (idle, total uint64) {
    contents, err := ioutil.ReadFile("/proc/stat")
    if err != nil {
        return
    }
    lines := strings.Split(string(contents), "\n")
    for _, line := range(lines) {
        fields := strings.Fields(line)
        if fields[0] == "cpu" {
            numFields := len(fields)
            for i := 1; i < numFields; i++ {
                val, err := strconv.ParseUint(fields[i], 10, 64)
                if err != nil {
                    fmt.Println("Error: ", i, fields[i], err)
                }
                total += val // tally up all the numbers to get total ticks
                if i == 4 {  // idle is the 5th field in the cpu line
                    idle = val
                }
            }
            return
        }
    }
    return
}


func main() {

     if targetLoad_str, err := os.LookupEnv("CPU_LOAD_TO"); err == true {

        if targetLoad, err := strconv.ParseFloat(targetLoad_str, 64); err == nil {
            runtime.GOMAXPROCS(numBurn)

            var loops_count int = int(math.Pow(2, float64( 16 ) ) ) // loops count

            var cpuUsage_channels = make([]chan float64, numBurn)
            for i := range cpuUsage_channels {
                cpuUsage_channels[i] = make(chan float64, 3)
            }

            fmt.Printf("Burning %d CPUs/cores\n", numBurn)
            for i := 0; i < numBurn; i++ {
                go cpuBurn(targetLoad, loops_count, cpuUsage_channels[i])
            }
            if updateInterval > 0 {
                t := time.Tick(time.Duration(updateInterval) * time.Millisecond)
                idle0, total0 := getCPUSample()
                for secs := updateInterval; ; secs += updateInterval {
                    <-t
                    idle1, total1 := getCPUSample()

                    idleTicks := float64(idle1 - idle0)
                    totalTicks := float64(total1 - total0)
                    cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks


                    for i := 0; i < numBurn; i++ {
                        for len(cpuUsage_channels[i]) > 0 {
                             <-cpuUsage_channels[i]
                        }
                        cpuUsage_channels[i] <- cpuUsage
                    }

                    //fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
                }
            } else {
                select {} // wait forever
            }

	    } else {
            fmt.Printf("Variable CPU_LOAD_TO is not integer\n")
	    }


     } else {
        fmt.Printf("Variable CPU_LOAD_TO is not set in Environment\n")
     }

}
