package requester

import (
	"context"
	"net/http"
)

func Send(ctx context.Context) {
	storiesURL := "https://hacker-news.firebaseio.com/v0/topstories.json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, storiesURL, body io.Reader)
}
