package chatgpt

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"os"
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

	return
}

func (c *Client) ListModels(ctx context.Context) (rsp openai.ModelsList, err error) {
	rsp, err = c.client.ListModels(ctx)

	return
}

func (c *Client) Edits(ctx context.Context, input, instruction, modelType string) (rsp openai.EditsResponse, err error) {
	conf := c.conf

	req := openai.EditsRequest{
		Input:       input,
		Instruction: instruction,
		N:           conf.N,
		Temperature: conf.Temperature,
		TopP:        conf.TopP,
	}

	mode := "text-davinci-edit-001"
	if modelType == "code" {
		mode = "code-davinci-edit-001"
	}

	req.Model = &mode

	rsp, err = c.client.Edits(ctx, req)

	return
}

func (c *Client) CreateImage(ctx context.Context, prompt, size, format string) (rsp openai.ImageResponse, err error) {
	conf := c.conf

	req := openai.ImageRequest{
		Prompt:         prompt,
		N:              conf.N,
		Size:           c.getImageSize(size),
		ResponseFormat: c.getImageFormat(format),
		User:           conf.User,
	}

	rsp, err = c.client.CreateImage(ctx, req)

	return
}

func (c *Client) getImageSize(size string) string {
	switch size {
	case "big":
		return openai.CreateImageSize1024x1024
	case "middle":
		return openai.CreateImageSize512x512
	default:
		return openai.CreateImageSize256x256

	}
}

func (c *Client) getImageFormat(format string) string {
	if format == "json" {
		return openai.CreateImageResponseFormatB64JSON
	}

	return openai.CreateImageResponseFormatURL
}

func (c *Client) EditImage(ctx context.Context, imageFilePath, maskPath, prompt, size string) (rsp openai.ImageResponse, err error) {
	conf := c.conf

	imageFile, _ := os.Open(imageFilePath)
	mask, _ := os.Open(maskPath)
	defer imageFile.Close()
	defer mask.Close()

	req := openai.ImageEditRequest{
		Image:  imageFile,
		Mask:   mask,
		Prompt: prompt,
		N:      conf.N,
		Size:   c.getImageSize(size),
	}

	rsp, err = c.client.CreateEditImage(ctx, req)

	return
}

func (c *Client) VariationImage(ctx context.Context, image *os.File, size string) (rsp openai.ImageResponse, err error) {
	conf := c.conf

	req := openai.ImageVariRequest{
		Image: image,
		N:     conf.N,
		Size:  c.getImageSize(size),
	}

	rsp, err = c.client.CreateVariImage(ctx, req)

	return
}

func (c *Client) CreateEmbeddings(ctx context.Context, input string) (rsp openai.EmbeddingResponse, err error) {
	conf := c.conf

	req := openai.EmbeddingRequest{
		Input: []string{input},
		Model: openai.DavinciSearchQuery,
		User:  conf.User,
	}

	rsp, err = c.client.CreateEmbeddings(ctx, req)

	return
}

func (c *Client) CreateTranscription(ctx context.Context, filePath, prompt string) (rsp openai.AudioResponse, err error) {
	conf := c.conf

	req := openai.AudioRequest{
		Model:       "whisper-3",
		FilePath:    filePath,
		Prompt:      prompt,
		Temperature: conf.Temperature,
	}

	rsp, err = c.client.CreateTranscription(ctx, req)

	return
}

func (c *Client) CreateTranslation(ctx context.Context, filePath, prompt string) (rsp openai.AudioResponse, err error) {
	conf := c.conf

	req := openai.AudioRequest{
		Model:       "whisper-3",
		FilePath:    filePath,
		Prompt:      prompt,
		Temperature: conf.Temperature,
	}

	rsp, err = c.client.CreateTranslation(ctx, req)

	return
}

func (c *Client) Moderations(ctx context.Context, input string) (rsp openai.ModerationResponse, err error) {
	model := "text-moderation-stable"

	req := openai.ModerationRequest{
		Input: input,
		Model: &model,
	}

	rsp, err = c.client.Moderations(ctx, req)

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
