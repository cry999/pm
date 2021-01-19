package persistence

import (
	"github.com/volatiletech/null/v8"
)

func nilIfEmpty(s string) null.String {
	return null.NewString(s, s != "")
}
