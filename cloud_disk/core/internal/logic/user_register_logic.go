package logic

import (
	"context"
	"errors"
	"log"

	"cloud_disk/core/etc/helper"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// todo: add your logic here and delete this line
	//判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取该邮箱的验证码为空")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	//判断用户名是否存在
	cnt, err := l.svcCtx.Engine.Where("name=?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("用户名已存在")
		return
	}
	//数据入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}

	n, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row:", n)
	return
}
