package dhtmlbs

import "github.com/mitoteam/dhtml"

//https://getbootstrap.com/docs/5.3/components/card/

type CardElement struct {
	header                        dhtml.HtmlPiece
	body                          dhtml.HtmlPiece
	class, headerClass, bodyClass dhtml.Classes
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*CardElement)(nil)

func NewCard() *CardElement {
	return &CardElement{}
}

// Add card classes
func (e *CardElement) Class(v ...any) *CardElement {
	e.class.Add(v...)
	return e
}

// Appends something to header
func (e *CardElement) Header(v ...any) *CardElement {
	e.header.Append(v...)
	return e
}

func (c *CardElement) GetHeader() *dhtml.HtmlPiece {
	return &c.header
}

// Add card header classes
func (e *CardElement) HeaderClass(v ...any) *CardElement {
	e.headerClass.Add(v...)
	return e
}

// Appends something to body
func (e *CardElement) Body(v ...any) *CardElement {
	e.body.Append(v...)
	return e
}

// Pointer to card's body
func (e *CardElement) GetBody() *dhtml.HtmlPiece {
	return &e.body
}

// Add card body classes
func (e *CardElement) BodyClass(v ...any) *CardElement {
	e.bodyClass.Add(v...)
	return e
}

func (e *CardElement) GetTags() dhtml.TagList {
	root := dhtml.Div().Class("card", e.class)

	if !e.header.IsEmpty() {
		root.Append(dhtml.Div().Class("card-header", e.headerClass).Append(e.header))
	}

	if !e.body.IsEmpty() {
		root.Append(
			dhtml.Div().Class("card-body", e.bodyClass).Append(e.body),
		)
	}

	return root.GetTags()
}
