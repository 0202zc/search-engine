package main

import (
	"context"
	"fmt"
	"log"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"searchengine3090ti/kitex_gen/SearchApi/search"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := search.NewClient("example", client.WithHostPorts("127.0.0.1:2022"))
	if err != nil {
		log.Fatal(err)
	}
	//1. 客户发起添加请求
	resp, err := client.Add(context.Background(),
		&searchapi.AddRequest{Id: 10, Text: "电子科技大学", Url: "https://www.uestc.edu.cn/"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	//2. 客户发起查询请求
	// respQ, err := client.Query(context.Background(),
	// 	&searchapi.QueryRequest{QueryText: "上海", FilterText: "", Page: 1, Limit: 10, Order: 0})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(respQ)
	//3. 相关搜索
	respR, err := client.RelatedQuery(context.Background(), &searchapi.RelatedQueryRequest{QueryText: "电"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(respR)
	//4. 查询ID
	respG, err := client.FindID(context.Background(), &searchapi.FindIDRequest{Ids: []int64{1}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(respG.Entries)
}
