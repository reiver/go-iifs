package iifs

import (
	"github.com/reiver/go-bravo16"
	"github.com/reiver/go-digestfs/driver"

	"path/filepath"
)

const (
	algo = "SHA-256.SHA-256"
)

type MountPoint struct {
	path string
	loaded bool
}

func NotMounted() MountPoint {
	return MountPoint{}
}

func Mount(path string) MountPoint {
	return MountPoint{
		loaded: true,
		path: path,
	}
}

func (receiver MountPoint) location(digest []byte) (location string, err error) {
	if NotMounted() == receiver {
		return "", digestfs_driver.ErrNotMounted()
	}

	var b16digest string
	{
		lenDigest := len(digest)
		lenDst    := bravo16.EncodeLiteralLen(lenDigest)

		var dst []byte = make([]byte, lenDst)

		_, err := bravo16.EncodeLiteral(dst, digest)
		if nil != err {
			return "", err
		}

		b16digest = string(dst)
	}

	var bucket string = b16digest[2:4]

	//@TODO: Escape possible "/" in algorithm
	path := filepath.Join(receiver.path, algo, bucket, b16digest)

	return path, nil
}
