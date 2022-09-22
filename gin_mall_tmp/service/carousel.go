package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/e/util"
	"gin_mall_tmp/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List() serializer.Response {

	code := e.Success
	crouselsCtx := dao.NewCarouselDao(context.Background())
	carousels, err := crouselsCtx.ListCarousel()
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousel, uint(len(carousels)))
}
