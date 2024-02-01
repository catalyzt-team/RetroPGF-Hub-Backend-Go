package server

import (
	datacenterhttphandler "RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/datacenter/datacenterHttpHandler"
	datacenterPb "RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/datacenter/datacenterPb"
	datacenterrepository "RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/datacenter/datacenterRepository"
	datacenterusecase "RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/datacenter/datacenterUsecase"

	grpcconn "RetroPGF-Hub/RetroPGF-Hub-Backend-Go/pkg/grpcConn"
	"log"
)

func (s *server) datacenterService() {
	datacenterRepo := datacenterrepository.NewDatacenterRepository(s.db)
	datacenterUsecase := datacenterusecase.NewDatacenterUsecase(datacenterRepo, &s.cfg.Grpc)
	datacenterHttpHandler := datacenterhttphandler.NewDatacenterHttpHandler(datacenterUsecase)

	datacenterGrpc := datacenterhttphandler.NewdatacenterGrpcHandler(datacenterUsecase)
	// Grpc client
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.DatacenterUrl)
		datacenterPb.RegisterDataCenterGrpcServiceServer(grpcServer, datacenterGrpc)

		log.Printf("datacenter grpc listening on %s", s.cfg.Grpc.DatacenterUrl)
		grpcServer.Serve(lis)
	}()
	_ = datacenterHttpHandler
}