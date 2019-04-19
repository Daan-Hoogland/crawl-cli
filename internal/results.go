package internal

import (
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type sortSlice2d [][]string

func (c sortSlice2d) Len() int      { return len(c) }
func (c sortSlice2d) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

//Less compare strings by first key of the sub-slice
func (c sortSlice2d) Less(i, j int) bool { return strings.Compare(c[i][0], c[j][0]) == -1 }

//ResultTo2DSlice converts a resultList to a 2d slice to be used in table generation.
func ResultTo2DSlice(res *resultList) [][]string {
	res.Lock()
	defer res.Unlock()

	data := make(sortSlice2d, 0)

	for _, file := range res.fileMatches {
		var submit string
		if file.submitted {
			submit = "✓"
		} else {
			submit = "x"
		}
		data = append(data, []string{file.info.Name(), file.path, submit})
	}

	sort.Sort(data)

	return data
}

//GenerateTable generates a table with given data
func GenerateTable(data [][]string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stderr)
	table.SetHeader([]string{"Name", "Path", "Submitted"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_DEFAULT, tablewriter.ALIGN_DEFAULT, tablewriter.ALIGN_CENTER})
	table.SetAutoMergeCells(false)
	table.SetRowLine(true)
	table.AppendBulk(data)

	return table
}
