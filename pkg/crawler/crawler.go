package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	types "clie/pkg/types"
)

const baseURL = "https://hacker-news.firebaseio.com/v0"
const item = baseURL + "/item/%d.json"
const topstories = baseURL + "/topstories.json"

func GetTopStories(ctx context.Context) ([]int, error) {
	var tops []int

	err := getJSON[[]int](ctx, topstories, &tops)
	if err != nil {
		log.Println("failed to get top stories:", err)
		return nil, err
	}

	return tops, nil
}

func GetItem(ctx context.Context, id int) (*types.Item, error) {
	var i types.Item
	url := fmt.Sprintf(item, id)

	err := getJSON[types.Item](ctx, url, &i)

	if err != nil {
		log.Println("failed to get item:", err)
		return nil, err
	}

	return &i, nil
}

func getJSON[T interface{}](ctx context.Context, url string, x *T) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("failed to create request:", err)
		return err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("failed to request:", err)
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("failed to read response body:", err)
		return err
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		log.Println("Failed to unmarshall response body:", err)
		return err
	}

	return nil
}
