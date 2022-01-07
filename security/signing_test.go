package security

import (
	"github.com/forta-protocol/forta-node/protocol"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	signature = "0x815136705413e8608fb33c7eab05057d1c697db2b8f8fc22e4e29c0d980002626a292cf12a863192f162c19576288c937e162301bc79dcbd006b1e76aea264b101"
	signer    = "0xeE0D82ac806efe2b9a0003a27a785458bC67bbf0"
	ref       = "Qmc2Dmb3wAycyeg3E7Nf6AANqDeBhiX4rdSy3ZJqg2PpMP"
)

func TestSignString(t *testing.T) {
	key, err := LoadKeyWithPassphrase("testkey", "Forta123")
	assert.NoError(t, err)

	res, err := SignString(key, ref)
	assert.NoError(t, err)
	assert.Equal(t, signature, res.Signature)
	assert.Equal(t, signer, res.Signer)
}

func TestVerifySignature(t *testing.T) {
	err := VerifySignature([]byte(ref), signer, signature)
	assert.NoError(t, err)
}

func TestSignProtoMessage(t *testing.T) {
	key, err := LoadKeyWithPassphrase("testkey", "Forta123")
	assert.NoError(t, err)

	obj := &protocol.AlertBatch{
		ChainId: 1,
		Parent:  "test",
	}

	sigHex, err := SignProtoMessage(key, obj)
	assert.NoError(t, err)

	err = VerifyProtoSignature(obj, signer, sigHex.Signature)
	assert.NoError(t, err)
}

func TestSignProtoMessage_BadSignature(t *testing.T) {
	obj := &protocol.AlertBatch{
		ChainId: 1,
		Parent:  "test",
	}

	// this signature isn't for this obj, so it should fail
	err := VerifyProtoSignature(obj, signer, signature)

	// should return error
	assert.Error(t, err)
}
