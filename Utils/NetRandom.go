package Utils

import (
	"encoding/base64"
	"encoding/binary"
)

//NetRandom Enc:
//NetRandom constructor takes seed
//NUM_INIT_RAND is set to 5
//RAND_VAL_MUL is set to 15189
//RAND_VAL_ADD is set to 1
//Seed is set to 3723 in webscocket sendparam

const NUM_INIT_RAND uint32 = 5
const RAND_VAL_MUL uint32 = 15189
const RAND_VAL_ADD uint32 = 1

var Seed uint32

func setSeed(seed uint32) {
	Seed = seed
}

func get16() uint32 {
	ret := (Seed * RAND_VAL_MUL) + RAND_VAL_ADD
	Seed = ret
	return ret
}

func get8() uint32 {
	return get16() >> 8
}

func WSDecrypt(enc string, seed uint32) (plaintext []byte, err error) {
	data, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return nil, err
	}
	setSeed(seed)
	for i := 0; i < len(data); i++ {
		key := get8()
		a := make([]byte, 4)
		binary.LittleEndian.PutUint32(a, key)
		data[i] = data[i] ^ a[0]
	}

	return data, nil
}

func WSEncrypt(data []byte, seed uint32) (enc string, err error) {
	setSeed(seed)
	for i := 0; i < len(data); i++ {
		key := get8()
		a := make([]byte, 4)
		binary.LittleEndian.PutUint32(a, key)
		data[i] = data[i] ^ a[0]
	}
	enc = base64.StdEncoding.EncodeToString(data)
	return enc, nil
}
