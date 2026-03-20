package interruptor

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type Interruptor struct {
	gRPCInterruptor *grpc.Server
	signal          chan os.Signal
	logger          *slog.Logger
}

func NewInterruptor(srv *grpc.Server, logger *slog.Logger) *Interruptor {
	return &Interruptor{
		gRPCInterruptor: srv,
		signal:          make(chan os.Signal, 1),
		logger:          logger,
	}
}

func (i *Interruptor) Run() {
	i.startCatchingSignal()
	go func() {
		i.shutdown()
	}()
}

func (i *Interruptor) startCatchingSignal() {
	i.logger.Info("Starting signal catching")
	signal.Notify(i.signal, syscall.SIGTERM, syscall.SIGINT)
}

func (i *Interruptor) shutdown() {
	<-i.signal
	i.logger.Info("Server is shutting down gracefully...")
	i.gRPCInterruptor.GracefulStop()
}
