package model

import (
	"gorm.io/gorm"
	"go_bin_blog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}


//新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类下所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article,int){
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?",id).Find(&cateArtList).Error
	if err != nil {
		return nil,errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateArtList,errmsg.SUCCESS
}

//查询单个文章
func GetArtInfo(id int) (Article,int){
	var art Article
	err := db.Where("id = ?",id).Preload("Category").First(&art).Error
	if err != nil {
		return art,errmsg.ERROR_ART_NOT_EXIST
	}
	return art,errmsg.SUCCESS
}


//查询文章列表
func GetArt(pageSize int, pageNum int) []Article {
	var articllist []Article
	err = db.Preload("Categoty").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articllist).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return articllist
}

//编辑文章信息
func UpdateArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
