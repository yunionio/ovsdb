package types

import (
	"os"
	"testing"
)

func TestParseSchema(t *testing.T) {
	cases := []struct {
		fname string
	}{
		{
			fname: "vswitch.ovsschema",
		},
		{
			fname: "vtep.ovsschema",
		},
		{
			fname: "ovn-nb.ovsschema",
		},
		{
			fname: "ovn-sb.ovsschema",
		},
	}
	for _, c := range cases {
		t.Run(c.fname, func(t *testing.T) {
			f, err := os.Open(c.fname)
			if err != nil {
				t.Fatalf("open: %v", err)
			}
			schema, err := ParseSchema(f)
			if err != nil {
				t.Errorf("parse: %v", err)
			}
			t.Logf("schema: %s", schema.Name)
		})
	}
}
