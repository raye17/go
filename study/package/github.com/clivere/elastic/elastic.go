package elasticS

import "github.com/olivere/elastic/v7"

var (
	ElasticClient *elastic.Client
)

func NewElasticClient() *elastic.Client {
	ElasticClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return ElasticClient
}
