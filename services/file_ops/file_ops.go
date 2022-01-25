package file_ops

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"lid/services/logger"
	"os"
)

func Size(f os.File) (int64, error) {
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

func MD5Hash(filename string) string {
	log := logger.CreateLogger("file_ops.Hash")
	log.Trace("hashing " + filename)

	h := md5.New()
	f, err := os.Open(filename)
	if err != nil {
		log.Error("something went wrong opening file", err, "file", filename)
		return ""
	}

	buf := make([]byte, 1024*1024)
	for {
		bytesRead, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Error("something went wrong reading from file", err, "file", filename)
			return ""
		}
		h.Write(buf[:bytesRead])
		if err == io.EOF {
			break
		}
	}

	computedHash := hex.EncodeToString(h.Sum(nil))
	log.Info("computed hash", "file", filename, "hash", computedHash)
	return computedHash
}
