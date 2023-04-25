package encode

import "encoding/base64"

func CodeEncode(src []byte) []byte {
	return encode(src)
}
func encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}
func CodeDecode(str string) string {
	return decode(str)
}
func decode(encoded string) string {
	d, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	decoded := string(d)
	return decoded
}
