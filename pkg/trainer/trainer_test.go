package trainer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIovecPair(t *testing.T) {
	assert := assert.New(t)

	trainer := &Trainer{
		storageAddress: 0x1000,
	}

	var value int32
	local, remote := trainer.createIovecPair(2, &value)

	assert.NotNil(local.Base, "Base pointer should be set")
	assert.Equal(uint64(4), local.Len, "Length should be size of int32")

	expectedRemoteBase := uintptr(0x1000 + 2*4)
	assert.Equal(expectedRemoteBase, remote.Base, "Remote Base address should be correctly calculated")
	assert.Equal(4, remote.Len, "Remote Length should be size of int32")

	assert.Panics(func() {
		_, _ = trainer.createIovecPair(-1, &value)
	}, "Should panic for negative index")
}
