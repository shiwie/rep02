package main
import (
	"go_code/familyaccount/utils"
	"fmt"
)
func main () {
	 // 提示用户输入用户名和密码
	 fmt.Print("请输入用户名: ")
	 var username string
	 fmt.Scanln(&username)
	 fmt.Print("请输入密码: ")
	 var password string
	 fmt.Scanln(&password)
  
	 // 创建 FamilyAccount 实例并传递用户名和密码
	 account := utils.NewFamilyAccount(username, password)
  
	 // 运行主程序
	 account.MainMenu()
}
  