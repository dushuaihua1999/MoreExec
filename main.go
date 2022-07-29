package main

import (
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"path/filepath"
)

/*
	1.根据传入的ip, 用户名，密码以及要执行的命令，批量执行 利用ssh连接执行
	2.由配置好的csv表格,根据填写的ip地址，以及要更改的机器名字修改xhbs数据库
*/

var (
	mode = flag.String("mode", "", "批量执行")
)

func main() {
	flag.Parse()
	if *mode == "" {
		fmt.Println(" -mode= ", " name or ssh")
		return
	}
	curDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatalln(err)
	}
	if *mode == "mysql" {
		ipName := make(map[string]string)
		f, err := excelize.OpenFile(curDir + "\\" + "name.xlsx")
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		// Get all the rows in the Sheet1.
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, row := range rows {
			for _, colCell := range row {
				
			}
		}
	}
}
