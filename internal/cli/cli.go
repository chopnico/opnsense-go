package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/chopnico/output"
	"github.com/olekukonko/tablewriter"
)

func PrintList(i *[]interface{}, p string) {
	var o string

	if p == "" {
		o = output.FormatList(i, nil)
	} else {
		b := strings.Split(p, ",")
		o = output.FormatList(i, b)
	}

	fmt.Print(o)
}

func PrintTable(data [][]string, header []string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader(header)
	t.SetAutoWrapText(false)
	t.SetAutoFormatHeaders(true)
	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetCenterSeparator("")
	t.SetColumnSeparator("")
	t.SetRowSeparator("")
	t.SetHeaderLine(false)
	t.SetBorder(false)
	t.SetTablePadding("\t")
	t.SetNoWhiteSpace(true)
	t.AppendBulk(data)
	t.Render()
}

func PrintJson(i interface{}) {
	var a []interface{}
	a = append(a, i)
	o := output.FormatJson(&a)

	fmt.Print(o)
}
