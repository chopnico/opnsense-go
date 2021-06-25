package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

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

func PrintJson(i interface{}) error {
	j, err := json.Marshal(i)

	if err != nil {
		return err
	}

	fmt.Printf("%s\n", j)
	return nil
}
