package yougile_api_wrapp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL            = "https://yougile.com/api-v2/"
	DefaultContentType = "application/json"
)

type YouGileClient struct {
	_Client     *http.Client
	CompanyId   string
	BearerToken string
	ctx         context.Context
}

func NewYouGileClient(companyId string, bearerToken string) *YouGileClient {
	return &YouGileClient{
		_Client:     http.DefaultClient,
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

func (c *YouGileClient) Put(path string, kwargs Arguments, target interface{}) error {
	params := kwargs.ToURLValues()

	payload, err := json.Marshal(target)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	url := fmt.Sprintf("%s%s", BaseURL, path)
	urlWithKwargs := fmt.Sprintf("%s%s", url, params.Encode())
	fmt.Println(urlWithKwargs)
	fmt.Println(string(payload))
	req, err := http.NewRequest("PUT", urlWithKwargs, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	req.Header.Set("Content-Type", DefaultContentType)
	if c.BearerToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.BearerToken))
	} else {
		return errors.New("no bearer token provided")
	}
	return c.executeRequest(req, target)
}

func (c *YouGileClient) Post(path string, kwargs Arguments, target interface{}) error {
	params := kwargs.ToURLValues()

	// кодируем переданную структуру в JSON
	payload, err := json.Marshal(target)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	// создаем урл для ендпоинта
	url := fmt.Sprintf("%s%s", BaseURL, path)
	urlWithKwargs := fmt.Sprintf("%s%s", url, params.Encode())

	// создаем новый пост запрос с телом запроса
	req, err := http.NewRequest("POST", urlWithKwargs, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	// задаем авторизационный заголовок и заголовок типа контента (он единственный который принимает АПИ)
	req.Header.Add("Content-Type", DefaultContentType)
	if c.BearerToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.BearerToken))
	}
	fmt.Println(string(payload)) // TODO: добавить логирование + ppjson
	fmt.Println(urlWithKwargs)   // TODO: добавить логирование + ppjson
	return c.executeRequest(req, target)

}

func (c *YouGileClient) Get(path string, kwarg Arguments, target interface{}) error {
	params := kwarg.ToURLValues()
	url := fmt.Sprintf("%s%s", BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())
	req, err := http.NewRequest("GET", urlWithParams, nil)

	fmt.Println(urlWithParams) // TODO: добавить логирование + ppjson

	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	if c.BearerToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.BearerToken))
	}
	req.Header.Set("Content-Type", DefaultContentType)
	return c.executeRequest(req, target)
}

func (c *YouGileClient) executeRequest(req *http.Request, target interface{}) error {
	response, err := c._Client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		str, _ := io.ReadAll(response.Body)
		return fmt.Errorf("unexpected response code: %d, %v", response.StatusCode, string(str))
	}
	_bytes, err := io.ReadAll(response.Body)
	fmt.Printf("|Response data: %v|\n|Status code: %v|\n", string(_bytes), response.StatusCode) // TODO: добавить логирование + ppjson
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	err = json.Unmarshal(_bytes, target)
	if err != nil {
		return fmt.Errorf("error unmarshal response: %w", err)
	}
	// TODO сдлеать prettyJson логи
	return nil
}
