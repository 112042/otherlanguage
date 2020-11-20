package main

import "fmt"


func main() {
	var key string                //key為選項的操作
	var opt string                //opt為新增菜單的操作
	 var meal []string             //建立一個空菜單meal
	 for{
		 fmt.Printf("請輸入你的操作行為:(1).新增餐點 (2).刪除餐點 (3).按q離開(輸入錯誤的話，請重新輸入~)\n")
		 //fmt.Scanf("%s", &key)                                         //在mac下只需要這一行，無須處理輸入問題
		 n, err := fmt.Scanf("%s\n", &key)
		
		 if err != nil || n != 1 {                                     //處理輸入重複問題(windows下才有此情況)
			 return
			}
			
		if key=="1" {                                                 //開始處理輸入操作的選項
			fmt.Printf("輸入你要新增的菜單\n")
 
			//  fmt.Scanf("%s", &opt)                                //在mac下只需要這一行，無須處理輸入問題
			n, err := fmt.Scanf("%s\n", &opt)                       //處理輸入重複問題(windows下才有此情況)
			if err != nil || n != 1 {
				return
			}
			
			meal=append(meal,opt)
			fmt.Printf("你的菜單有:\n")

			for i:=0;i<len(meal);i++{
				fmt.Printf("第%d："+meal[i]+"\n",i+1)
			}
		}else if key =="2"{
			if len(meal)>0{
				fmt.Printf("你目前的菜單有:\n")
				for i:=0;i<len(meal);i++{
					fmt.Printf("第%d："+meal[i]+"\n",i+1)
				}
				
				fmt.Printf("輸入你要刪除的菜單(若查無菜單請重新選擇)\n")
				//fmt.Scanf("%s", &opt)                                 //在mac下只需要這一行，無須處理輸入問題
				
				n, err := fmt.Scanf("%s\n", &opt)
				if err != nil || n != 1 {                               //處理輸入重複問題(windows下才有此情況)
					return
				}
				
				for i:=0;i<len(meal);i++{                              //將所有的菜單資訊列印出來
					if meal[i]== opt{                                  //尋找菜單資訊有沒有符合opt
						meal=append(meal[:i],meal[i+1:]...)            //找到後，將菜單的位置重新調整
						i--
				}
			}
			
			for i:=0;i<len(meal);i++{                                        //列印出所有菜單
				fmt.Printf("你剩餘的菜單-第%d："+meal[i]+"\n",i+1)
				}
			}else{
				fmt.Printf("找不到任何菜單請從新選擇\n")
			}
		}else if key =="q"{                                                //離開
			break
		}
	}
	fmt.Printf("結束\n")
}