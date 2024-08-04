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

// TestNewFormItem tests the NewFormItem function.
// It creates a new FormItem and checks the fields.
func TestNewFormItem(t *testing.T) {
	label := "test"
	inputType := "test"
	name := "test"
	value := "test"
	required := true
	options := map[int64]string{
		1: "test",
		2: "test2",
	}
	selected := []int64{1}
	formItem := NewFormItem(label, inputType, name, value, required, options, selected)
	if formItem.Label != label {
		t.Errorf("Label field is not the same as the input.")
	}
	if formItem.Type != inputType {
		t.Errorf("InputType field is not the same as the input.")
	}
	if formItem.Name != name {
		t.Errorf("Name field is not the same as the input.")
	}
	if formItem.Value != value {
		t.Errorf("Value field is not the same as the input.")
	}
	if formItem.Required != required {
		t.Errorf("Required field is not the same as the input.")
	}
	for i, option := range formItem.Options {
		if option.Value != options[i] {
			t.Errorf("Option field is not the same as the input.")
		}
		// Check if the selected option is selected.
		if selected[0] == int64(i) {
			if !option.Selected {
				t.Errorf("Selected field is not the same as the input.")
			}
		}
	}
}
