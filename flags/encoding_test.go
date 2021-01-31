package flags

import(
	"testing"
	"github.com/gucchisk/bytestring"
)

func TestNewEncoding(t *testing.T) {
	type TestData struct {
		str string
		expect bytestring.Encoding
	}
	tests := []TestData{
		{
			str: "ascii",
			expect: bytestring.Ascii,
		},
		{
			str: "Ascii",
			expect: bytestring.Ascii,
		},
		{
			str: "hex",
			expect: bytestring.Hex,
		},
		{
			str: "Hex",
			expect: bytestring.Hex,
		},
		{
			str: "base64",
			expect: bytestring.Base64,
		},
		{
			str: "Base64",
			expect: bytestring.Base64,
		},
		{
			str: "base64url",
			expect: bytestring.Base64URL,
		},
		{
			str: "Base64URL",
			expect: bytestring.Base64URL,
		},
	}

	for _, data := range tests {
		enc, err := NewEncoding(data.str)
		if err != nil {
			t.Errorf("err: %v", err)
		}
		if enc != data.expect {
			t.Errorf("error")
		}
	}
}
