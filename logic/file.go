package logic

import (
	"io"
	pb "simpleNetdisk/protoc"
	"simpleNetdisk/repo/file"
	"simpleNetdisk/repo/storage"
)

type fileMgtImpl struct {
	storageWriter storage.Writer
}

func (f *fileMgtImpl) UploadFileToBucket(server pb.FileMgt_UploadFileServer) error {
	//使用事务保持一致性
	// 1. 保存元文件信息
	// 2. 保存文件信息
	for {
		req, err := server.Recv()
		if err != nil {
			return err
		}
		// 保存文件
		if err == io.EOF {
			//todo 接入minio存储数据
			return server.SendAndClose(&pb.UploadFileRsp{})
		}
		// 构建文件
		fileRepo := file.NewFile(req.GetName())
		if err = fileRepo.Write(req.GetContent()); err != nil {
			return err
		}
	}
}

func newFileMgtImpl() *fileMgtImpl {
	return &fileMgtImpl{
		storageWriter: storage.NewWriter(),
	}
}
