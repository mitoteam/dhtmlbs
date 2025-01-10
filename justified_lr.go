package dhtmlbs

import "github.com/mitoteam/dhtml"

// =================== JustifiedLR =========================

// couple of <div> tags justified by applying .d-flex and .justify-content-between classes
type JustifiedLRElement struct {
	class dhtml.Classes
	l     dhtml.HtmlPiece
	r     dhtml.HtmlPiece
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*JustifiedLRElement)(nil)

func NewJustifiedLR() *JustifiedLRElement {
	return &JustifiedLRElement{}
}

// Add class to element wrapper.
func (e *JustifiedLRElement) Class(v ...any) *JustifiedLRElement {
	e.class.Add(v...)
	return e
}

func (e *JustifiedLRElement) GetL() *dhtml.HtmlPiece {
	return &e.l
}

// Add content to the left side.
func (e *JustifiedLRElement) L(v any) *JustifiedLRElement {
	e.l.Append(v)
	return e
}

func (e *JustifiedLRElement) GetR() *dhtml.HtmlPiece {
	return &e.r
}

// Add content to the right side.
func (e *JustifiedLRElement) R(v any) *JustifiedLRElement {
	e.r.Append(v)
	return e
}

func (e *JustifiedLRElement) GetTags() dhtml.TagList {
	return dhtml.Div().Class("d-flex", "justify-content-between", e.class).
		Append(dhtml.Div().Append(e.l)).
		Append(dhtml.Div().Append(e.r)).
		GetTags()
}
