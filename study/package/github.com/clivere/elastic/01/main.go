package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	elasticS "study/package/github.com/clivere/elastic"

	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := elasticS.NewElasticClient()
	log.Println("connect elastic success")
	//search(client, "user")
	res := searchFuzzyNames(client, "user", "s", "l")
	for _, item := range res {
		fmt.Println(item)
	}
	// insert(client, Person{Name: "sxy", Age: 18}, "user")
	// insert(client, Person{Name: "sss", Age: 18}, "user")
	// insert(client, Person{Name: "lcx", Age: 18}, "user")

}
func search(client *elastic.Client, index string) {
	query := elastic.NewMatchAllQuery()
	searchResult, err := client.Search(index).Query(query).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	fmt.Printf("Found %d documents\n", searchResult.TotalHits())
	var p Person
	for _, item := range searchResult.Hits.Hits {
		err := json.Unmarshal(item.Source, &p)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Person: %s, %d\n", p.Name, p.Age)
	}
}
func insert(client *elastic.Client, p Person, index string) {
	put1, err := client.Index().
		Index(index).
		BodyJson(p).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index: %s, type: %s\n", put1.Id, put1.Index, put1.Type)
}
func searchFuzzyNames(client *elastic.Client, index string, names ...string) []Person {
	var res []Person
	var result []*elastic.SearchResult
	for _, name := range names {
		query := elastic.NewWildcardQuery("name", "*"+name+"*")
		searchResult, err := client.Search(index).Query(query).Do(context.Background())
		if err != nil {
			panic(err)
		}
		result = append(result, searchResult)
		fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
		fmt.Printf("Found %d documents\n", searchResult.TotalHits())
	}
	for _, item := range result {
		for _, hit := range item.Hits.Hits {
			var p Person
			err := json.Unmarshal(hit.Source, &p)
			if err != nil {
				panic(err)
			}
			res = append(res, p)
		}
	}
	return res
}
