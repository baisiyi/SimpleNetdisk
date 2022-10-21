package gateway

import (
	"context"
	pb "simpleNetdisk/protoc"
)

type fileMgtServiceImpl struct {
}

func NewFileMgtServiceImpl() *fileMgtServiceImpl {
	return &fileMgtServiceImpl{}
}

func (f fileMgtServiceImpl) UploadFile(ctx context.Context, req *pb.UploadFileReq) (*pb.UploadFileRsp, error) {
	//TODO implement me
	panic("implement me")
}
