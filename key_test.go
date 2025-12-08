package translate

import "testing"

func TestKey(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "Hello", want: "hello_8b1a"},
		{input: "Hello World", want: "hello_world_b10a"},
		{input: "The quick brown fox jumps over the lazy dog", want: "quick_brown_jumps_over_lazy_9e107d_43"},
		{input: "Expiry Date (MM/YY)", want: "expiry_date_mm_yy_25e40e_19"},
	}

	for _, test := range tests {
		if got := Key(test.input); got != test.want {
			t.Errorf("Key(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
