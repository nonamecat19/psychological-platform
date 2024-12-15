package repository

import (
	"gorm.io/gorm"
	"server/internal/database"
	"server/internal/model"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		db: database.GetPostgresConnection(),
	}
}

func (r *ArticleRepository) Create(data model.Article) (model.Article, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *ArticleRepository) FindAll() ([]model.Article, error) {
	var articles []model.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *ArticleRepository) GetOne(id uint) (model.Article, error) {
	var article model.Article
	err := r.db.First(&article, id).Error
	return article, err
}

func (r *ArticleRepository) UpdateOne(data model.Article) (model.Article, error) {
	var article model.Article
	err := r.db.First(&article, data.ID).Error
	if err != nil {
		return article, err
	}
	err = r.db.Model(&article).Updates(data).Error
	return article, err
}

func (r *ArticleRepository) DeleteOne(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}
