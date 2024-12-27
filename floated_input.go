package dhtmlbs

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
	"github.com/mitoteam/mttools"
)

const floatingInputControlKind = "floatinginput"

func init() {
	dhtmlform.RegisterFormControlHandler(floatingInputControlKind, &dhtmlform.FormControlHandler{
		RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := dhtml.Div().Class("form-floating")

			if control.IsError() {
				rootTag.Class(classesBlockError)
			}

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Class("form-control").
				Attribute("type", control.GetProp("type").(string)).
				Attribute("name", control.Name).Attribute("value", mttools.AnyToString(control.GetValue()))

			//placeholder is required
			if control.GetPlaceholder() == "" {
				inputTag.Attribute("placeholder", control.GetLabel().RawString())
			} else {
				inputTag.Attribute("placeholder", control.GetPlaceholder())
			}

			rootTag.Append(inputTag)

			//label should be after input for floating
			if !control.GetLabel().IsEmpty() {
				rootTag.Append(renderControlLabel(control))
			}

			if !control.GetNote().IsEmpty() {
				rootTag.Append(renderControlNote(control))
			}

			out.Append(rootTag)

			return out
		},
	})
}

func NewFloatingTextInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(floatingInputControlKind, name).SetProp("type", "text")
}

func NewFloatingPasswordInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(floatingInputControlKind, name).SetProp("type", "password")
}
