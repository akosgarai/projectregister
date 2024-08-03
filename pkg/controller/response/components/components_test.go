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

// TestNewContentHeader tests the NewContentHeader function.
// It creates a new ContentHeader and checks the fields.
func TestNewContentHeader(t *testing.T) {
	title := "test"
	buttons := []*Link{
		{Href: "test", Text: "test"},
		{Href: "test2", Text: "test2"},
	}
	header := NewContentHeader(title, buttons)
	if header.Title != title {
		t.Errorf("Title field is not the same as the input.")
	}
	for i, button := range header.Buttons {
		if button.Href != buttons[i].Href {
			t.Errorf("Href field is not the same as the input.")
		}
		if button.Text != buttons[i].Text {
			t.Errorf("Text field is not the same as the input.")
		}
	}
}
