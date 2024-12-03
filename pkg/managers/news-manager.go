package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"

	"gorm.io/gorm"
)

type NewsManager struct {
	repo *repositories.GenericRepository[model.NewsModel]
}

func NewNewsManager(database *gorm.DB) *NewsManager {
	repo := repositories.NewGenericRepository[model.NewsModel](database)

	return &NewsManager{repo: repo}
}

func (m *NewsManager) Create(news *model.NewsModel) (*model.NewsModel, error) {
	return m.repo.Create(news)
}

func (m *NewsManager) Update(news *model.NewsModel) error {
	return m.repo.Update(news)
}

func (m *NewsManager) GetById(id int) (*model.NewsModel, error) {
	return m.repo.Get(repositories.Where("id = ?", id))
}

func (m *NewsManager) GetAll() ([]model.NewsModel, error) {
	return m.repo.Find()
}

func (m *NewsManager) PublishNews(id int) error {
	news, err := m.GetById(id)
	if err != nil {
		return err
	}

	news.Published = true
	return m.repo.Update(news)
}

func (m *NewsManager) UnpublishNews(id int) error {
	news, err := m.GetById(id)
	if err != nil {
		return err
	}

	news.Published = false
	return m.repo.Update(news)
}
