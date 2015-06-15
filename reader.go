package json

import (
	"bufio"
	"bytes"
)

// Parse "json" from a given reader removing line comment (so it's not a conformal json)
// and white space, return a slice of bytes that can be marshalled into a struct
func JsonReader(r *bufio.Reader) ([]byte, error) {
	var (
		c_tkn []byte = []byte("//")
		s_tkn []byte = []byte("\"")
		skip  bool   = false
		out   []byte
	)

	l, e := r.ReadBytes('\n')
	for e == nil {
		for i := 0; i < len(l); i++ {
			/* skip state is persistent on line change */
			if bytes.Equal(l[i:], c_tkn) && skip == false {
				i = len(l)
			} else {
				out = append(out, l[i])

				if bytes.Equal(l[i:i+1], s_tkn) {
					skip = !skip
				}
			}
		}
		l, e = r.ReadBytes('\n')
	}

	return out, nil
}
