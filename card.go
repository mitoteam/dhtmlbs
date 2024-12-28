package dhtmlbs

import "github.com/mitoteam/dhtml"

//https://getbootstrap.com/docs/5.3/components/card/

type CardElement struct {
	header dhtml.HtmlPiece
	body   dhtml.HtmlPiece
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*CardElement)(nil)

func NewCard() *CardElement {
	return &CardElement{}
}

// Appends something to header
func (c *CardElement) Header(v any) *CardElement {
	c.header.Append(v)
	return c
}

func (c *CardElement) GetHeader() *dhtml.HtmlPiece {
	return &c.header
}

// Appends something to body
func (c *CardElement) Body(v any) *CardElement {
	c.body.Append(v)
	return c
}

// Pointer to card's body
func (c *CardElement) GetBody() *dhtml.HtmlPiece {
	return &c.body
}

// region Rendering
func (c *CardElement) GetTags() dhtml.TagList {
	root := dhtml.Div().Class("card")

	if !c.header.IsEmpty() {
		root.Append(dhtml.Div().Class("card-header").Append(c.header))
	}

	if !c.body.IsEmpty() {
		root.Append(
			dhtml.Div().Class("card-body").Append(c.body),
		)
	}

	return root.GetTags()
}

//endregion
