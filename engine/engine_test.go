package engine_test

import (
	hash "crypto"
	"fmt"
	"testing"

	c "github.com/nsheremet/esrp/crypto"
	e "github.com/nsheremet/esrp/engine"
	"github.com/nsheremet/esrp/group"
	"github.com/nsheremet/esrp/value"
)

var vectors = map[string]string{
	"N": "eeaf0ab9adb38dd69c33f80afa8fc5e86072618775ff3c0b9ea2314c9c256576d674df7496ea81d3383b4813d692c6e0e0d5d8e250b98be48e495c1d6089dad15dc7d7b46154d6b6ce8ef4ad69b15d4982559b297bcf1885c529f566660e57ec68edbc3c05726cc02fd4cbf4976eaa9afd5138fe8376435b9fc61d2fc0eb06e3",
	"g": "02",
	"k": "7556aa045aef2cdd07abaf0f665c3e818913186f",
	"x": "94b7555aabe9127cc58ccf4993db6cf84d16c124",
	"a": "60975527035cf2ad1989806f0407210bc81edc04e2762a56afd529ddda2d4393",
	"b": "e487cb59d31ac550471e81f00f6928e01dda08e974a004f49e61f5d105284d20",
	"u": "ce38b9593487da98554ed47d70a7ae5f462ef019",
	"A": "61d5e490f6f1b79547b0704c436f523dd0e560f0c64115bb72557ec44352e8903211c04692272d8b2d1a5358a2cf1b6e0bfcf99f921530ec8e39356179eae45e42ba92aeaced825171e1e8b9af6d9c03e1327f44be087ef06530e69f66615261eef54073ca11cf5858f0edfdfe15efeab349ef5d76988a3672fac47b0769447b",
	"B": "bd0c61512c692c0cb6d041fa01bb152d4916a1e77af46ae105393011baf38964dc46a0670dd125b95a981652236f99d9b681cbf87837ec996c6da04453728610d0c6ddb58b318885d7d82c7f8deb75ce7bd4fbaa37089e6f9c6059f388838e7a00030b331eb76840910440b1b27aaeaeeb4012b7d7665238a8e3fb004b117b58",
	"v": "7e273de8696ffc4f4e337d05b4b375beb0dde1569e8fa00a9886d8129bada1f1822223ca1a605b530e379ba4729fdc59f105b4787e5186f5c671085a1447b52a48cf1970b4fb6f8400bbf4cebfbb168152e08ab5ea53d15c1aff87b2b9da6e04e058ad51cc72bfc9033b564e26480d78e955a5e29e7ab245db2be315e2099afb",
	"S": "b0dc82babcf30674ae450c0287745e7990a3381f63b387aaf271a10d233861e359b48220f7c4693c9ae12b0a6f67809f0876e2d013800d6c41bb59b6d5979b5c00a172b4a2a5903a0bdcaf8a709585eb2afafa8f3499b200210dcc1f10eb33943cd67fc88a2f39a4be5bec4ec0a3212dc346d7e474b29ede8a469ffeca686e5a",
}

var g = int(value.New(vectors["g"]).Int().Int64())
var crypto = c.NewStandard(hash.SHA256)
var grp = group.New(512, g, vectors["N"])
var instance = e.New(crypto, grp)

func TestEngineCalcV(t *testing.T) {
	subj := instance.CalcV(value.New(vectors["x"]))

	if subj.Hex() != vectors["v"] {
		t.Error("hex should be equal")
	}
}

func TestEngineCalcA(t *testing.T) {
	subj := instance.CalcV(value.New(vectors["v"]))

	if subj.Hex() != vectors["A"] {
		t.Error("hex should be equal")
	}
}

func TestEngineCalcB(t *testing.T) {
	crypto = c.NewStandard(hash.SHA1)
	instance = e.New(crypto, grp)

	b := value.New(vectors["b"])
	v := instance.CalcV(value.New(vectors["x"]))

	subj := instance.CalcB(b, v)

	if subj.Hex() != vectors["B"] {
		t.Error("hex should be equal")
	}
}

func TestEngineCalcClientS(t *testing.T) {
	a := value.New(vectors["a"])
	u := value.New(vectors["u"])
	x := value.New(vectors["x"])

	val := instance.CalcV(x)
	bb := instance.CalcB(value.New(vectors["b"]), val)
	subj := instance.CalcClientS(bb, a, x, u)

	fmt.Printf("%s\n", subj.Hex())
	fmt.Printf("%s\n", vectors["S"])

	if subj.Hex() != vectors["S"] {
		t.Error("ClientS should be equal to hex")
	}
}

func TestEngineCalcServerS(t *testing.T) {
	crypto = c.NewStandard(hash.SHA1)
	instance = e.New(crypto, grp)

	aa := instance.CalcA(value.New(vectors["a"]))
	b := value.New(vectors["b"])
	v := instance.CalcV(value.New(vectors["x"]))
	u := value.New(vectors["u"])

	subj := instance.CalcServerS(aa, b, v, u)

	if subj.Hex() != vectors["S"] {
		t.Error("ServerS should be equal to hex")
	}
}
