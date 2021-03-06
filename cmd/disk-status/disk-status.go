package main

import (
	"fmt"
	"os"

	"github.com/bborbe/disk_utils/status"
	flag "github.com/bborbe/flagenv"
)

var (
	pathPtr = flag.String("path", "/", "Path")
)

func main() {
	flag.Parse()
	disk, err := status.DiskUsage(*pathPtr)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	fmt.Print(disk.String())
}
