package main

import (
	"github.com/andreburgaud/crypt2go/ecb"
	_ "github.com/xeodou/go-sqlcipher"
	"golang.org/x/crypto/blowfish"
)

func decrypt(ct, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	if err != nil {
		panic(err.Error())
	}
	return pt
}
