package chatgpt

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

type Config struct {
	MaxTokens        int            `toml:"max_token" json:"max_tokens"`
	Temperature      float32        `toml:"temperature" json:"temperature"`
	TopP             float32        `toml:"top_p"json:"top_p"`
	N                int            `toml:"n" json:"n"`
	Stop             []string       `toml:"stop" json:"stop"`
	PresencePenalty  float32        `toml:"presence_penalty" json:"presence_penalty"`
	FrequencyPenalty float32        `toml:"frequency_penalty" json:"frequency_penalty"`
	LogitBias        map[string]int `toml:"logit_bias" json:"logit_bias"`
	User             string         `toml:"user" json:"user"`
	ProxyUrl         string         `toml:"proxy_url" json:"proxy_url"`
	ApiKey           string         `toml:"api_key" json:"api_key"`
}

type Client struct {
	conf   *Config
	client *openai.Client
}

func (c *Client) GetConfig() *Config {
	return c.conf
}

func (c *Client) GetOpenAIClient() *openai.Client {
	return c.client
}

func (c *Client) CreateChatCompletion(ctx context.Context, role, content string) (rsp openai.ChatCompletionResponse, err error) {
	conf := c.conf

	rsp, err = c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    role,
					Content: content,
				},
			},
			MaxTokens:        conf.MaxTokens,
			Temperature:      conf.Temperature,
			TopP:             conf.TopP,
			N:                conf.N,
			Stream:           false,
			Stop:             conf.Stop,
			PresencePenalty:  conf.PresencePenalty,
			FrequencyPenalty: conf.FrequencyPenalty,
			LogitBias:        conf.LogitBias,
			User:             conf.User,
		},
	)

	if err != nil {
		return rsp, err
	}

	return
}

func (c *Client) CreateCompletion(ctx context.Context, prompt string) (rsp openai.CompletionResponse, err error) {
	conf := c.conf

	rsp, err = c.client.CreateCompletion(
		ctx,
		openai.CompletionRequest{
			Model:            openai.GPT3TextDavinci003,
			Prompt:           prompt,
			MaxTokens:        1500,
			Temperature:      conf.Temperature,
			TopP:             conf.TopP,
			N:                conf.N,
			Stream:           false,
			Stop:             conf.Stop,
			PresencePenalty:  conf.PresencePenalty,
			FrequencyPenalty: conf.FrequencyPenalty,
			LogitBias:        conf.LogitBias,
			User:             conf.User,
		},
	)

	if err != nil {
		return rsp, err
	}

	return
}

func NewClient(conf *Config) (*Client, error) {
	c := openai.DefaultConfig(conf.ApiKey)
	if conf.ProxyUrl != "" {
		proxyUrl, err := url.Parse(conf.ProxyUrl)

		if err != nil {
			return nil, err
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		c.HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	openaiClient := openai.NewClientWithConfig(c)

	client := &Client{
		conf:   conf,
		client: openaiClient,
	}

	return client, nil
}
