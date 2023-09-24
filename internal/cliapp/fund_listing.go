package cliapp

import (
	"context"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ueihvn/fmarket-cli/pkg/fmarketclient"
)

func (c *cliApp) TableFundListing(ctx context.Context) (*tview.Table, error) {
	req := fmarketclient.NewFilterFundsRequestBuilder().
		SetPage(1).
		SetPageSize(100).
		SetSortOrder(fmarketclient.DescSortOrder).
		SetTypes([]string{fmarketclient.NewFundType, fmarketclient.TradingFundType}).
		SetSortField(fmarketclient.AnnualizedReturn36MonthsSortField).
		Build()
	funds, err := c.fc.FilterFunds(context.TODO(), &req)
	if err != nil {
		return nil, err
	}
	table := tview.NewTable().SetBorders(true).SetSelectedStyle(tcell.StyleDefault)
	const nameCol = 0
	const issuerCol = 1
	const navCol = 2
	const nav6monthsCol = 3
	const annualizedReturn36Months = 4

	table.SetCell(0, nameCol,
		tview.NewTableCell("Name").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter),
	)
	table.SetCell(0, navCol,
		tview.NewTableCell("Nav").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter),
	)
	table.SetCell(0, issuerCol,
		tview.NewTableCell("Issuer").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter),
	)
	table.SetCell(0, nav6monthsCol,
		tview.NewTableCell("Nav 6 Months").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter),
	)
	table.SetCell(0, annualizedReturn36Months,
		tview.NewTableCell("Annualized return 3 years").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter),
	)
	rows := funds.Total
	for row := 0; row < rows; row++ {
		table.SetCell(
			row+1,
			nameCol,
			tview.NewTableCell(funds.Rows[row].ShortName).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft),
		)
		table.SetCell(
			row+1,
			issuerCol,
			tview.NewTableCell(funds.Rows[row].Owner.ShortName).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft),
		)
		table.SetCell(
			row+1,
			navCol,
			tview.NewTableCell(strconv.FormatFloat(funds.Rows[row].NAV, 'f', 2, 64)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft),
		)
		table.SetCell(
			row+1,
			nav6monthsCol,
			tview.NewTableCell(strconv.FormatFloat(funds.Rows[row].ProductNavChange.NavTo6Months, 'f', 2, 64)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft),
		)
		table.SetCell(
			row+1,
			annualizedReturn36Months,
			tview.NewTableCell(strconv.FormatFloat(funds.Rows[row].ProductNavChange.AnnualizedReturn36Months, 'f', 2, 64)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft),
		)
	}

	table.Select(0, 0).
		SetFixed(1, 1).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				c.app.Stop()
			}
			if key == tcell.KeyCtrlC {
				c.app.Stop()
			}
			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		}).
		SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		})
	return table, nil
}
