package gpt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"search-conversa/biz/model/api/gpt"
	"search-conversa/config"
	entitygpt "search-conversa/internal/entity/gpt"
	"strconv"
	"time"
)

func CreateConversa(ctx context.Context, c *app.RequestContext) (*gpt.ConversaResp, error) {
	id := strconv.FormatInt(time.Now().Unix(), 10)
	conversa := entitygpt.Conversa{
		ID:      id,
		Role:    "system",
		Content: "You are a helpful assistant.",
	}
	err := config.DB.Create(&conversa).Error
	if err != nil {
		hlog.Error("create conversa faild:", err)
		return nil, err
	}

	return &gpt.ConversaResp{
		Id: id,
		Message: &gpt.Message{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
	}, nil
}
