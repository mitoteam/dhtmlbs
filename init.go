package dhtmlbs

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
)

func init() {
	dhtmlform.Settings().FormErrorsRenderF = func(fe *dhtmlform.FormErrors) (out dhtml.HtmlPiece) {
		container := dhtml.Div().Class("border border-danger-subtle bg-danger-subtle p-3 mb-3")
		for _, controlErrors := range *fe {
			for _, controlError := range controlErrors.Errors {
				errorOut := dhtml.Div()

				if container.HasChildren() {
					//not first error, so add separator
					errorOut.Class("border-top border-1 mt-1 pt-1")
				}

				errorOut.Append(dhtml.Span().Append("✘").Class("me-1", "fw-bold", "text-danger"))

				if !controlErrors.Label.IsEmpty() {
					errorOut.
						Append(dhtml.Span().Class("fw-bold").Append(controlErrors.Label)).
						Text(":")
				}

				errorOut.Append(controlError)

				container.Append(errorOut)
			}
		}

		out.Append(container)

		return out
	}

	dhtml.Settings().EmptyLabelRendererF = func(label string, span *dhtml.Tag) {
		if label == "" {
			label = "empty"
		}

		span.Append("[" + label + "]").Class("text-muted")
	}
}
