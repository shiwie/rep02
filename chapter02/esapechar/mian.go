package main

import "fmt" //fmt包中提供可视化，输出，输入的函数。
//这是一个main函数，是程序的入口
func main() {
	//如果希望一次性注释 ctrl+/ 第一次注释
	//演示转义字符的使用\t
	fmt.Println("tom\tjack")
	// \n换行字符的使用
	fmt.Println("hello\nwold")
	// \\字符的使用代表一个\
	fmt.Println("D:D:\\英雄联盟\\Game\\DATA\\FINAL\\Champions")
	// \"代表一个"
	fmt.Println("tom说\"i love you\"")
	// \r 回车 从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r张飞")

}
