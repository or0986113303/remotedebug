package device

import (
	"context"
	"net"

	"github.com/shirou/gopsutil/host"
	"github.com/sirupsen/logrus"
)

type HostInfo interface {
	GetHostname() (string, error)
	GetDiskStatus(path string) (result DiskStatus, err error)
	GetMacAddress() (macAddrs map[string]string, err error)
}

type Worker struct {
	ctx    context.Context
	cancel context.CancelFunc
	log    *logrus.Entry
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func New() (worker *Worker) {
	ctx, cancel := context.WithCancel(context.TODO())
	worker = &Worker{
		ctx:    ctx,
		cancel: cancel,
		log:    logrus.WithFields(logrus.Fields{"origin": "device"}),
	}
	return
}

// Getdiskstatus ...
func (w *Worker) GetDiskStatus(path string) (result DiskStatus, err error) {
	/*
		if runtime.GOOS == "linux" {
			fs := syscall.Statfs_t{}
			err = syscall.Statfs(path, &fs)
			if err != nil {
				w.log.Panic(err.Error())
				return
			}
			result.All = fs.Blocks * uint64(fs.Bsize)
			result.Used = fs.Bfree * uint64(fs.Bsize)
			result.Free = result.All - result.Used
			return
		}
	*/
	return
}

// GetMacAddress ...
func (w *Worker) GetMacAddress() (macAddrs map[string]string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		w.log.Panic(err.Error())
	}
	tmpmap := make(map[string]string)
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		tmpmap[netInterface.Name] = macAddr
		macAddrs = tmpmap
	}
	return
}

func (w *Worker) GetHostname() (string, error) {
	hInfo, err := host.Info()
	return hInfo.Hostname, err
}
