package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; expected: Hello GitHub Actions", result)
	}
}
