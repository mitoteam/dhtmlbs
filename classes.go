package dhtmlbs

import "slices"

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

	MarginClasses = []string{
		"mt-0", "mt-1", "mt-2", "mt-3", "mt-4", "mt-5",
		"mb-0", "mb-1", "mb-2", "mb-3", "mb-4", "mb-5",
		"ms-0", "ms-1", "ms-2", "ms-3", "ms-4", "ms-5",
		"me-0", "me-1", "me-2", "me-3", "me-4", "me-5",
		"mx-0", "mx-1", "mx-2", "mx-3", "mx-4", "mx-5",
		"my-0", "my-1", "my-2", "my-3", "my-4", "my-5",
		"m-0", "m-1", "m-2", "m-3", "m-4", "m-5",
	}

	PaddingClasses = []string{
		"pt-0", "pt-1", "pt-2", "pt-3", "pt-4", "pt-5",
		"pb-0", "pb-1", "pb-2", "pb-3", "pb-4", "pb-5",
		"ps-0", "ps-1", "ps-2", "ps-3", "ps-4", "ps-5",
		"pe-0", "pe-1", "pe-2", "pe-3", "pe-4", "pe-5",
		"px-0", "px-1", "px-2", "px-3", "px-4", "px-5",
		"py-0", "py-1", "py-2", "py-3", "py-4", "py-5",
		"p-0", "p-1", "p-2", "p-3", "p-4", "p-5",
	}

	GridItemClasses = []string{
		"col", "col-1", "col-2", "col-3", "col-4", "col-5", "col-6",
		"col-7", "col-8", "col-9", "col-10", "col-11", "col-12",

		"col-sm-1", "col-sm-2", "col-sm-3", "col-sm-4", "col-sm-5", "col-sm-6",
		"col-sm-7", "col-sm-8", "col-sm-9", "col-sm-10", "col-sm-11", "col-sm-12",

		"col-md-1", "col-md-2", "col-md-3", "col-md-4", "col-md-5", "col-md-6",
		"col-md-7", "col-md-8", "col-md-9", "col-md-10", "col-md-11", "col-md-12",

		"col-lg-1", "col-lg-2", "col-lg-3", "col-lg-4", "col-lg-5", "col-lg-6",
		"col-lg-7", "col-lg-8", "col-lg-9", "col-lg-10", "col-lg-11", "col-lg-12",

		"col-xl-1", "col-xl-2", "col-xl-3", "col-xl-4", "col-xl-5", "col-xl-6",
		"col-xl-7", "col-xl-8", "col-xl-9", "col-xl-10", "col-xl-11", "col-xl-12",

		"col-xxl-1", "col-xxl-2", "col-xxl-3", "col-xxl-4", "col-xxl-5", "col-xxl-6",
		"col-xxl-7", "col-xxl-8", "col-xxl-9", "col-xxl-10", "col-xxl-11", "col-xxl-12",
	}

	GridRowColClasses = []string{
		"row-cols-1", "row-cols-2", "row-cols-3", "row-cols-4", "row-cols-5", "row-cols-6",
		"row-cols-7", "row-cols-8", "row-cols-9", "row-cols-10", "row-cols-11", "row-cols-12",
	}
	GridRowColXsClasses = []string{
		"row-cols-xs-1", "row-cols-xs-2", "row-cols-xs-3", "row-cols-xs-4", "row-cols-xs-5", "row-cols-xs-6",
		"row-cols-xs-7", "row-cols-xs-8", "row-cols-xs-9", "row-cols-xs-10", "row-cols-xs-11", "row-cols-xs-12",
	}
	GridRowColSmClasses = []string{
		"row-cols-sm-1", "row-cols-sm-2", "row-cols-sm-3", "row-cols-sm-4", "row-cols-sm-5", "row-cols-sm-6",
		"row-cols-sm-7", "row-cols-sm-8", "row-cols-sm-9", "row-cols-sm-10", "row-cols-sm-11", "row-cols-sm-12",
	}
	GridRowColMdClasses = []string{
		"row-cols-md-1", "row-cols-md-2", "row-cols-md-3", "row-cols-md-4", "row-cols-md-5", "row-cols-md-6",
		"row-cols-md-7", "row-cols-md-8", "row-cols-md-9", "row-cols-md-10", "row-cols-md-11", "row-cols-md-12",
	}
	GridRowColLgClasses = []string{
		"row-cols-lg-1", "row-cols-lg-2", "row-cols-lg-3", "row-cols-lg-4", "row-cols-lg-5", "row-cols-lg-6",
		"row-cols-lg-7", "row-cols-lg-8", "row-cols-lg-9", "row-cols-lg-10", "row-cols-lg-11", "row-cols-lg-12",
	}
	GridRowColXlClasses = []string{
		"row-cols-xl-1", "row-cols-xl-2", "row-cols-xl-3", "row-cols-xl-4", "row-cols-xl-5", "row-cols-xl-6",
		"row-cols-xl-7", "row-cols-xl-8", "row-cols-xl-9", "row-cols-xl-10", "row-cols-xl-11", "row-cols-xl-12",
	}
	GridRowColXxlClasses = []string{
		"row-cols-xxl-1", "row-cols-xxl-2", "row-cols-xxl-3", "row-cols-xxl-4", "row-cols-xxl-5", "row-cols-xxl-6",
		"row-cols-xxl-7", "row-cols-xxl-8", "row-cols-xxl-9", "row-cols-xxl-10", "row-cols-xxl-11", "row-cols-xxl-12",
	}
	GridRowClasses = slices.Concat(
		[]string{"row-cols-auto"},
		GridRowColClasses, GridRowColXsClasses, GridRowColSmClasses, GridRowColMdClasses, GridRowColLgClasses, GridRowColXlClasses, GridRowColXxlClasses,
	)

	GapClasses = []string{
		"g-1", "g-2", "g-3", "g-4", "g-5",
	}
	DefaultGapClass = GapClasses[2]
)
