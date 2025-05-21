package services

import (
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GenerateInvoice() {
	cfg := config.NewBuilder().WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithRightMargin(15).
		WithTopMargin(15).
		WithBottomMargin(15).
		Build()

	m := maroto.New(cfg)
	generate(m)

}

func generate(m core.Maroto) {
	addHeader(m)
	addInvoiceDetails(m)
}

func addHeader(m core.Maroto) {
	m.AddRow(59, image.NewFromFileCol(12, "./pkg/asset/mihoyo.png", props.Rect{
		Center:  true,
		Percent: 75,
	}))

	m.AddRow(20, text.NewCol(12, "Mihoyo", props.Text{
		Top:   5,
		Style: fontstyle.Bold,
		Align: align.Center,
		Size:  16,
	}))

	m.AddRow(20, text.NewCol(12, "Invoice Tesating", props.Text{
		Top:   6,
		Style: fontstyle.Bold,
		Align: align.Center,
		Size:  12,
	}))

}

func addInvoiceDetails(m core.Maroto) {

	m.AddRow(10,
		text.NewCol(6, "Date: "+time.Now().Format("01 Jan 2024"), props.Text{
			Align: align.Left,
			Size:  10,
		}),
		text.NewCol(6, "Invoice #7", props.Text{
			Align: align.Left,
			Size:  10,
		}))

	m.AddRow(10, line.NewCol(12))

}
