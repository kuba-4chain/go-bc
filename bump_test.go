package bc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jsonExample = `{
		"blockHeight": 814435,
		"path": [
		  {
			"20": {
			  "hash": "0dc75b4efeeddb95d8ee98ded75d781fcf95d35f9d88f7f1ce54a77a0c7c50fe"
			},
			"21": {
			  "hash": "3ecead27a44d013ad1aae40038acbb1883ac9242406808bb4667c15b4f164eac",
			  "txid": true
			}
		  },
		  {
			"11": {
			  "hash": "5745cf28cd3a31703f611fb80b5a080da55acefa4c6977b21917d1ef95f34fbc"
			}
		  },
		  {
			"4": {
			  "hash": "522a096a1a6d3b64a4289ab456134158d8443f2c3b8ed8618bd2b842912d4b57"
			}
		  },
		  {
			"3": {
			  "hash": "191c70d2ecb477f90716d602f4e39f2f81f686f8f4230c255d1b534dc85fa051"
			}
		  },
		  {
			"0": {
			  "hash": "1f487b8cd3b11472c56617227e7e8509b44054f2a796f33c52c28fd5291578fd"
			}
		  },
		  {
			"1": {
			  "hash": "5ecc0ad4f24b5d8c7e6ec5669dc1d45fcb3405d8ce13c0860f66a35ef442f562"
			}
		  },
		  {
			"1": {
			  "hash": "31631241c8124bc5a9531c160bfddb6fcff3729f4e652b10d57cfd3618e921b1"
			}
		  }
		]
	  }`

	hexExample = `fe636d0c0007021400fe507c0c7aa754cef1f7889d5fd395cf1f785dd7de98eed895dbedfe4e5bc70d1502ac4e164f5bc16746bb0868404292ac8318bbac3800e4aad13a014da427adce3e010b00bc4ff395efd11719b277694cface5aa50d085a0bb81f613f70313acd28cf4557010400574b2d9142b8d28b61d88e3b2c3f44d858411356b49a28a4643b6d1a6a092a5201030051a05fc84d531b5d250c23f4f886f6812f9fe3f402d61607f977b4ecd2701c19010000fd781529d58fc2523cf396a7f25440b409857e7e221766c57214b1d38c7b481f01010062f542f45ea3660f86c013ced80534cb5fd4c19d66c56e7e8c5d4bf2d40acc5e010100b121e91836fd7cd5102b654e9f72f3cf6fdbfd0b161c53a9c54b12c841126331`
)

func TestNewBUMPFromStr(t *testing.T) {
	bump, err := NewBUMPFromStr(hexExample)
	assert.NoError(t, err)
	str, err := bump.String()
	assert.NoError(t, err)
	assert.Equal(t, hexExample, str)
}

func TestNewBUMPFromJson(t *testing.T) {
	jBump, err := NewBUMPFromJSON(jsonExample)
	assert.NoError(t, err)
	jStr, err := jBump.String()
	assert.NoError(t, err)
	assert.Equal(t, hexExample, jStr)
}
