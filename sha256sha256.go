package iifs

import (
	"crypto/sha256"
)

// sha256sha256 is basically: sha256(sha256(p))
func sha256sha256(p []byte) ([sha256.Size]byte, error) {

	var digest1 [sha256.Size]byte
	{
		sha256hasher := sha256.New()
		if nil == sha256hasher {
			return [sha256.Size]byte{}, errInternalError("could not create SHA-256 hasher")
		}

		{
			pp := p

			n, err := sha256hasher.Write(pp)
			if nil != err {
				return [sha256.Size]byte{}, errInternalErrorf("problem writing to SHA-256 hasher: %s", err)
			}
			if expected, actual := len(pp), n; expected != actual {
				return [sha256.Size]byte{}, errInternalErrorf("the number of bytes written to the SHA-256 hasher was not what was expected: expected=%d actual=%d", expected, actual)
			}
		}

		digest := sha256hasher.Sum(nil)
		if nil == digest {
			return [sha256.Size]byte{}, errInternalError("received nil digest from SHA-256 subsystem")
		}

		copy(digest1[:], digest)
	}

	var digest2 [sha256.Size]byte
	{
		sha256hasher := sha256.New()
		if nil == sha256hasher {
			return [sha256.Size]byte{}, errInternalError("could not create SHA-256 hasher")
		}

		{
			pp := digest1[:]

			n, err := sha256hasher.Write(pp)
			if nil != err {
				return [sha256.Size]byte{}, errInternalErrorf("problem writing to SHA-256 hasher: %s", err)
			}
			if expected, actual := len(pp), n; expected != actual {
				return [sha256.Size]byte{}, errInternalErrorf("the number of bytes written to the SHA-256 hasher was not what was expected: expected=%d actual=%d", expected, actual)
			}
		}

		digest := sha256hasher.Sum(nil)
		if nil == digest {
			return [sha256.Size]byte{}, errInternalError("received nil digest from SHA-256 subsystem")
		}

		copy(digest2[:], digest)
	}

	return digest2, nil
}
