package iifs

import (
	"testing"
)

func TestMountPointLocation(t *testing.T) {

	tests := []struct{
		Content []byte
		MountPath string
		Expected string
	}{
		{
			Content: []byte(nil),
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/5d/0r5dfxE0E2Vx1m5kdm0AH2V505HE2kkfCC0mH15m4545f55Cf4mE41kHmf5d4Ck45x",
		},
		{
			Content: []byte{},
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/5d/0r5dfxE0E2Vx1m5kdm0AH2V505HE2kkfCC0mH15m4545f55Cf4mE41kHmf5d4Ck45x",
		},



		{
			Content: []byte("apple"),
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/2x/0r2xf1VCAAdbHCmHb04mdffAkVfff0f5bkmf5x4fbbkCxmdk1A0E4Eb05410dHbCkE",
		},
		{
			Content: []byte("BANANA"),
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/10/0r1025104HmmbE0kAkA51VAdVxC4fkd1AC4C1xH2Ek2Vk2Hd4fkEE0f1EdVx1411df",
		},
		{
			Content: []byte("Cherry"),
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/A0/0rA0241kHC40m5fV0x5E51d4fH0bCkE2C25ddkHmfEVmV2EV4xb55bVxdmbdAA2bkk",
		},
		{
			Content: []byte("dATE"),
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/2x/0r2xEmd02E0f0kf41CddEx0fVVHEmCk1C1VCbCCEH04A1mAb54AbxkHbxVmd4EEmb5",
		},



		{
			Content: []byte("Hello world!"),
			MountPath: "/home/joeblow/workspaces/articles/.ii/content",
			Expected:  "/home/joeblow/workspaces/articles/.ii/content/SHA-256.SHA-256/Vk/0rVkH2kV05m4E0HkbHmkk5VbVE1V4V25CE1HVHVm1EdxdV00VxxE5kCb1xf1C25E2V",
		},
	}

	for testNumber, test := range tests {

		digest32, err := sha256sha256(test.Content)
		if nil != err {
			t.Errorf("For test #%d, expected to get an error, but did not actually get one: (%T) %q", testNumber, err, err)
			continue
		}
		var digest []byte = digest32[:]

		var mountpoint MountPoint
		mountpoint = Mount(test.MountPath)

		location, err := mountpoint.location(digest)
		if nil != err {
			t.Errorf("For test #%d, expected to get an error, but did not actually get one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, location; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Content: %q", test.Content)
			t.Logf("Digest: 0x%x", digest)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
