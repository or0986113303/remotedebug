package system

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Parameters struct {
	Execute string   `json:"execute"`
	Args    []string `json:"args"`
}

type System interface {
	Sleep()
}

type Worker struct {
	ctx    context.Context
	cancel context.CancelFunc
	log    *logrus.Entry
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

func (w *Worker) Sysinfo() (Uptime int64, err error) {
	go func() {
		/*
			var sysinfo syscall.Sysinfo_t
			syscall.Sysinfo(&sysinfo)
			Uptime = int64(sysinfo.Uptime)
		*/
		time.Sleep(3 * time.Second)
		defer w.cancel()
	}()
	for {
		select {
		case <-w.ctx.Done():
			return
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func (w *Worker) getCPUClockbyTimePoint() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
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

func (w *Worker) CPUUseage() (result float64) {
	go func() {
		idle0, total0 := w.getCPUClockbyTimePoint()
		time.Sleep(1 * time.Second)
		idle1, total1 := w.getCPUClockbyTimePoint()
		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)
		result = 100 * (totalTicks - idleTicks) / totalTicks
		w.log.Infof("CPU usage is %f%% [busy: %f, total: %f]\n", result, totalTicks-idleTicks, totalTicks)
		defer w.cancel()
	}()
	for {
		select {
		case <-w.ctx.Done():
			return
		case <-time.After(10 * time.Millisecond):
		}
	}
}
