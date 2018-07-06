package image

import (
	"testing"
)

func TestMax(t *testing.T) {
	if max(1, 2, 3) != 3 {
		t.Error("max error")
	}

	if max(-1, -2, -3, -4) != -1 {
		t.Error("max error")
	}

	if max(1) != 1 {
		t.Error("max error")
	}
}

func TestMin(t *testing.T) {
	if min(1, 2, 3) != 1 {
		t.Error("min error")
	}

	if min(-1, -2, -3, -4) != -4 {
		t.Error("min error")
	}

	if min(1) != 1 {
		t.Error("min error")
	}
}
