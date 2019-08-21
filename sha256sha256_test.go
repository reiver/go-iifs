package iifs

import (
	"encoding/hex"

	"testing"
)

func TestSha256Sha256(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected string
	}{
		{
			Data: []byte(nil),
			Expected: "5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456",
		},
		{
			Data: []byte{},
			Expected: "5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456",
		},


		{
			Data: []byte("apple"),
			Expected: "26f17caadb8c38b043dffa97fff0f5b93f564fbb9c63d91a0e4eb05410d8bc9e",
		},
		{
			Data: []byte("BANANA"),
			Expected: "1025104833be09a9a517ad76c4f9d1ac4c1682e927928d4f9ee0f1ed761411df",
		},
		{
			Data: []byte("Cherry"),
			Expected: "a024198c4035f7065e51d4f80bc9e2c25dd983fe7372e746b55b76d3bdaa2b99",
		},
		{
			Data: []byte("dATE"),
			Expected: "26e3d02e0f09f41cdde60f778e3c91c17cbcce804a13ab54ab698b673d4ee3b5",
		},



		{
			Data: []byte("Hello world!"),
			Expected: "7982970534e089b839957b7e174725ce1878731ed6d700766e59cb16f1c25e27",
		},
	}

	for testNumber, test := range tests {

		actual32, err := sha256sha256(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		expectedBytes, err := hex.DecodeString(test.Expected)
		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := string(expectedBytes), string(actual32[:]); expected != actual {
			t.Errorf("For test #%d, the actual digest is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
