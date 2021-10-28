package usr

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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

func (w *Worker) Usrinfo(id string) (user *User) {
	go func() {
		if id == "16701" {
			user = &User{ID: id, Name: "Mir"}
			defer w.cancel()
			return
		}
	}()
	for {
		select {
		case <-w.ctx.Done():
			return
		case <-time.After(10 * time.Millisecond):
		}
	}
}
