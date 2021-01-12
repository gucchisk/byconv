package flags

import (
	"fmt"
	"strings"
	"github.com/gucchisk/bytestring"
)

func NewEncoding(str string) (bytestring.Encoding, error) {
	switch strings.ToLower(str) {
	case "ascii", "":
		return bytestring.Ascii, nil
	case "hex":
		return bytestring.Hex, nil
	case "base64":
		return bytestring.Base64, nil
	}
	return nil, fmt.Errorf("unkown encoding: %s", str)
}
