package main

import (
	"bytes"
	"fmt"
)

func Quote(str string) string {
	if str == "" {
		return "\"\""
	}

	var (
		cursor    = 0
		strLength = len(str)
		buffer    bytes.Buffer
	)
	buffer.Grow(strLength + 16)

	buffer.WriteString("\"")
	for cursor < strLength {
		nextchar := str[cursor]
		switch nextchar {
		case '\\':
			buffer.WriteString("\\\\") // Substitue \ por \\
		case '"':
			buffer.WriteString("\\\"") // Substitue ' por \'
		case '\r':
			if (cursor+1) < strLength && str[cursor+1] == '\n' {
				cursor++
			}
			buffer.WriteString("\\n") // Substitue quebra de linha por \n
		case '\n':
			buffer.WriteString("\\n") // Substitue quebra de linha por \n
		default:
			buffer.WriteByte(nextchar)
		}
		cursor++
	}
	buffer.WriteString("\"")
	return buffer.String()
}

func main() {
	fmt.Println(Quote(`he say:"1122"`))
}
