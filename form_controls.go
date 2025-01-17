package dhtmlbs

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
	"github.com/mitoteam/mttools"
)

const (
	inputControlKind         = "bs_input"
	textareaControlKind      = "bs_textarea"
	checkboxControlKind      = "bs_checkbox"
	floatingInputControlKind = "bs_floatinginput"
	submitControlKind        = "bs_submit_btn"
)

var classBlockError dhtml.Classes = dhtml.NewClasses("border border-danger-subtle bg-danger-subtle py-1")

func init() {
	dhtmlform.RegisterFormControlHandler(inputControlKind, &dhtmlform.FormControlHandler{
		RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := formControlWrapper(control)

			if !control.GetLabel().IsEmpty() {
				rootTag.Append(renderControlLabel(control))
			}

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Class("form-control").
				Attribute("type", control.GetProp("type").(string)).
				Attribute("name", control.GetName()).Attribute("value", mttools.AnyToString(control.GetValue()))

			if control.GetPlaceholder() != "" {
				inputTag.Attribute("placeholder", control.GetPlaceholder())
			}

			rootTag.Append(inputTag)

			if !control.GetNote().IsEmpty() {
				rootTag.Append(renderControlNote(control))
			}

			out.Append(rootTag)
			return out
		},
	})

	dhtmlform.RegisterFormControlHandler(floatingInputControlKind, &dhtmlform.FormControlHandler{
		RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := formControlWrapper(control).Class("form-floating")

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Class("form-control").
				Attribute("type", control.GetProp("type").(string)).
				Attribute("name", control.GetName()).Attribute("value", mttools.AnyToString(control.GetValue()))

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

	dhtmlform.RegisterFormControlHandler(textareaControlKind, &dhtmlform.FormControlHandler{
		RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := formControlWrapper(control)

			if !control.GetLabel().IsEmpty() {
				rootTag.Append(renderControlLabel(control))
			}

			rows := control.GetProp("rows")
			if mttools.IsEmpty(rows) {
				rows = 3
			}

			textareaTag := dhtml.NewTag("textarea").Id(control.GetId()).Class("form-control").
				Attribute("rows", mttools.AnyToString(rows)).
				Attribute("name", control.GetName()).Append(control.GetValue())

			if control.GetPlaceholder() != "" {
				textareaTag.Attribute("placeholder", control.GetPlaceholder())
			}

			rootTag.Append(textareaTag)

			if !control.GetNote().IsEmpty() {
				rootTag.Append(renderControlNote(control))
			}

			out.Append(rootTag)
			return out
		},
	})

	dhtmlform.RegisterFormControlHandler(checkboxControlKind, &dhtmlform.FormControlHandler{
		RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
			rootTag := formControlWrapper(control)

			inputTag := dhtml.NewTag("input").Id(control.GetId()).Attribute("type", "checkbox").
				Attribute("name", control.GetName()).Attribute("value", "on")

			if mttools.AnyToBool(control.GetValue()) {
				inputTag.Attribute("checked", "")
			}

			rootTag.Append(inputTag)

			if !control.GetLabel().IsEmpty() {
				rootTag.Append(renderControlLabel(control))
			}

			if !control.GetNote().IsEmpty() {
				rootTag.Append(renderControlNote(control))
			}

			out.Append(rootTag)
			return out
		},

		ProcessPostValueF: func(controlData *dhtmlform.FormControlData) {
			controlData.Value = controlData.Value == "on"
		},
	})

	dhtmlform.RegisterFormControlHandler(submitControlKind, &dhtmlform.FormControlHandler{
		RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
			tag := dhtml.NewTag("button").Attribute("type", "submit").Class("btn btn-success")

			if !mttools.IsEmpty(control.GetValue()) {
				tag.Attribute("name", control.GetName()).Attribute("value", mttools.AnyToString(control.GetValue()))
			}

			if control.GetLabel().IsEmpty() {
				tag.Append("Submit")
			} else {
				tag.Append(control.GetLabel())
			}

			out.Append(tag)
			return out
		},
	})
}

func NewTextInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "text")
}

func NewPasswordInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "password")
}

func NewEmailInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "email")
}

func NewDateInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "date")
}

func NewNumberInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "number")
}

func NewTelInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "tel")
}

func NewTimeInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "time")
}

func NewUrlInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(inputControlKind, name).SetProp("type", "url")
}

func NewFloatingTextInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(floatingInputControlKind, name).SetProp("type", "text")
}

func NewFloatingPasswordInput(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(floatingInputControlKind, name).SetProp("type", "password")
}

func NewTextarea(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(textareaControlKind, name)
}

func NewCheckbox(name string) *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(checkboxControlKind, name)
}

func NewSubmitBtn() *dhtmlform.FormControlElement {
	return dhtmlform.NewFormControl(submitControlKind, "submit")
}

// -------------------- internal helpers --------------------------

func formControlWrapper(controlElement *dhtmlform.FormControlElement) (tag *dhtml.Tag) {
	tag = dhtml.Div()

	var wrapperClass any

	if controlElement.HasProp("wrapper-class") {
		wrapperClass = controlElement.GetProp("wrapper-class")
	} else {
		wrapperClass = "mb-3"
	}

	if !mttools.IsEmpty(wrapperClass) {
		tag.Class(wrapperClass)
	}

	if controlElement.IsError() {
		tag.Class(classBlockError)
	}

	return tag
}

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
