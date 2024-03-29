package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/e/util"
	"gin_mall_tmp/serializer"
	"strconv"
)

// CartService 创建购物车
type CartService struct {
	Id        uint `form:"id" json:"id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	// 判断有无这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	cart, status, err := cartDao.CreateCart(service.ProductId, uId, service.BossID)
	if status == e.Error {
		return serializer.Response{
			Status: status,
			Msg:    e.GetMsg(status),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossID)
	return serializer.Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

//Show 购物车
func (service *CartService) Show(ctx context.Context, uId string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	userId, _ := strconv.Atoi(uId)
	carts, err := cartDao.ListCartByUserId(uint(userId))
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCarts(carts),
		Msg:    e.GetMsg(code),
	}
}

// Update 修改购物车信息
func (service *CartService) Update(ctx context.Context, cId string) serializer.Response {
	code := e.Success
	cartId, _ := strconv.Atoi(cId)

	cartDao := dao.NewCartDao(ctx)
	err := cartDao.UpdateCartNumById(uint(cartId), service.Num)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 删除购物车
func (service *CartService) Delete(ctx context.Context) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartById(service.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
