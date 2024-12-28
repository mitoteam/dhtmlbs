package dhtmlbs

import "github.com/mitoteam/dhtml"

type CardListElement struct {
	cards   []*CardElement
	classes dhtml.Classes
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*CardListElement)(nil)

func NewCardList() *CardListElement {
	return &CardListElement{}
}

func (e *CardListElement) Add(card *CardElement) *CardListElement {
	e.cards = append(e.cards, card)
	return e
}

func (e *CardListElement) Class(v any) *CardListElement {
	e.classes.Add(v)
	return e
}

func (e *CardListElement) GetTags() dhtml.TagList {
	e.classes.AddFromSet(GapClasses, DefaultGapClass)
	e.classes.AddFromSet(GridRowColMdClasses, "row-cols-md-2")
	e.classes.AddFromSet(GridRowColXlClasses, "row-cols-xl-3")

	root := dhtml.Div().Class("card-list row").Class(e.classes)

	for _, card := range e.cards {
		root.Append(dhtml.Div().Class("col").Append(card))
	}

	return root.GetTags()
}
