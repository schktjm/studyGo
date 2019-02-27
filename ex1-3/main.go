package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    unefficient()
    useStringJoin()
}

func unefficient() {
    start := time.Now()
    var s,sep string
    for i:=1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    secs := time.Since(start).Seconds()
    fmt.Printf("for reading %fs: %s\n",secs,s)
}

func useStringJoin() {
    start := time.Now()
    s := strings.Join(os.Args[1:]," ")
    secs := time.Since(start).Seconds()
    fmt.Printf("string join reading %fs: %s\n",secs,s)
}
