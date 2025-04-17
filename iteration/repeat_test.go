package iteration

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name      string
		character string
		count     int
		want      string
	}{
		{"repeat 0 times", "a", 0, ""},
		{"repeat 1 times", "b", 1, "b"},
		{"repeat 2 times", "c", 2, "cc"},
		{"repeat 3 times", "d", 3, "ddd"},
		{"repeat 4 times", "e", 4, "eeee"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Repeat(tt.character, tt.count)
			assertCorrectMessage(t, got, tt.want)
		})
	}
}

const repeatCount = 5
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}
