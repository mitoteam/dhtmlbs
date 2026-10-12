package dhtmlbs

import "github.com/mitoteam/dhtml"

type GridCardsElement struct {
	cards   []dhtml.ElementI
	classes dhtml.Classes
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*GridCardsElement)(nil)

// Set of cards for pages like dashboards
// https://getbootstrap.com/docs/5.3/components/card/#grid-cards
func NewGridCards() *GridCardsElement {
	return &GridCardsElement{}
}

func (e *GridCardsElement) Add(card dhtml.ElementI) *GridCardsElement {
	e.cards = append(e.cards, card)
	return e
}

func (e *GridCardsElement) Class(v ...any) *GridCardsElement {
	e.classes.Add(v...)
	return e
}

// Adds default flex grid classes.
func (e *GridCardsElement) DefaultGrid() *GridCardsElement {
	e.classes.AddFromSet(GridRowColMdClasses, "row-cols-md-2")
	e.classes.AddFromSet(GridRowColXlClasses, "row-cols-xl-3")

	return e
}

func (e *GridCardsElement) GetTags() dhtml.TagList {
	e.classes.AddFromSet(GapClasses, DefaultGapClass)

	root := dhtml.Div().Class("card-list row").Class(e.classes)

	for _, card := range e.cards {
		root.Append(dhtml.Div().Class("col").Append(card))
	}

	return root.GetTags()
}
