package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"devto/graph/generated"
	"devto/graph/model"
	"encoding/json"
	"fmt"
	"log"

	resty "github.com/go-resty/resty/v2"
)

func (r *queryResolver) Articles(ctx context.Context, input *model.ArticlesQuery) ([]*model.Article, error) {
	// Manipulate Query
	queryInterface := make(map[string]interface{})
	var jsonString []byte
	if input != nil {
		jsonString, _ = json.Marshal(*input)
	} else {
		jsonString = []byte("")
	}
	json.Unmarshal(jsonString, &queryInterface)
	queryString := manipulateQuery(queryInterface)
	// Send Request
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(queryString).
		SetHeader("Accept", "application/json").
		ForceContentType("application/json").
		Get("https://dev.to/api/articles")
	if err != nil {
		log.Fatal(err)
	}
	//Parse Data
	var articles []*model.Article
	if err := json.Unmarshal(resp.Body(), &articles); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return articles, nil
}

func (r *queryResolver) Article(ctx context.Context, id int) (*model.Article, error) {
	// Send Request
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		ForceContentType("application/json").
		Get(fmt.Sprintf("https://dev.to/api/articles/%d", id))
	if err != nil {
		log.Fatal(err)
	}
	//Parse Data
	var article *model.Article
	if err := json.Unmarshal(resp.Body(), &article); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return article, nil
}

func (r *queryResolver) LatestArticles(ctx context.Context, input *model.PaginationQuery) ([]*model.Article, error) {
	// Manipulate Query
	queryInterface := make(map[string]interface{})
	var jsonString []byte
	if input != nil {
		jsonString, _ = json.Marshal(*input)
	} else {
		jsonString = []byte("")
	}
	json.Unmarshal(jsonString, &queryInterface)
	queryString := manipulateQuery(queryInterface)
	// Send Request
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(queryString).
		SetHeader("Accept", "application/json").
		ForceContentType("application/json").
		Get("https://dev.to/api/articles/latest")
	if err != nil {
		log.Fatal(err)
	}
	//Parse Data
	var articles []*model.Article
	if err := json.Unmarshal(resp.Body(), &articles); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return articles, nil
}

func (r *queryResolver) VideoArticles(ctx context.Context, input *model.PaginationQuery) ([]*model.VideoArticle, error) {
	// Manipulate Query
	queryInterface := make(map[string]interface{})
	var jsonString []byte
	if input != nil {
		jsonString, _ = json.Marshal(*input)
	} else {
		jsonString = []byte("")
	}
	json.Unmarshal(jsonString, &queryInterface)
	queryString := manipulateQuery(queryInterface)
	// Send Request
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(queryString).
		SetHeader("Accept", "application/json").
		ForceContentType("application/json").
		Get("https://dev.to/api/videos")
	if err != nil {
		log.Fatal(err)
	}
	//Parse Data
	var videoArticles []*model.VideoArticle
	if err := json.Unmarshal(resp.Body(), &videoArticles); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return videoArticles, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func manipulateQuery(input map[string]interface{}) map[string]string {
	query := make(map[string]string)
	for key, element := range input {
		if element != nil {
			query[key] = fmt.Sprintf("%v", element)
			continue
		}
	}
	return query
}
