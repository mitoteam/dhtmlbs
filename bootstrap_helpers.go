package dhtmlbs

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
)

var classesBlockError dhtml.Classes = dhtml.NewClasses("border border-danger bg-danger")

func renderControlLabel(controlElement *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
	labelElement := dhtml.NewLabel().For(controlElement.GetId()).Class("form-label").
		Append(controlElement.GetLabel())

	if controlElement.IsRequired() {
		labelElement.Append(dhtml.Span().Class("text-danger fw-bold ms-1").Text("*"))
	}

	out.Append(labelElement)

	return out
}

func renderControlNote(controlElement *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
	out.Append(dhtml.Div().Class("small text-muted").Append(controlElement.GetNote()))
	return out
}
