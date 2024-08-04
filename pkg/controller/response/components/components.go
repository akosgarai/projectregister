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

// DetailValue is the struct for the detail values.
// It holds the value, and if it is a link, the link.
type DetailValue struct {
	Value string
	Link  string
}

// DetailValues is the struct for the detail values.
type DetailValues []*DetailValue

// DetailItem is the struct for the detail items.
// It holds the label and the value.
// The value is a list of DetailValue
type DetailItem struct {
	Label string
	Value *DetailValues
}

// DetailItems is the struct for the detail items.
type DetailItems []*DetailItem

// ListingHeader is the struct for the listing header.
// It contains the header elements of the listing.
type ListingHeader struct {
	Headers []string
}

// ListingRow is the struct for the listing item.
// It contains the values of the item.
type ListingRow struct {
	Columns *ListingColumns
}

// ListingRows is the struct for the listing items.
type ListingRows []*ListingRow

// ListingColumn is the struct for the listing column.
type ListingColumn struct {
	Values *ListingColumnValues
}

// ListingColumns is the struct for the listing columns.
type ListingColumns []*ListingColumn

// ListingColumnValue is the struct for a listing column entry (one column might contain multiple values).
// It contains the value of the column.
// Also contains the link if it is a link.
// On case of the Form set to true, the link is the action of a POST form.
type ListingColumnValue struct {
	Value string
	Link  string
	Form  bool
}

// ListingColumnValues is the struct for the listing column values.
type ListingColumnValues []*ListingColumnValue

// Listing is the struct for the listing response.
// It contains the header block and the list items.
type Listing struct {
	Header *ListingHeader
	Rows   *ListingRows
}

// FormItemOption is the struct for the form item options.
// It contains the value of the option.
type FormItemOption struct {
	Value    string
	Selected bool
}

// FormItem is the struct for the form items.
type FormItem struct {
	Label    string
	Name     string
	Type     string
	Value    string
	Required bool
	// On case of select / checkbox group type we need the options.
	Options map[int64]*FormItemOption
}

// NewFormItem is a constructor for the FormItem struct.
func NewFormItem(label, name, itemType, value string, required bool, options map[int64]string, selectedOptions []int64) *FormItem {
	// convert the options to the FormItemOption
	// and set the selected options.
	formItemOptions := map[int64]*FormItemOption{}
	if options != nil {
		for key, value := range options {
			isSelected := false
			for _, selectedOption := range selectedOptions {
				if selectedOption == key {
					isSelected = true
					break
				}
			}
			formItemOptions[key] = &FormItemOption{
				Value:    value,
				Selected: isSelected,
			}
		}
	}
	return &FormItem{
		Label:    label,
		Name:     name,
		Type:     itemType,
		Value:    value,
		Required: required,
		Options:  formItemOptions,
	}
}

// Form is the struct for a form.
// It contains the form items.
// Also contains the action, method, and the submit button text.
type Form struct {
	Items  []*FormItem
	Action string
	Method string
	Submit string
}
