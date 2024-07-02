package ii

import (
	"encoding/hex"
	"encoding/json"
	"github.com/aviate-labs/agent-go/certification"
	"github.com/aviate-labs/agent-go/principal"
	"testing"
)

func TestVerifyChallenge(t *testing.T) {
	challenge, err := hex.DecodeString("e7875e69ce7beda6fc7b6dfbd9b75be1c6f6d5debae3ae1ed7c7f873de1b6f9f75e9e7dcddcf37efaddcdf6f7b69a7b57377b5ddaef87dee386ddd75e39e9cd39d7d77debc79df1b7b469df36eb8e7cef47b4d5cefa7f5df67dbefc73debdf5c")
	if err != nil {
		t.Fatal(err)
	}
	delegationJSON := []byte(`{
	"delegations": [
		{
			"delegation":{
				"expiration":"17b5b384762bfd21",
				"pubkey":"e7875e69ce7beda6fc7b6dfbd9b75be1c6f6d5debae3ae1ed7c7f873de1b6f9f75e9e7dcddcf37efaddcdf6f7b69a7b57377b5ddaef87dee386ddd75e39e9cd39d7d77debc79df1b7b469df36eb8e7cef47b4d5cefa7f5df67dbefc73debdf5c"
            },
            "signature":"d9d9f7a26b6365727469666963617465590547d9d9f7a3647472656583018301830183024863616e697374657283018301830183018301830182045820640c48458731be868c750243066312f4e06b2bfde48309a3cfd0617ee3c8f3448301820458204042fb2844db206e1724a248eef393f5cb1d22280f298d948fc18e0a408533438301820458208d3dbc5b1ac807eb4f313b91712db94fdf4a50068207719f1cba37771b2ac8ef83024a000000000060002701018301830183024e6365727469666965645f6461746182035820a61cee2397ab0f006060d4a7bf4a9bef463d5b2381c502a6c66a26b6d088b64d820458206ccd6bb31a54761d4a56e9cfd8cba384d5b8fb47184e8ca13cb70e04f2209ace82045820c64354fe1474e905acdcf09f6569cfb29c305d0b06806908f2da5ee9404726bf820458203de781de0811f5a8469166c594f9433d966f686f4f4065ad9395e30bfac153e282045820cb2a94057004ae336fb52ba39117cf90aaadefe02ddfe9205bcc13c8f6150a0282045820bc1f9b4c54f66eb8fc25381e90641ae59ef87c590186355162a52cb4875242cb8204582001f9f57686d9eb1af846b6ee42c48b02289fe9cf134f84d527a000e65e4d7443820458201c1f10e2904ed9819f3cf7e051c473151700ea5b8038bf1413ba894b3afac4608204582045c96fb30bf784be7d9da2f7e41a2fa93f728bf07829da23acad05006286c269820458204ffce0d4d1e2124180daef5447fe496bbec7ef22b53786138b4acf523453fa75830182045820d5523abdfb2963caffc236cfe5a7f30a832b152c2f827d6acdf79ed5bb9a690e83024474696d65820349a1fa9b83afaae6da17697369676e6174757265583092eaf174a665a296e8968d910ab5a6130fb7deca606a68f5903d8e6a4b64a0fc609b7b7f6a68146e6c51b35e367deb8b6a64656c65676174696f6ea2697375626e65745f6964581d2c55b347ecf2686c83781d6c59d1b43e7b4cba8deb6c1b376107f2cd026b636572746966696361746559026ed9d9f7a2647472656583018204582075d2df1ca388b2596be5564ca726dbcadf77bbc535811734b704a8846153be1383018302467375626e657483018301830183018204582035bc207266aa1f9a1b4eea393efe91ae33ed4ce77069ed8e881d86716adf7b6b830182045820f8c3eae0377ee00859223bf1c6202f5885c4dcdc8fd13b1d48c3c838688919bc83018302581d2c55b347ecf2686c83781d6c59d1b43e7b4cba8deb6c1b376107f2cd02830183024f63616e69737465725f72616e67657382035832d9d9f782824a000000000060000001014a00000000006000ae0101824a00000000006000b001014a00000000006fffff010183024a7075626c69635f6b657982035885308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c0503020103610090075120778eb21a530a02bcc763e7f4a192933506966af7b54c10a4d2b24de6a86b200e3440bae6267bf4c488d9a11d0472c38c1b6221198f98e4e6882ba38a5a4e3aa5afce899b7f825ed95adfa12629688073556f2747527213e8d73e40ce8204582036f3cd257d90fb38e42597f193a5e031dbd585b6292793bb04db4794803ce06e82045820028fc5e5f70868254e7215e7fc630dbd29eefc3619af17ce231909e1faf97e9582045820ef8995c410ed405731c9b913f67879e3b6a6b4d659d2746db9a6b47d7e70d3d582045820f9a6810df003d2188a807e8370076bd94a996877ec8bd11aa2c4e1358c01c6ab83024474696d65820349e2c9c9e480f6edd917697369676e61747572655830833724e450e6e1c8848118e82b04c5db3964f0869b6fb52af9bdbf3876435a19c798c03b41d5eb5fd39535c4ab24e70464747265658301820458209a7cc9ffcec2242e2e15b45a4e1fb9983c87c5b7e8badb7b92a891b40382f73683024373696783025820c9f3b4b781360e36240c549029e4b0857a6cc31e7230a680e551cab71aae0df38301820458203e26edaf16f66c93c238503a3d2077176e9ce6f0438940679b22cb31a636bfee83025820f49c0d7056981c0f2fdfaf02d219db038e2c448193bbf19642fbf118a8f4739a820340"
		}
	],
	"publicKey":"303c300c060a2b0601040183b8430102032c000a00000000006000270101f3ffab2278616508ad5ebfa0cb79a21e08dbb7132f6875b95f81e72067f31302"
}`)
	expiration := uint64(1708469015156620577)
	canisterID := principal.MustDecode("fgte5-ciaaa-aaaad-aaatq-cai")
	rootKey, _ := hex.DecodeString(certification.RootKey)

	var dc DelegationChain
	if err := json.Unmarshal(delegationJSON, &dc); err != nil {
		t.Fatal(err)
	}
	if err := dc.VerifyChallenge(
		challenge,
		expiration-42,
		canisterID,
		rootKey,
	); err != nil {
		t.Fatal(err)
	}
}
