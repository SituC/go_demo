package main

import (
	"fmt"

	"github.com/extrame/xls"

)

var res [][]string

func main() {
		if xlFile, err := xls.Open("test.xls", "utf-8"); err == nil {
			fmt.Println(xlFile.Author)
			//第一个sheet
			sheet := xlFile.GetSheet(0)
			if sheet.MaxRow != 0 {
				temp := make([][]string, sheet.MaxRow)
				for i := 0; i < int(sheet.MaxRow); i++ {
					row := sheet.Row(i)
					data := make([]string, 0)
					if row.LastCol() > 0 {
						for j := 0; j < row.LastCol(); j++ {
							col := row.Col(j)
							data = append(data, col)
						}
						temp[i] = data
						fmt.Println(data)
					}
				}
				// res = append(res, temp...)

			}
		} else {
			fmt.Println("open_err:", err)
		}
		// fmt.Println("open_err:", res)
}