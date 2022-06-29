package server

import (
	"fmt"
	"github.com/bhatti/GSSI/internal/config"
	"google.golang.org/grpc"
	"net"
)

type Handle struct {
	grpcServer *grpc.Server
	listner    net.Listener
}

func (h *Handle) Addr() net.Addr {
	return h.listner.Addr()
}

func (h *Handle) Close() (err error) {
	if h.grpcServer != nil {
		h.grpcServer.Stop()
	}
	if h.listner != nil {
		err = h.listner.Close()
	}
	h.grpcServer = nil
	h.listner = nil
	return
}

func (h *Handle) Serve() error {
	return h.grpcServer.Serve(h.listner)
}

type NullAuthorizer struct {
}

func (n NullAuthorizer) Authorize(string, string, string) error {
	return nil
}

type NoAuthorizer struct {
}

func (n NoAuthorizer) Authorize(string, string, string) error {
	return fmt.Errorf("auth error")
}

func StartServers(
	config *config.Config,
	grpcOpts ...grpc.ServerOption) (handle *Handle, err error) {
	return nil, fmt.Errorf("not implemented")
}
