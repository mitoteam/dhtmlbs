package dhtmlbs

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
	"github.com/mitoteam/mttools"
)

// https://getbootstrap.com/docs/5.3/forms/select/

const selectControlKind = "bs_select"

func init() {
	dhtmlform.RegisterFormControlHandler(selectControlKind, selectFormControlHandler)
}

type SelectFormControlElement struct {
	dhtmlform.FormControlElement

	options  mttools.Values
	options2 map[any]dhtml.HtmlPiece // value => label
}

// force interface implementation
var _ dhtmlform.FormControlElementI = (*SelectFormControlElement)(nil)

func NewSelect(name string) *SelectFormControlElement {
	c := &SelectFormControlElement{
		options: mttools.NewValues(),
	}

	c.FormControlElement = *dhtmlform.NewFormControl(selectControlKind, name)

	c.FormControlElement.SetProp("ref", c) //reference to *SelectFormControlElement to be used in FormControlHandler.RenderF

	return c
}

func (c *SelectFormControlElement) Option(value any, label any) *SelectFormControlElement {
	c.options.Set(mttools.AnyToString(value), label)

	return c
}

// dhtmlform control handler implementation
var selectFormControlHandler = &dhtmlform.FormControlHandler{
	RenderF: func(control *dhtmlform.FormControlElement) (out dhtml.HtmlPiece) {
		selectFC := control.GetProp("ref").(*SelectFormControlElement)

		rootTag := formControlWrapper(control)

		if !control.GetLabel().IsEmpty() {
			rootTag.Append(renderControlLabel(control))
		}

		selectElement := dhtml.NewSelect().Class("form-select").Id(control.GetId()).Attribute("name", control.GetName())

		defaultValue := mttools.AnyToString(control.GetValue())

		for value := range selectFC.options.GetNamesIterator() {
			option := selectElement.Option(value, selectFC.options.Get(value))

			if defaultValue != "" && value == defaultValue {
				option.Selected(true)
			}
		}

		rootTag.Append(selectElement)

		if !control.GetNote().IsEmpty() {
			rootTag.Append(renderControlNote(control))
		}

		out.Append(rootTag)
		return out
	},

	ProcessPostValueF: func(controlData *dhtmlform.FormControlData) {
		//log.Printf("dhtmlbs.SelectFormControlElement.ProcessPostValueF: %+v\n", controlData)
		//controlData.Value = controlData.Value == "on"
	},
}
