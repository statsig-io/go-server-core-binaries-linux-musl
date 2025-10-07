//go:build linux && musl && arm64

package go_server_core_binaries_linux_musl

import (
	_ "embed"
)

//go:embed linux_musl_aarch64.so
var binaryData []byte

//go:embed linux_musl_aarch64.so.sig
var signatureData []byte

func getBinaryData() []byte {
	return binaryData
}

func getSignatureData() []byte {
	return signatureData
}
