package bc_test

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/libsv/go-bc"
	"github.com/libsv/go-bt/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	ebh, _ := bc.NewBlockHeaderFromStr("0000002043453154ad6d8209030ada359e07d2ce354cbed1f6169db497a5f2726e0bb51df5bc41a43429c7469dbb3501a186bf1f9238f9e886f84da057e7571c3472d12af33a1561ffff7f2001000000")
	etx0, _ := bt.NewTxFromString("02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05024c0b0101ffffffff0106270000000000002321033ac208f182e7fe982b1c25027ada05e6fc44590e3f862b0a8422eda03ea5951bac00000000")
	etx1, _ := bt.NewTxFromString("020000000353d4f38490033f3baf11135175c011c61db6cb3e1d9c8d5579da464bd6d7500d000000004847304402205069ed8be3ea22953232328f4594b542655211ce103261ec9278900f8e4a7844022017baa239129970ab92dc4f3f18626954a298e179cc41457e94ea26232fa60de741feffffffd6db9360d48d9084e60d9e9e93ee187ec785768fc38a1826224cda54b436c198000000004847304402203a322b5c2145a8c6194f7575684cf877504a08e07c6718b633c1c7a88bfb71f3022079a87efe2bed70d886cd82f7c747b20a148c79f5adcaec1da05cc18df615fcee41feffffff07c023d3e3bc13b64025000002d2c565521b418562ae0e92e18553c5fafbc781010000006b483045022100abd8d9aed279921efe7be9fd9e24ff2e80b223106355a2e67ecb545cdfbfbf1002207c3861d13bbb08b4aa8e6d5f075f7505a70b98469c4b586c1674bd62b73cf8f2412102d86a9727d885baa389532bba48e37fc529c797939204c78d441a122b2f7a5c32feffffff02bd440f00000000001976a9142621c6863e947d83172bc677640d88cbe5b2477d88aca0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588ac4b0b0000")
	etxs := []*bt.Tx{etx0, etx1}
	eb := &bc.Block{
		BlockHeader: ebh,
		Txs:         etxs,
	}

	blockBytes := "0000002043453154ad6d8209030ada359e07d2ce354cbed1f6169db497a5f2726e0bb51df5bc41a43429c7469dbb3501a186bf1f9238f9e886f84da057e7571c3472d12af33a1561ffff7f20010000000202000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05024c0b0101ffffffff0106270000000000002321033ac208f182e7fe982b1c25027ada05e6fc44590e3f862b0a8422eda03ea5951bac00000000020000000353d4f38490033f3baf11135175c011c61db6cb3e1d9c8d5579da464bd6d7500d000000004847304402205069ed8be3ea22953232328f4594b542655211ce103261ec9278900f8e4a7844022017baa239129970ab92dc4f3f18626954a298e179cc41457e94ea26232fa60de741feffffffd6db9360d48d9084e60d9e9e93ee187ec785768fc38a1826224cda54b436c198000000004847304402203a322b5c2145a8c6194f7575684cf877504a08e07c6718b633c1c7a88bfb71f3022079a87efe2bed70d886cd82f7c747b20a148c79f5adcaec1da05cc18df615fcee41feffffff07c023d3e3bc13b64025000002d2c565521b418562ae0e92e18553c5fafbc781010000006b483045022100abd8d9aed279921efe7be9fd9e24ff2e80b223106355a2e67ecb545cdfbfbf1002207c3861d13bbb08b4aa8e6d5f075f7505a70b98469c4b586c1674bd62b73cf8f2412102d86a9727d885baa389532bba48e37fc529c797939204c78d441a122b2f7a5c32feffffff02bd440f00000000001976a9142621c6863e947d83172bc677640d88cbe5b2477d88aca0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588ac4b0b0000"
	b, err := bc.NewBlockFromStr(blockBytes)

	assert.NoError(t, err)
	assert.Equal(t, eb, b)
}

