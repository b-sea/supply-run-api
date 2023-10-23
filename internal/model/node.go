package model

import "encoding/base64"

type Kind string

const (
	idSep = ":"
	idLen = 2
)

type ID struct {
	Key  string
	Kind Kind
}

func (m ID) String() string {
	return base64.StdEncoding.EncodeToString([]byte(m.Key + idSep + m.Key))
}
