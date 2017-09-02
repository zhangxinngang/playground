package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
	"context"
	//"encoding/json"
	"reflect"
)

func main() {
	fmt.Println("vim-go")
	esClinet,err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))	
	if err != nil{
		fmt.Println(err)
	}
	exits,_ := esClinet.IndexExists("website").Do(context.Background())
	fmt.Println(exits)
// 	mapping := `
// {
// 	"settings":{
// 		"number_of_shards":1,
// 		"number_of_replicas":0
// 	},
// 	"mappings":{
// 		"_default_": {
// 			"_all": {
// 				"enabled": true
// 			}
// 		},
// 		"tweet":{
// 			"properties":{
// 				"user":{
// 					"type":"keyword"
// 				},
// 				"message":{
// 					"type":"text",
// 					"store": true,
// 					"fielddata": true
// 				},
//             "retweets":{
//                 "type":"long"
//             },
// 				"tags":{
// 					"type":"keyword"
// 				},
// 				"location":{
// 					"type":"geo_point"
// 				},
// 				"suggest_field":{
// 					"type":"completion"
// 				}
// 			}
// 		}
// 	}
// }
// `
// 	//es , err := esClinet.CreateIndex("report").Body(mapping).Do(context.Background())
// 	//fmt.Println(es,err)
	// tweet1 := Report{User: "olivere", Message: "Take Five", Retweets: 0}
	// put1, err := esClinet.Index().Index("report").Type("tweet").Id("1").BodyJson(tweet1).Do(context.Background())
	// tweet2 := Report{User: "olivere", Message: "It's a Raggy Waltz"}
	// esClinet.Index().Index("report").Type("tweet").Id("2").BodyJson(tweet2).Do(context.Background())
	// fmt.Sprintln("%#v",put1)
	// get,err := esClinet.Get().Index("report").Type("tweet").Id("1").Do(context.Background())
	// fmt.Println(get,err)

	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := esClinet.Search().
		Index("report").        // search in index "twitter"
		Query(termQuery).        // specify the query
		Sort("user", true).      // sort by "user" field, ascending
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	fmt.Println(err,"err")
	var ttyp Report
	for _,item := range searchResult.Each(reflect.TypeOf(ttyp)){
		t := item.(Report)
    	fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	}
}

type Report struct{
	User string `json:"user"`
	Message string `json:"message"`
	Retweets int `json:"retweets"`
}