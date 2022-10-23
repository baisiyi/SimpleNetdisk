package storage

type Writer interface {
}

type Reader interface {
}

func NewWriter() Writer {
	return newWriterImpl()
}

func NewReader() Reader {
	return newReaderImpl()
}
