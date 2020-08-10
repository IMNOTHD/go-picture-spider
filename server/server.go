package server

import (
	"context"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	pb "go-picture-spider/api"
	. "go-picture-spider/model"
	"go-picture-spider/service"
)

type PictureSpiderService struct{}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func Start() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPictureSpiderServer(grpcServer, &PictureSpiderService{})
	log.Println(Address + " net.Listening...")

	httpServer := ProvideHttp(Address, grpcServer)

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	if err = httpServer.Serve(listener); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 实现PictureUrl方法
func (p *PictureSpiderService) PictureUrl(ctx context.Context, req *pb.PictureUrlRequest) (*pb.PictureUrlResponse, error) {
	var r *UrlAnalyzerResult
	switch strings.ToLower(req.Type) {
	case "pixiv":
		r = service.Pixiv(req.Id, req.Manga)
		break
	case "donmai":
		r = service.Donmai(req.Id)
		break
	case "medibang":
		r = service.Medibang(req.Id)
		break
	case "yande":
		r = service.Yande(req.Id)
		break
	default:
		return &pb.PictureUrlResponse{
			Code: 400,
			Msg:  "no such type",
			Data: nil,
		}, nil
	}

	if r.Code != Success {
		return &pb.PictureUrlResponse{
			Code: r.Code,
			Msg:  r.Err.Error(),
		}, nil
	}

	res := pb.PictureUrlResponse{
		Code: r.Code,
		Msg:  "",
		Data: &pb.PictureUrlData{
			OriginalUrl: r.OriginalUrl,
			IsR18:       r.IsR18,
			Referer:     r.Referer,
		},
	}

	return &res, nil
}
