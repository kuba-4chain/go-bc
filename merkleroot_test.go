package bc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/libsv/go-bc"
	"github.com/libsv/go-p2p/chaincfg/chainhash"
)

func TestBuildMerkleRoot(t *testing.T) {
	txids := []string{
		"b6d4d13aa08bb4b6cdb3b329cef29b5a5d55d85a85c330d56fddbce78d99c7d6",
		"426f65f6a6ce79c909e54d8959c874a767db3076e76031be70942b896cc64052",
		"adc23d36cc457d5847968c2e4d5f017a6f12a2f165102d10d2843f5276cfe68e",
		"728714bbbddd81a54cae473835ae99eb92ed78191327eb11a9d7494273dcad2a",
		"e3aa0230aa81abd483023886ad12790acf070e2a9f92d7f0ae3bebd90a904361",
		"4848b9e94dd0e4f3173ebd6982ae7eb6b793de305d8450624b1d86c02a5c61d9",
		"912f77eefdd311e24f96850ed8e701381fc4943327f9cf73f9c4dec0d93a056d",
		"397fe2ae4d1d24efcc868a02daae42d1b419289d9a1ded3a5fe771efcc1219d9",
	}

	expected := "1a1e779cd7dfc59f603b4e88842121001af822b2dc5d3b167ae66152e586a6b0"

	root, err := bc.BuildMerkleRoot(txids)

	require.NoError(t, err)
	require.Equal(t, expected, root)
}

func TestTxsToTxIDs(t *testing.T) {
	txs := []string{
		"02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff05021f010101ffffffff0162fe029500000000232102004b59f0c993d90db48b2edc3ad70a12ae395461b4835f05c0809ae65f97e17cac00000000",
		"0200000001984f28af78b9444aa74524144b57e597ae6fa41b98d890b9386405791b93edb1010000006a4730440220615896bb8fe32d0134bdb9f1bb3e11973f0055285e5555fe29436666dceb71e20220143ecb9eb7dcd64ec1b06204fe965fa465e660fc6cf2b8efd5f6c437daa83682412103be25626f7529374909607d5ceb0480d1eea0edc20f008a7d16fed9a531ebd173feffffff0200e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac5e351789000000001976a914d650965966844f9078b404b5bf8025480c180aa188ac1e010000",
		"0200000001567f5fbe0fdbf7f986d53890dbef2853ab3ac484c7a297268083f37f3c419681000000004847304402204e268b71bfc2010204a34a1f12ee5184ad4559b0db39f17c79b0fcf20f13f6a902201f0316188eb9655c4d48c5eba08c5b180a9cdd6cdde152b76e742febd388246241feffffff0200e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac40170d8f000000001976a914effc2aa10ef8cbc60efc2b6da372c3f95ad10b5888ac1e010000",
		"0200000001d9cf788dcfc011efe7087a91065a97b78fc3e5b0dafc09a6945c15b83951cf890000000049483045022100a52ac91c41aeb4c0ff82cc1a4fab2c3a10523c6354cb8fb7429c6f90df8826690220521502978174c051fbe675f4fd1112b0e129b625a07d690d14a517da8827705441feffffff0200e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac40170d8f000000001976a91440457151d8715085fe06ae10423d6532f28ca71988ac1e010000",
		"0200000001908fea3e4d233c39aab4aa62ba983ba8a045a2005fb97d0e297e6e1b022589430000000048473044022013c9d5c45f7e0bdbf0015310c2dd94187e499848f52fd381904e311436427bfb022037b932177072210adcdaa15af68c1fe0ce99068cd9c83f22459c1777d680757e41feffffff0200e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac40170d8f000000001976a91477c8960abf6b9edfacd90e4fc72bcc2d2cdf933c88ac1e010000",
		"02000000010f8512d49417c32bc7c4ae0936586864b2f40d4aadb466b11133846a82e13cdb00000000484730440220741df170e4a81eeaabe9d7a038b9bf0122e938cb5278ce32c57041bad716efd802203cf128209a0f38384514bc2200831f39a614bb30982879a9b019d45ad9fac85841feffffff0240170d8f000000001976a91424af82ffd00cbaf6b4073f6998e6b4e43e38b00688ac00e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac1e010000",
		"0200000001b61bee2fe8a3443d00daa912d9416a40d4ecf2ae67ec1b476a707d79fede06de0000000049483045022100f4d40d7662f0eb8bf4d7034ed6ec279253d56dc56a9668bce6113dd3eeb18e30022023726a68a6fd8a5702f10af432790939afa1a0339128ded7fd785d104b493c6141feffffff0240170d8f000000001976a9143ae6dbf899b67f929109ec5accf532f89a67982288ac00e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac1e010000",
		"020000000193ac7106b2af54f326c74ed53bb3ff58d0ce18e19128371a4d43ae367b27acce000000004847304402205c38de0fbb4279a5f62108616de33e33fa681b59ad01c07648ff1795e82da01602203ca1c3f38b4c1b5809602429709805c968b959971d032a223716f2ed2fb5495441feffffff0200e1f505000000001976a914e296a740f5d9ecc22e0a74f9799f54ec44ee215a88ac40170d8f000000001976a914158679fabbb52644256ee2fd69ce6a42145d43a288ac1e010000",
	}

	expected := []string{
		"b6d4d13aa08bb4b6cdb3b329cef29b5a5d55d85a85c330d56fddbce78d99c7d6",
		"426f65f6a6ce79c909e54d8959c874a767db3076e76031be70942b896cc64052",
		"adc23d36cc457d5847968c2e4d5f017a6f12a2f165102d10d2843f5276cfe68e",
		"728714bbbddd81a54cae473835ae99eb92ed78191327eb11a9d7494273dcad2a",
		"e3aa0230aa81abd483023886ad12790acf070e2a9f92d7f0ae3bebd90a904361",
		"4848b9e94dd0e4f3173ebd6982ae7eb6b793de305d8450624b1d86c02a5c61d9",
		"912f77eefdd311e24f96850ed8e701381fc4943327f9cf73f9c4dec0d93a056d",
		"397fe2ae4d1d24efcc868a02daae42d1b419289d9a1ded3a5fe771efcc1219d9",
	}

	txids, err := bc.TxsToTxIDs(txs)

	require.NoError(t, err)
	require.Equal(t, expected, txids)
}

