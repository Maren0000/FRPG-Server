package Utils

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

func ReadHashMap(data []byte) (ret map[string]any, err error) {
	ret = map[string]any{}
	elemCount := binary.BigEndian.Uint32(data[1:5])
	pos := 5
	i := 0

	for i != int(elemCount) {
		keylen := binary.BigEndian.Uint16(data[pos : pos+2])
		pos += 2
		key := string(data[pos : pos+int(keylen)])
		pos += int(keylen)

		typeId := int(data[pos])
		pos++
		switch typeId {
		case 2:
			val := binary.BigEndian.Uint32(data[pos : pos+4])
			pos += 4
			ret[key] = val
			i++
		case 3:
			bits := binary.LittleEndian.Uint64(data[pos : pos+8])
			val := math.Float64frombits(bits)
			pos += 8
			ret[key] = val
			i++
		case 4:
			if int(data[pos]) == 0 {
				ret[key] = false
			} else {
				ret[key] = true
			}
			pos++
			i++
		case 5:
			strlen := binary.BigEndian.Uint16(data[pos : pos+2])
			pos += 2
			val := string(data[pos : pos+int(strlen)])
			pos += int(strlen)
			ret[key] = val
			i++
		case 6:
			fmt.Println("array has no support")
			i++
		default:
			fmt.Println("Unkown Element!")
		}
	}

	return ret, nil
}

func WriteRequest(data map[string]any) (ret []byte) {
	ret = WriteHashMap(data)
	return ret
}

func WriteArray(data []any) (Req []byte) {
	Req = append(Req, byte(6))
	arryLen := uint32(len(data))
	Req = binary.BigEndian.AppendUint32(Req, arryLen)
	for i, v := range data {
		Req = binary.BigEndian.AppendUint32(Req, uint32(i))
		switch reflect.TypeOf(v).String() {
		case "map[string]interface {}":
			encode := WriteHashMap(v.(map[string]any))
			Req = append(Req, encode...)
		case "string":
			encode := WriteString(v.(string))
			Req = append(Req, encode...)
		}
	}
	return Req
}

func WriteString(data string) (Req []byte) {
	Req = append(Req, byte(5))
	vlen := uint16(len(data))
	Req = binary.BigEndian.AppendUint16(Req, vlen)
	Req = append(Req, []byte(data)...)
	return Req
}

func WriteHashMap(data map[string]any) (Req []byte) {
	Req = append(Req, byte(1))
	mapLen := uint32(len(data))
	Req = binary.BigEndian.AppendUint32(Req, mapLen)
	for k, v := range data {
		klen := uint16(len(k))
		Req = binary.BigEndian.AppendUint16(Req, klen)
		Req = append(Req, []byte(k)...)
		switch reflect.TypeOf(v).String() {
		case "string":
			encode := WriteString(v.(string))
			Req = append(Req, encode...)
		case "[]interface {}":
			encode := WriteArray(v.([]any))
			Req = append(Req, encode...)
		}
	}
	return Req
}
