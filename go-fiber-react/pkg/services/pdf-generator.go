package services

import (
	"log"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
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
	addItemList(m)
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

type InvoiceItem struct {
	Item            string
	Description     string
	Quantity        string
	Price           string
	DiscountedPrice string
	Total           string
}

func (i InvoiceItem) GetHeader() core.Row {

	return row.New(10).Add(
		text.NewCol(2, "Item", props.Text{
			Style: fontstyle.Bold},
		),
		text.NewCol(3, "Description", props.Text{
			Style: fontstyle.Bold,
		}),
		text.NewCol(1, "Quantity", props.Text{
			Style: fontstyle.Bold,
		}),
		text.NewCol(2, "Price", props.Text{
			Style: fontstyle.Bold,
		}),
		text.NewCol(2, "Discounted Price", props.Text{
			Style: fontstyle.Bold,
		}),
		text.NewCol(2, "Total", props.Text{
			Style: fontstyle.Bold,
		}),
	)
}

func (o InvoiceItem) GetContent(i int) core.Row {
	r := row.New(5).Add(
		text.NewCol(2, o.Item),
		text.NewCol(3, o.Description),
		text.NewCol(2, o.Quantity),
		text.NewCol(2, o.Price),
		text.NewCol(2, o.DiscountedPrice),
		text.NewCol(2, o.Total),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: &props.Color{Red: 240, Green: 240, Blue: 240},
		})
	}

	return r
}

func getObjects() []InvoiceItem {
	var items []InvoiceItem

	contents := [][]string{
		{"Laptop", "14inc, 16GB", "2", "121212", "21121", "12121"},
		{"Laptop", "14inc, 16GB", "2", "121212", "21121", "12121"},
	}

	for i := 0; i < len(contents); i++ {
		items = append(items, InvoiceItem{
			Item:            contents[i][0],
			Description:     contents[i][1],
			Quantity:        contents[i][2],
			Price:           contents[i][3],
			DiscountedPrice: contents[i][4],
			Total:           contents[i][5],
		})
	}

	return items
}

func addItemList(m core.Maroto) {
	rows, err := list.Build[InvoiceItem](getObjects())

	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(rows...)

}