func TestBuildMerkleTreeStoreChainHash(t *testing.T) {
	txids := []string{
		"b6d4d13aa08bb4b6cdb3b329cef29b5a5d55d85a85c330d56fddbce78d99c7d6",
		"426f65f6a6ce79c909e54d8959c874a767db3076e76031be70942b896cc64052",
		"adc23d36cc457d5847968c2e4d5f017a6f12a2f165102d10d2843f5276cfe68e",
		"728714bbbddd81a54cae473835ae99eb92ed78191327eb11a9d7494273dcad2a",
		"e3aa0230aa81abd483023886ad12790acf070e2a9f92d7f0ae3bebd90a904361",
		"4848b9e94dd0e4f3173ebd6982ae7eb6b793de305d8450624b1d86c02a5c61d9",
		"912f77eefdd311e24f96850ed8e701381fc4943327f9cf73f9c4dec0d93a056d",
		"397fe2ae4d1d24efcc868a02daae42d1b419289d9a1ded3a5fe771efcc1219d9",
	}

	transactionHashes := make([]*chainhash.Hash, len(txids))
	for i, txid := range txids {
		transactionHashes[i], _ = chainhash.NewHashFromStr(txid)
	}

	merkleTreeChainStore := bc.BuildMerkleTreeStoreChainHash(transactionHashes)

	actual := merkleTreeChainStore[len(merkleTreeChainStore)-1].String()

	expected := "1a1e779cd7dfc59f603b4e88842121001af822b2dc5d3b167ae66152e586a6b0"
	require.Equal(t, expected, actual)
}

func TestBuildMerkleTreeStoreChainHashDifferentSizes(t *testing.T) {
	txids := []string{
		"b6d4d13aa08bb4b6cdb3b329cef29b5a5d55d85a85c330d56fddbce78d99c7d6",
		"3783b6638131c8e573410597f2418b7c55be00f6c45aee63f5c1c6d04671ef22",
		"8ac670905831ee210f1abd206ca4c468979709564ca27450c5fb6c3ab78886cc",
		"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b",
		"eb1e33e8a81b697b75855af6bfcdbcbf7cbbde9f94962ceaec1ed8af21f5a50f",
		"524b2d27a1e7fbc3a1614fa661e2dcad68462352feeb8bf633deaccfb8aa84f3",
		"1897ccddda13e71fa04d46700db9dfc4f020fa139943fc9f43e6b3066a0a4fac",
		"c664049c48318a7965d7fa61233fc7ee0f4ea556f2f6481be328a50388c019d2",
	}

	expectations := []string{
		"b6d4d13aa08bb4b6cdb3b329cef29b5a5d55d85a85c330d56fddbce78d99c7d6",
		"1e2ead1d2c9c5471e885cdd2842c516deda57eeebe249385a8650fa7b6e9fc5d",
		"a85c9ebf2b24c881eab3b3a81c886e7b882ec706fa95f5d33bfefe1e22fa6158",
		"113447163e01962e56f6e953f7107f034ab58a8bcac785f0f5e39441e99c9ccf",
		"292356bcb20ac8ea6d84afc176ab8dbc08e73dd8c94ff3aca2df2ec7f369f0eb",
		"d68413306c79b03459f02d126bbb0ec2fd55b880f9ad2de18ca004d260231f65",
		"f4ae9da203450b32785dd17374afe94591347794c59c7c025430fc55ed1ff6c2",
		"684f9b477809d2db9f3ac1dbf660ad578c287130f5ff83bf9cbe8bddf2a6061c",
	}

	for i := 1; i <= 8; i++ {
		transactionHashes := make([]*chainhash.Hash, i)
		for idx := range transactionHashes {
			h, _ := chainhash.NewHashFromStr(txids[idx])
			transactionHashes[idx] = h
		}

		merkleTreeChainStore := bc.BuildMerkleTreeStoreChainHash(transactionHashes)

		actual := merkleTreeChainStore[len(merkleTreeChainStore)-1].String()

		expected := expectations[i-1]
		require.Equal(t, expected, actual)
	}
}
