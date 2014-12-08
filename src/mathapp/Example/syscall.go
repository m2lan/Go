package main

import (
	"fmt"
	"syscall"
)

type MemStatus struct {
	All  uint32 `json:"all"`
	Used uint32 `json:"used"`
	Free uint32 `json:"free"`
	Self uint64 `json:"self"`
}

// 获取磁盘占用情况
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return disk
}

func main() {
	fmt.Println(DiskUsage("/Users/m2lan"))

	disk := DiskUsage("/Users/m2lan")
	value := ((disk.All / 1000) / 1000) / 1000
	fmt.Println(value)
}

// B kb M GB TB
//998973898752 975560448 952695.75 930.366
