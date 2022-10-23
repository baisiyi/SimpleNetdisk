package gateway

import (
	"simpleNetdisk/logic"
	pb "simpleNetdisk/protoc"
)

type fileMgtServiceImpl struct {
	fileMgt logic.FileMgt
}

func NewFileMgtServiceImpl() *fileMgtServiceImpl {
	return &fileMgtServiceImpl{
		fileMgt: logic.NewFileMgt(),
	}
}

//UploadFile 文件上传
//fileName 文件名
//path 保存的文件目录
//在设想中，每一个用户实体对应minio的一个桶，桶名为用户uin， 文件保存形式为路径名拼接文件名经过hash后写入桶内
//设计两张表，一张为文件表，记录文件信息，文件名，文件路径名，文件类型，文件所有者
//另一张为用户信息表，记录用户信息，用户名，登录账号，后台唯一识别码uin，登录认证信息token
func (f *fileMgtServiceImpl) UploadFile(server pb.FileMgt_UploadFileServer) error {
	return f.fileMgt.UploadFileToBucket(server)
}
