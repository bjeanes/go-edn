package edn

func ParseString(data string) (edn EDN, err error) {
	return ParseBytes([]byte(data))
}