func TestBlockString(t *testing.T) {
	expectedBlock := "000000208340568a93304c2b327d901fde726e26825a753e9d9681697d60f13b5033691540dddb67dc3caf63b5ac5945e62eed5e7b328901c3bad1be775ca773152be5f8023d1561ffff7f20000000000302000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05024e0b0101ffffffff01cc28000000000000232102af5e52d92723981deef3865309f04807a4cb16cc3da8270b203e482c43a370feac00000000020000000372545d8b76a366701abf79c5219a2f70748c2f888e933b82ada34ed070e66d2100000000494830450221009e8c1ec9c0bb567c47e153946c48dbb1c904d892dd149f92721d9fe87b816f1702207584a0fa85d39056a55685e2c7a1ed6f663b995670bc17fefc00b8ed781591d841feffffffef6f13ab6366f7a670869505630fdee12338ef12efbb223e223b44115f3c273100000000484730440220303ebd18633704633c3b92f261173fa833ca0376578e6d54c213d058c42c6716022077ec705a52337011cd7dd86ebcd207e613618b3da1252bae19355ad45cc04acd41feffffffad5cf4c165fde449155b4de8d1eee9f65e9bb66ff7665f4cb4788a38d665adcc010000006b483045022100ac2e344a9ec980b0c2625a5784c17e62ee59b674a146e6268ae56d49016b57e202202e2e7beb60d879148fdb3f0ed98b7b1148780bb31d82794cddc1c4a2f77d1ed5412102b691a69957cf30c1a7ceae9ba719d5f8891662623f0e797146446df73aa83872feffffff02a0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588acbd440f00000000001976a914fe88c4aeccc229c1bf9913e65fc6ff22f6c9d1fe88ac4d0b000002000000038bf51c82898c0f633f3bab38cdc737a4f666a3640c7128151d6d14bfa911aeb9000000004948304502210095cb2822a8ac066e074a06bf299fd4d2724f869e27e85b02365c2ba54da34e6902202191ffa313b9c4cf55d4893a18e99108d20720bafbbc7c5486238c1e502b254f41feffffffbbba0582b6dc50cce76a0b9d5e00e0cb3afa656db5000eeabad69c3c7b045b860000000049483045022100833865334ae594028a00460dd90575047cdfb9e40d3517051f4841a76035898e0220330e1321e99a59481513978d3fcd34db7b178c8f318176a0eccc0cdb308293a141feffffff5e6584b9ccc112673740ad8fe0f98db8b57585da611a727938fc6702c595827f000000006b483045022100e07f8411e6fd3fdc9ebc9360df6a18a45e49ce80f7e34f387930a16f07d3df6202206eba79ebe9e3760bdb21fa0bb10e4087a51bae88af8b16038d27b89256f9529e412103ba0acf181c9c111451fc5201b8008c33348b49f0b8337e6575312a39eb16852ffeffffff02bd440f00000000001976a914b7a6f23683c5570019094d61429c3c9cbe64533088aca0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588ac4d0b0000"

	bh, _ := bc.NewBlockHeaderFromStr("000000208340568a93304c2b327d901fde726e26825a753e9d9681697d60f13b5033691540dddb67dc3caf63b5ac5945e62eed5e7b328901c3bad1be775ca773152be5f8023d1561ffff7f2000000000")
	tx0, _ := bt.NewTxFromString("02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05024e0b0101ffffffff01cc28000000000000232102af5e52d92723981deef3865309f04807a4cb16cc3da8270b203e482c43a370feac00000000")
	tx1, _ := bt.NewTxFromString("020000000372545d8b76a366701abf79c5219a2f70748c2f888e933b82ada34ed070e66d2100000000494830450221009e8c1ec9c0bb567c47e153946c48dbb1c904d892dd149f92721d9fe87b816f1702207584a0fa85d39056a55685e2c7a1ed6f663b995670bc17fefc00b8ed781591d841feffffffef6f13ab6366f7a670869505630fdee12338ef12efbb223e223b44115f3c273100000000484730440220303ebd18633704633c3b92f261173fa833ca0376578e6d54c213d058c42c6716022077ec705a52337011cd7dd86ebcd207e613618b3da1252bae19355ad45cc04acd41feffffffad5cf4c165fde449155b4de8d1eee9f65e9bb66ff7665f4cb4788a38d665adcc010000006b483045022100ac2e344a9ec980b0c2625a5784c17e62ee59b674a146e6268ae56d49016b57e202202e2e7beb60d879148fdb3f0ed98b7b1148780bb31d82794cddc1c4a2f77d1ed5412102b691a69957cf30c1a7ceae9ba719d5f8891662623f0e797146446df73aa83872feffffff02a0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588acbd440f00000000001976a914fe88c4aeccc229c1bf9913e65fc6ff22f6c9d1fe88ac4d0b0000")
	tx2, _ := bt.NewTxFromString("02000000038bf51c82898c0f633f3bab38cdc737a4f666a3640c7128151d6d14bfa911aeb9000000004948304502210095cb2822a8ac066e074a06bf299fd4d2724f869e27e85b02365c2ba54da34e6902202191ffa313b9c4cf55d4893a18e99108d20720bafbbc7c5486238c1e502b254f41feffffffbbba0582b6dc50cce76a0b9d5e00e0cb3afa656db5000eeabad69c3c7b045b860000000049483045022100833865334ae594028a00460dd90575047cdfb9e40d3517051f4841a76035898e0220330e1321e99a59481513978d3fcd34db7b178c8f318176a0eccc0cdb308293a141feffffff5e6584b9ccc112673740ad8fe0f98db8b57585da611a727938fc6702c595827f000000006b483045022100e07f8411e6fd3fdc9ebc9360df6a18a45e49ce80f7e34f387930a16f07d3df6202206eba79ebe9e3760bdb21fa0bb10e4087a51bae88af8b16038d27b89256f9529e412103ba0acf181c9c111451fc5201b8008c33348b49f0b8337e6575312a39eb16852ffeffffff02bd440f00000000001976a914b7a6f23683c5570019094d61429c3c9cbe64533088aca0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588ac4d0b0000")
	txs := []*bt.Tx{tx0, tx1, tx2}
	b := &bc.Block{
		BlockHeader: bh,
		Txs:         txs,
	}

	assert.Equal(t, expectedBlock, b.String())
}

