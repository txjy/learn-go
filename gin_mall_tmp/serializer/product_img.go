package serializer

import (
	"gin_mall_tmp/conf"
	"gin_mall_tmp/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productimg []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productimg = append(productimg, product)
	}
	return
}
