package dhtmlbs

import "github.com/mitoteam/dhtml"

type ListCardElement struct {
	header  dhtml.HtmlPiece
	classes dhtml.Classes

	items []dhtml.HtmlPiece
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*ListCardElement)(nil)

// Single card with items list
// https://getbootstrap.com/docs/5.3/components/card/#list-groups
func NewListCard() *ListCardElement {
	return &ListCardElement{}
}

func (e *ListCardElement) Class(v ...any) *ListCardElement {
	e.classes.Add(v...)
	return e
}

// Appends something to header
func (e *ListCardElement) Header(v ...any) *ListCardElement {
	e.header.Append(v...)
	return e
}

func (c *ListCardElement) GetHeader() *dhtml.HtmlPiece {
	return &c.header
}

// Adds list items
func (e *ListCardElement) Add(item any) *ListCardElement {
	if e.items == nil {
		e.items = make([]dhtml.HtmlPiece, 0, 1)
	}

	e.items = append(e.items, *dhtml.Piece(item))
	return e
}

func (e *ListCardElement) GetTags() dhtml.TagList {
	root := dhtml.Div().Class("card").Class(e.classes)

	if !e.header.IsEmpty() {
		root.Append(dhtml.Div().Class("card-header").Append(e.header))
	}

	if len(e.items) > 0 {
		listElement := dhtml.NewTag("ul").Class("list-group list-group-flush")

		for _, item := range e.items {
			listElement.Append(
				dhtml.NewTag("li").Class("list-group-item").Append(item),
			)
		}

		root.Append(listElement)
	}

	return root.GetTags()
}
