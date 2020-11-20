package main

import "fmt"
import "os"


func main() {
	var key string                //key為選項的操作
	var opt string                //opt為新增菜單的操作
	var meal = []string{"大麥克","雙層吉士牛肉堡", "麥脆雞原味", "麥香魚","薯條大包","可樂","冰紅茶"}                      //將meal初始化(基本菜單值)
	 for{
		 
		f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)                                  //打開test.txt
		if err != nil {
			fmt.Println("os OpenFile error: ", err)
			return
		}
		defer f.Close()                                                                                            //讀取檔案全部內容
		
		fmt.Printf("(1).新增餐點\n(2).刪除餐點\n(3).按q離開(結束程式)\n請輸入指令:")
		 //fmt.Scanf("%s", &key)                                         //在mac下只需要這一行，無須處理輸入問題
		 n, err := fmt.Scanf("%s\n", &key)
		
		 if err != nil || n != 1 {                                     //處理輸入重複問題(windows下才有此情況)
			 return
			}
			
		if key=="1" {                                                 //開始處理輸入操作的選項
			fmt.Printf("目前菜單餐點\n")
			for i:=0;i<len(meal);i++{
				fmt.Printf("(%d)."+meal[i]+"\n",i+1)
			}
			fmt.Printf("(q).離開(結束程式)\n(p).回到指令主選單\n輸入你要新增的菜單\n")
 
			//  fmt.Scanf("%s", &opt)                                //在mac下只需要這一行，無須處理輸入問題
			n, err := fmt.Scanf("%s\n", &opt)                       //處理輸入重複問題(windows下才有此情況)
			if err != nil || n != 1 {
				return
			}
			os.Truncate("test.txt",0)                               //將記事本初始化
			if opt=="p"{
				fmt.Printf("返回主選單\n")
			}else if opt=="q"{
				break;
			}else{
				meal=append(meal,opt)
				fmt.Printf("你的菜單有:\n")

			for i:=0;i<len(meal);i++{
				f.WriteString(meal[i]+"\n")                       //將新增餐點寫入記事本內
				fmt.Printf("(%d)."+meal[i]+"\n",i+1)
			}
			fmt.Printf("-------------------------------------\n")
			}
			
		}else if key =="2"{
			if len(meal)>0{
				fmt.Printf("你目前的菜單有:\n")
				for i:=0;i<len(meal);i++{
					fmt.Printf("(%d)."+meal[i]+"\n",i+1)
				}
				fmt.Printf("(q).離開(結束程式)\n(p).回到指令主選單\n請輸入餐點名稱\n")
				//fmt.Scanf("%s", &opt)                                 //在mac下只需要這一行，無須處理輸入問題
				
				n, err := fmt.Scanf("%s\n", &opt)
				if err != nil || n != 1 {                               //處理輸入重複問題(windows下才有此情況)
					return
				}
				
				if opt=="p"{
					fmt.Printf("返回主選單\n")
				}else if opt=="q"{
					break;
				}else{
					for i:=0;i<len(meal);i++{                              //將所有的菜單資訊列印出來
					if meal[i]== opt{                                  //尋找菜單資訊有沒有符合opt
						meal=append(meal[:i],meal[i+1:]...)            //找到後，將菜單的位置重新調整
						fmt.Printf("[刪除餐點成功]\n")
				}
			}
			os.Truncate("test.txt",0)                                    //將記事本初始化
			for i:=0;i<len(meal);i++{                                        
				f.WriteString(meal[i]+"\n")                             //將刪除餐點從新寫入記事本內
				fmt.Printf("(%d)."+meal[i]+"\n",i+1)                        //列印出所有菜單
				}
			}
			fmt.Printf("-------------------------------------\n")
				
			}else{
				fmt.Printf("找不到任何菜單請從新選擇\n")
			}
		}else if key =="q"{                                                //離開
			break
		}
	}
	fmt.Printf("結束\n")
}