package grpc_server

import "net"

func (server *Server) Start(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	server.listener = listener

	return server.grpcServer.Serve(listener)
}

func (server *Server) Close() error {
	server.grpcServer.GracefulStop()
	if server.listener == nil {
		return nil
	}

	return server.listener.Close()
}
