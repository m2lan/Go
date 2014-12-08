package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"syscall"
)

// 获取磁盘占用情况
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

// 获取磁盘占用情况
func diskUsage(path string) (disk DiskStatus) {
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

func DiskUsages(path string) (all, used uint64) {
	disk := diskUsage(path)
	all = ((disk.All / 1000) / 1000) / 1000
	used = ((disk.Used / 1000) / 1000) / 1000

	return all, used
}
