package main

import (
	"fmt"

	"./rsa"
)

const (
	PubKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDak3y6bu639sSUv9a3FCbfHHQU8xd8WB+9fxk7ap1RcVuplUenLcD2GIAlUTHNVz6hXz9XDlOgYacDhvqGJ18R+sBos/Fof0Hfb8xxFhG69soALaw2AjGi64e12DYeXM83DW3oP8qg4VRPhIM6EmLFI5lfQh5aXFc/qB7aSbmaVwIDAQAB"
	PriKey = "MIICXAIBAAKBgQDak3y6bu639sSUv9a3FCbfHHQU8xd8WB+9fxk7ap1RcVuplUenLcD2GIAlUTHNVz6hXz9XDlOgYacDhvqGJ18R+sBos/Fof0Hfb8xxFhG69soALaw2AjGi64e12DYeXM83DW3oP8qg4VRPhIM6EmLFI5lfQh5aXFc/qB7aSbmaVwIDAQABAoGBAMNv0rvESvD01csHSU6D10nxYB4HSj2lciP3DDGjX60dxc75KotiDsI9H/+9/ODVHKwfxTvrFH04M6SkwUbI12hN1jt63VruyVrGStTLjxSqbxQsTtyy8bAPZwTO7zulb7UqU+1e79wiLyZB4/TdDgQ1kPqcAdDEzSIkapQ0fRlRAkEA5AiIfgrlC43v0bwPq04rfwKZvBmJdZDSYnDTZfnYCnEiiL2/wfSoTjA7rMaE7d5modQe14Mr4HFIX8gsQBGcPwJBAPViB6wgUhgT0Bg4/KtYuTZzayEpt3D36bCgJJ5nDxgqnVGga01i3oPzJpoFGhtnj5u4OTVLGtBCAoMUvZ7YW+kCQA/E10uADV7DLfZe/uZlvXRHBcKQMYVDwoi0LKo8oMf4Et9J39zVwl3F0Bqej8qzOn2uYYOduocGzm5xNT4o+ZkCQH2DQQ1oEtXB/KUHMSar4UDa1KdH4U7lYCE6KMDlUjE4SvPfrEMPGWZAzzPk/C+cTzaFiq2Z1wNDYGmQCFXhsVECQEHbQCEqad8cdVT5uy3koG8CS24yAMver6hg9tA81zK39zjTl9uG9PfOdfMi72I0C/r9w69VFqLKEnrUI7dl14A="
)

var cipher rsa.Cipher

func init() {
	priKey := fmt.Sprintf(`
-----BEGIN PRIVATE KEY-----
%s
-----END PRIVATE KEY-----
`, PriKey)
	pubKey := fmt.Sprintf(`
-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----
`, PubKey)
	client, err := rsa.NewDefault(priKey, pubKey)

	if err != nil {
		fmt.Println(err)
	}

	cipher = client
}
func main() {

}
