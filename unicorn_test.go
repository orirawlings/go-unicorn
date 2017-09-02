package unicorn

import (
	"bytes"
	"testing"
)

func TestFspawnAndSspawn(t *testing.T) {
	var buf bytes.Buffer
	n, err := Fspawn(&buf)
	if err != nil {
		t.Error(err)
	}
	if n != (&buf).Len() {
		t.Errorf("buf Len not %d, was %d", n, (&buf).Len())
	}
	bs, s := (&buf).String(), Sspawn()
	if bs != s {
		t.Errorf("Fspawn and Sspawn not equal:\n  Fspawn = %s\nSspawn = %s\n", bs, s)
	}
}

func BenchmarkFspawn(b *testing.B) {
	var buf bytes.Buffer
	Fspawn(&buf)
	for i := 0; i < b.N; i++ {
		(&buf).Reset()
		Fspawn(&buf)
	}
}
