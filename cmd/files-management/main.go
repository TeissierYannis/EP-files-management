// Package main provides the entry point to the files-management application.
// It sets up the gRPC server and registers the FileService.
package main

import (
	"context"
	"files-management/pkg/config"
	"files-management/pkg/filehandler"
	"files-management/pkg/logger"
	"files-management/proto/gen"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// fileServiceServer implements the FileServiceServer interface.
// It uses a FileHandler to handle file operations.
type fileServiceServer struct {
	gen.UnimplementedFileServiceServer
	fileHandler filehandler.FileHandler
}

// NewFileServiceServer creates a new fileServiceServer with the provided FileHandler.
// It returns a pointer to the created fileServiceServer.
func NewFileServiceServer(fh filehandler.FileHandler) *fileServiceServer {
	return &fileServiceServer{fileHandler: fh}
}

// OpenFile is a method of fileServiceServer that opens a file at the provided path.
// It returns a FileResponse with the file content or an error if the file could not be opened.
func (s *fileServiceServer) OpenFile(ctx context.Context, in *gen.FileRequest) (*gen.FileResponse, error) {
	img, err := s.fileHandler.LoadImage(in.GetPath())
	if err != nil {
		return nil, err
	}
	return &gen.FileResponse{Content: img}, nil
}

// main is the entry point of the application.
// It sets up the gRPC server, registers the FileService and starts listening for incoming connections.
func main() {
	// Setup and parse flags
	port := flag.Int("port", 50051, "The server port")
	flag.Parse()

	// Initialize configuration and logger
	cfg, err := config.LoadConfig("")
	if err != nil {
		logger.Fatal("Failed to load configuration: ", err)
	}
	logger.Init(cfg.LogLevel)

	// Log configuration details
	logger.Info("============== Configuration ==============")
	logger.Info("Log level: ", cfg.LogLevel)
	logger.Info("Server starting on port: ", *port)
	logger.Info("===========================================\n")

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal("Failed to listen: ", err)
	}
	var opts []grpc.ServerOption
	fh := filehandler.DefaultFileHandler{}
	grpcServer := grpc.NewServer(opts...)
	gen.RegisterFileServiceServer(grpcServer, NewFileServiceServer(fh))

	logger.Info("gRPC server listening at ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to serve: ", err)
	}
}