func TestBlockStringAndBytesMatch(t *testing.T) {
	blockStr := "000000208340568a93304c2b327d901fde726e26825a753e9d9681697d60f13b5033691540dddb67dc3caf63b5ac5945e62eed5e7b328901c3bad1be775ca773152be5f8023d1561ffff7f20000000000302000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05024e0b0101ffffffff01cc28000000000000232102af5e52d92723981deef3865309f04807a4cb16cc3da8270b203e482c43a370feac00000000020000000372545d8b76a366701abf79c5219a2f70748c2f888e933b82ada34ed070e66d2100000000494830450221009e8c1ec9c0bb567c47e153946c48dbb1c904d892dd149f92721d9fe87b816f1702207584a0fa85d39056a55685e2c7a1ed6f663b995670bc17fefc00b8ed781591d841feffffffef6f13ab6366f7a670869505630fdee12338ef12efbb223e223b44115f3c273100000000484730440220303ebd18633704633c3b92f261173fa833ca0376578e6d54c213d058c42c6716022077ec705a52337011cd7dd86ebcd207e613618b3da1252bae19355ad45cc04acd41feffffffad5cf4c165fde449155b4de8d1eee9f65e9bb66ff7665f4cb4788a38d665adcc010000006b483045022100ac2e344a9ec980b0c2625a5784c17e62ee59b674a146e6268ae56d49016b57e202202e2e7beb60d879148fdb3f0ed98b7b1148780bb31d82794cddc1c4a2f77d1ed5412102b691a69957cf30c1a7ceae9ba719d5f8891662623f0e797146446df73aa83872feffffff02a0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588acbd440f00000000001976a914fe88c4aeccc229c1bf9913e65fc6ff22f6c9d1fe88ac4d0b000002000000038bf51c82898c0f633f3bab38cdc737a4f666a3640c7128151d6d14bfa911aeb9000000004948304502210095cb2822a8ac066e074a06bf299fd4d2724f869e27e85b02365c2ba54da34e6902202191ffa313b9c4cf55d4893a18e99108d20720bafbbc7c5486238c1e502b254f41feffffffbbba0582b6dc50cce76a0b9d5e00e0cb3afa656db5000eeabad69c3c7b045b860000000049483045022100833865334ae594028a00460dd90575047cdfb9e40d3517051f4841a76035898e0220330e1321e99a59481513978d3fcd34db7b178c8f318176a0eccc0cdb308293a141feffffff5e6584b9ccc112673740ad8fe0f98db8b57585da611a727938fc6702c595827f000000006b483045022100e07f8411e6fd3fdc9ebc9360df6a18a45e49ce80f7e34f387930a16f07d3df6202206eba79ebe9e3760bdb21fa0bb10e4087a51bae88af8b16038d27b89256f9529e412103ba0acf181c9c111451fc5201b8008c33348b49f0b8337e6575312a39eb16852ffeffffff02bd440f00000000001976a914b7a6f23683c5570019094d61429c3c9cbe64533088aca0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588ac4d0b0000"
	b, err := bc.NewBlockFromStr(blockStr)
	assert.NoError(t, err)
	assert.Equal(t, hex.EncodeToString(b.Bytes()), b.String())
}

