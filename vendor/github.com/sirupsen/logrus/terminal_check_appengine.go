// +build appengine gopherjs js

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
