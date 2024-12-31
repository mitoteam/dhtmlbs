package dhtmlbs

import (
	"slices"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
	"github.com/mitoteam/mttools"
)

// https://getbootstrap.com/docs/5.3/forms/select/

const selectControlKind = "bs_select"
const selectControlDataValuesProp = "values"

func init() {
	dhtmlform.RegisterFormControlHandler(selectControlKind, selectFormControlHandler)
}

type SelectFormControlElement struct {
	dhtmlform.FormControlElement

	options mttools.Values
}

// force interface implementation
var _ dhtmlform.FormControlElementI = (*SelectFormControlElement)(nil)

func NewSelect(name string) *SelectFormControlElement {
	c := &SelectFormControlElement{
		options: mttools.NewValues(),
	}

	c.FormControlElement = *dhtmlform.NewFormControl(selectControlKind, name)

	//reference to *SelectFormControlElement to be used in FormControlHandler.RenderF
	c.FormControlElement.SetProp("ref", c)

	return c
}

func (c *SelectFormControlElement) Option(value string, label any) *SelectFormControlElement {
	var knownValues []string

	if c.GetControlData().HasProp(selectControlDataValuesProp) {
		knownValues = c.GetControlData().GetProp(selectControlDataValuesProp).([]string)
	} else {
		knownValues = make([]string, 0)
	}

	knownValues = append(knownValues, value)

	c.GetControlData().SetProp(selectControlDataValuesProp, knownValues)

	c.options.Set(value, label)

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
		knownValues := controlData.GetProp(selectControlDataValuesProp).([]string)
		//log.Printf("dhtmlbs.SelectFormControlElement.ProcessPostValueF: %+v\n", possibleValues)

		// unknown values protection
		if !slices.Contains(knownValues, mttools.AnyToString(controlData.Value)) {
			controlData.Value = nil
		}
	},
}