func TestBlockInvalid(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		expectedBlock string
		expErr        error
	}{
		"empty string": {
			expectedBlock: "",
			expErr:        errors.New("block cannot be empty"),
		},
		"invalid tx": {
			expectedBlock: "000000208340568a93304c2b327d901fde726e26825a753e9d9681697d60f13b5033691540dddb67dc3caf63b5ac5945e62eed5e7b328901c3bad1be775ca773152be5f8023d1561ffff7f20000000000302000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05024e0b0101ffffffff01cc28000000000000232102af5e52d92723981deef3865309f04807a4cb16cc3da8270b203e482c43a370feac00000000020000000372545d8b76a366701abf79c5219a2f70748c2f888e933b82ada34ed070e66d2100000000494830450221009e8c1ec9c0bb567c47e153946c48dbb1c904d892dd149f92721d9fe87b816f1702207584a0fa85d39056a55685e2c7a1ed6f663b995670bc17fefc00b8ed781591d841feffffffef6f13ab6366f7a670869505630fdee12338ef12efbb223e223b44115f3c273100000000484730440220303ebd18633704633c3b92f261173fa833ca0376578e6d54c213d058c42c6716022077ec705a52337011cd7dd86ebcd207e613618b3da1252bae19355ad45cc04acd41feffffffad5cf4c165fde449155b4de8d1eee9f65e9bb66ff7665f4cb4788a38d665adcc010000006b483045022100ac2e344a9ec980b0c2625a5784c17e62ee59b674a146e6268ae56d49016b57e202202e2e7beb60d879148fdb3f0ed98b7b1148780bb31d82794cddc1c4a2f77d1ed5412102b691a69957cf30c1a7ceae9ba719d5f8891662623f0e797146446df73aa83872feffffff02a0860100000000001976a914b85524abf8202a961b847a3bd0bc89d3d4d41cc588acbd440f00000000001976a914fe88c4aeccc229c1bf9913e65fc6ff22f6c9d1fe88ac4d0b000002000000038bf51c82898c0f633f3bab38cdc737a4f666a3640c7128151d6d14bfa911aeb9000000004948304502210095cb2822a8ac066e074a06bf299fd4d2724f869e27e85b02365c2ba54da34e6902202191ffa313b9c4cf55d4893a18e99108d20720bafbbc7c5486238c1e502b254f41feffffffbbba0582b6dc50cce76a0b9d5e00e0cb3afa656db5000eeabad69c3c7b045b860000000049483045022100833865334ae594028a00460dd90575047cdfb9e40d3517051f4841a76035898e0220330e1321e99a59481513978d3fcd34db7b178c8f318176a0eccc0cdb308293a141feffffff5e6584b9ccc112673740ad8fe0f98db8b57585da611a727938fc6702c595827f000000006b483045022100e07f8411e6fd3fdc9ebc9360df6a18a45e49ce80f7e34f387930a16f07d3df6202206eba79ebe9e3760bdb21fa0bb10e4087a51bae88af8b16038d27b89256f9529e412103ba0acf181c9c111451fc5201b8008c33348b49f0b8337e6575312a39eb168123123123123123213123123123213123123123123123213123123c5570019094d61429c3c9cbe64533088aca0860100000000001976a914b85524ab123a961b81321312312321347a3bd0bc89d3d4d41cc588ac4d0b0000",
			expErr:        errors.New("encoding/hex: odd length hex string"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := bc.NewBlockFromStr(test.expectedBlock)
			assert.Error(t, err)
			assert.EqualError(t, err, test.expErr.Error())
		})
	}
}