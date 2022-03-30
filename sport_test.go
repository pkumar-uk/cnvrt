package sport

import (
	"sport"
	"testing"
)

func TestProcessFilePostive(t *testing.T) {
	_, err := sport.ProcessFile("in.txt")
	if err != nil {
		t.Errorf("Should have read")
	}
}
func TestProcessFileNeg(t *testing.T) {
	_, err := sport.ProcessFile("in1.txt")
	if err == nil {
		t.Errorf("Should have Failed")
	}
}

func BenchmarkProcessFilePostive(b *testing.B) {
	_, err := sport.ProcessFile("in.txt")
	if err != nil {
		b.Errorf("Should have read")
	}
}
