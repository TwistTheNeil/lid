package file_ops

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"lid/services/logger"
	"os"
)

func MD5Hash(f *os.File) string {
	log := logger.CreateLogger("file_ops.Hash")
	log.Trace("hashing " + f.Name())

	h := md5.New()
	buf := make([]byte, 1024*1024)

	for {
		bytesRead, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Error("something went wrong reading from file", err, "file", f.Name())
			return ""
		}
		h.Write(buf[:bytesRead])
		if err == io.EOF {
			break
		}
	}

	computedHash := hex.EncodeToString(h.Sum(nil))
	log.Info("computed hash", "node", f.Name(), "hash", computedHash)
	return computedHash
}
