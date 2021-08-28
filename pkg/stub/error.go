package stub

import (
	"context"

	"github.com/benalucorp/coinbase-commerce-go/pkg/entity"
)

func CreateErrResp(ctx context.Context) entity.ErrResp {
	return entity.ErrResp{
		Error: GetErrDetailResp(ctx),
		Warnings: []string{
			"stub: warning",
		},
	}
}
