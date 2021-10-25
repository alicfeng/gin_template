package helper

import (
	"github.com/minio/minio-go"
	"io"
	"os"
)

// FileExist 判断文件是否存在 /**
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func GetObject(object *minio.Object) (stream []byte) {
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := object.Read(buf)
		if err != nil && err != io.EOF {
			return stream
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		stream = append(stream, buf[:n]...)
	}

	return stream
}
