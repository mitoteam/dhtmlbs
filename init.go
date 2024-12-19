package dhtmlbs

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
)

func init() {
	dhtmlform.Settings().FormErrorsRenderF = func(fe *dhtmlform.FormErrors) (out dhtml.HtmlPiece) {
		/*
			container := dhtml.Div().Class("border p-3 border-danger border-2 mb-3")

			for name, itemErrors := range fd.GetErrors() {
				for _, itemError := range itemErrors {
					errorDiv := dhtml.Div().Class("item-error")

					if container.HasChildren() {
						//not first error, so add separating line
						errorDiv.Class("border-top border-1 mt-1 pt-1")
					}

					errorDiv.Append(Icon("circle-xmark").Class("text-danger").ElementClass("me-1"))

					if name != "" {
						errorDiv.Attribute("data-form-item-name", name).
							Append(dhtml.Span().Class("fw-bold").Append(fd.GetLabel(name))).
							Append(":")
					}

					errorDiv.Append(itemError)

					container.Append(errorDiv)
				}
			}

			out.Append(container)
		*/
		return out
	}

	dhtml.Settings().EmptyLabelRendererF = func(label string, span *dhtml.Tag) {
		if label == "" {
			label = "empty"
		}

		span.Append("[" + label + "]").Class("text-muted")
	}
}
