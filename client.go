package yougile_api_wrapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseURL            = "https://yougile.com/api-v2/"
	DefaultContentType = "application/json"
)

type YouGileClient struct {
	CompanyId   string
	BearerToken string
	ctx         context.Context
}

func NewYouGileClient(companyId string, bearerToken string) *YouGileClient {
	return &YouGileClient{
		CompanyId:   companyId,
		BearerToken: bearerToken,
		ctx:         context.Background(),
	}
}

func (c *YouGileClient) WithContext(ctx context.Context) *YouGileClient {
	newClient := *c
	newClient.ctx = ctx
	return &newClient
}

func (c *YouGileClient) Post(path string, kwargs KeyWordArguments, target interface{}) error {
	//params := kwargs.ToURLValues()
	// кодируем переданную структуру в JSON
	payload, err := json.Marshal(target)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	// создаем урл для ендпоинта
	url := fmt.Sprintf("%s%s", BaseURL, path)

	// создаем новый пост запрос с телом запроса
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	// задаем авторизационный заголовок и заголовок типа контента (он единственный который принимает АПИ)
	req.Header.Add("Content-Type", DefaultContentType)
	if c.BearerToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.BearerToken))
	}
	return nil // TODO реализвать метод <c.doRequest> для для фетча запроса

}
