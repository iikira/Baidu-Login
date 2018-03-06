package bdcrypto

import (
	"compress/gzip"
	"io"
)

// GZIPCompress GZIP 压缩
func GZIPCompress(src io.Reader, writeTo io.Writer) (err error) {
	w := gzip.NewWriter(writeTo)
	_, err = io.Copy(w, src)
	if err != nil {
		return
	}

	w.Flush()
	return w.Close()
}

// GZIPUncompress GZIP 解压缩
func GZIPUncompress(src io.Reader, writeTo io.Writer) (err error) {
	unReader, err := gzip.NewReader(src)
	if err != nil {
		return err
	}

	_, err = io.Copy(writeTo, unReader)
	if err != nil {
		return
	}

	return unReader.Close()
}
