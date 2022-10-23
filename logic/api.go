package logic

import pb "simpleNetdisk/protoc"

type FileMgt interface {
	UploadFileToBucket(server pb.FileMgt_UploadFileServer) error
}

func NewFileMgt() FileMgt {
	return newFileMgtImpl()
}
