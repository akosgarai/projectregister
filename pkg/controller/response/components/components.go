package components

// Link is the struct for a basic clickable text.
// It contains
// - the label of the item
// - the link.
type Link struct {
	Text string
	Href string
}

// NewLink is a constructor for the Link struct.
func NewLink(text, href string) *Link {
	return &Link{
		Text: text,
		Href: href,
	}
}
