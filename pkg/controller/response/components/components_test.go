package components

import "testing"

// TestNewLink tests the NewLink function.
// It creates a new Link and checks the fields.
func TestNewLink(t *testing.T) {
	text := "test"
	href := "test"
	link := NewLink(text, href)
	if link.Text != text {
		t.Errorf("Text field is not the same as the input.")
	}
	if link.Href != href {
		t.Errorf("Href field is not the same as the input.")
	}
}
