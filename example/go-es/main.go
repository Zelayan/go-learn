package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"math/rand"
	"time"
)

type LoginRecord struct {
	Username  string    `json:"username,omitempty"`
	LoginTime time.Time `json:"login_time"`
	IPAddress string    `json:"ip_address,omitempty"`
}

const indexName = "login_index"

func main() {
	client := client()
	err := createIndex(client, indexName)
	if err != nil {
		panic(err)
	}
	createDocData(client)

}

func client() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return client
}

func search(client *elastic.Client, indexName string) {
	query := elastic.NewBoolQuery()
	// 精确匹配
	query.Must(elastic.NewTermsQuery("username", "user1"))
	do, err := client.Search(indexName).Query(query).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	printHits(do)
}

func printHits(do *elastic.SearchResult) {
	hits := do.Hits.Hits
	for _, hit := range hits {
		var login LoginRecord
		err := json.Unmarshal(hit.Source, &login)
		if err != nil {
			panic(err)
		}
		log.Println(login)
	}
}

func scrollSearch(client *elastic.Client, indexName string) {
	g, ctx := errgroup.WithContext(context.Background())

	// 用于存放查询结果的channel
	hits := make(chan json.RawMessage)

	g.Go(func() error {
		// 定义一个scroll 查询
		// 这里可以定义其他查询条件
		scroll := client.Scroll(indexName).Size(10).KeepAlive("5m")
		for {
			results, err := scroll.Do(context.TODO())
			// 只要查询不结束就一直查询，如果查询失败了就结束查询
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			for _, hit := range results.Hits.Hits {
				select {
				case hits <- hit.Source:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}
	})

	// 用一个携程去消费， 这里其实可以使用多个携程去消费
	g.Go(func() error {
		for hit := range hits {
			var login LoginRecord
			err := json.Unmarshal(hit, &login)
			if err != nil {
				return err
			}
			log.Println(login)
			select {
			default:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		return nil
	})

	// 等待携程结束
	if err := g.Wait(); err != nil {
		panic(err)
	}
}

func fromSizeSearch(client *elastic.Client, indexName string) {
	// 从第10开始，找10个
	do, err := client.Search(indexName).From(10).Size(10).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	printHits(do)
}

// 创建1000条随机数据
func createDocData(client *elastic.Client) {
	for i := 0; i < 1000; i++ {
		record := LoginRecord{
			Username:  fmt.Sprintf("user%d", i),
			LoginTime: time.Now(),
			IPAddress: fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		}
		_, err := client.Index().Index(indexName).BodyJson(record).Do(context.TODO())
		if err != nil {
			panic(err)
		}
	}
}

func createIndex(client *elastic.Client, indexName string) error {
	do, err := client.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		return err
	}
	if !do {
		_, err := client.CreateIndex(indexName).Do(context.TODO())
		if err != nil {
			return err
		}
	}
	return nil
}
