package dhtmlbs

import (
	"html"

	"github.com/mitoteam/dhtml"
)

//https://getbootstrap.com/docs/5.3/components/buttons/

var (
	BtnSizeClasses = []string{
		"btn-sm", "btn-md", "btn-lg",
	}

	BtnVariantClasses = []string{
		"btn-outline-primary", "btn-outline-secondary", "btn-outline-success", "btn-outline-danger", "btn-outline-warning",
		"btn-outline-info", "btn-outline-light", "btn-outline-dark",
		"btn-primary", "btn-secondary", "btn-success", "btn-danger", "btn-warning", "btn-info", "btn-light", "btn-dark",
	}
	DefaultBtnVariantClass = BtnVariantClasses[0]
)

type BtnElement struct {
	tag *dhtml.Tag
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*BtnElement)(nil)

func NewBtn() *BtnElement {
	return &BtnElement{
		tag: dhtml.NewTag("a"),
	}
}

func (e *BtnElement) Href(url string) *BtnElement {
	e.tag.Attribute("href", url)
	return e
}

func (e *BtnElement) Label(v any) *BtnElement {
	e.tag.Append(v)
	return e
}

func (e *BtnElement) Title(title string) *BtnElement {
	e.tag.Title(title)
	return e
}

func (e *BtnElement) Confirm(message string) *BtnElement {
	e.tag.AttributeUnsafe("onclick", "return confirm(\""+html.EscapeString(message)+"\");")

	return e
}

func (e *BtnElement) Class(v any) *BtnElement {
	e.tag.Class(v)
	return e
}

func (e *BtnElement) GetTags() dhtml.TagList {
	e.tag.GetClasses().
		Add("btn text-nowrap").
		AddFromSet(BtnVariantClasses, DefaultBtnVariantClass)

	return e.tag.GetTags()
}
