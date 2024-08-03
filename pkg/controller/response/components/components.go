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

// ContentHeader is the struct for the content header.
// It contains
// - the title of the page
// - the action buttons (links)
type ContentHeader struct {
	Title   string
	Buttons []*Link
}

// NewContentHeader is a constructor for the ContentHeader struct.
func NewContentHeader(title string, buttons []*Link) *ContentHeader {
	return &ContentHeader{
		Title:   title,
		Buttons: buttons,
	}
}
