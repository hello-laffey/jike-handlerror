// 主入口

package main

import (
	"fmt"
	"handlerror/internal/service/dao"
	"os"
)

// main 主函数
// 程序的顶部，使用%+v把堆栈打清楚
func main() {
	if err := dao.QuerySql(); err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Process finish ok\n")
}
