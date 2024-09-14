package workerext

import (
	"context"

	"github.com/ThisJohan/delivery-microservice/pkg/di"
)

type worker interface {
	RegisterWorker(context.Context)
}

func StartWorker(w worker, builders ...di.Builder) {
	ctx := context.Background()
	for _, builder := range builders {
		ctx = builder(ctx)
	}
	go func() {
		w.RegisterWorker(ctx)
	}()
}
