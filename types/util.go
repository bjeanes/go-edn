package types

import "bytes"

func collectionString(start string, seq Sequence, end string) string {
	var buffer bytes.Buffer
	buffer.WriteString(start)

	seq.Each(func(item Value, i int) {
		if item == nil {
			buffer.WriteString("nil")
		} else {
			buffer.WriteString(item.String())
		}
		if i < seq.Length()-1 {
			buffer.WriteString(" ")
		}
	})

	buffer.WriteString(end)
	return buffer.String()
}
