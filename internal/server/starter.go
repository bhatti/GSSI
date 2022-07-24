package server

import (
	"fmt"
	"github.com/bhatti/GSSI/internal/config"
	"google.golang.org/grpc"
	"net"
)

type Handle struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func (h *Handle) Addr() net.Addr {
	return h.listener.Addr()
}

func (h *Handle) Close() (err error) {
	if h.grpcServer != nil {
		h.grpcServer.Stop()
	}
	if h.listener != nil {
		err = h.listener.Close()
	}
	h.grpcServer = nil
	h.listener = nil
	return
}

func (h *Handle) Serve() error {
	return h.grpcServer.Serve(h.listener)
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
