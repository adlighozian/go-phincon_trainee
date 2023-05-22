package repository

import (
	"contact-go/model"
	"log"

	"gorm.io/gorm"
)

type gormRepository struct {
	Conn *gorm.DB
}

func NewContactRepositoryGorm(connection *gorm.DB) ContactRepository {
	return &gormRepository{
		Conn: connection,
	}
}

func (repo gormRepository) List() ([]model.Client, error) {
	log.Println("list repository gorm")

	var result []model.Client
	query := repo.Conn.Model(&model.Client{}).Find(&result)

	return result, query.Error
}

func (repo *gormRepository) Add(req []model.ContactRequest) ([]model.Client, error) {
	log.Println("add repository gorm")

	var result []model.Client

	for _, data := range req {
		result = append(result, model.Client{
			Name:   data.Name,
			NoTelp: data.NoTelp,
		})
	}

	query := repo.Conn.Model(&model.Client{}).Create(&result)

	return result, query.Error
}
func (repo *gormRepository) Update(id int, req model.ContactRequest) error {
	log.Println("update repository gorm")

	query := repo.Conn.Model(&model.Client{}).Where("id = ?", id).Updates(&req)

	return query.Error
}
func (repo *gormRepository) Delete(id int) error {
	log.Println("delete repository gorm")

	query := repo.Conn.Delete(&model.Client{}, id)

	return query.Error
}
