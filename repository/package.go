package repository

import (
	"restaurant/models"

	"gorm.io/gorm"
)

type PackageRepository interface {
	CreatePackage() error
	GetAllPackages() ([]models.Package, error)
}

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) PackageRepository {
	return &packageRepository{db}
}

func (repo *packageRepository) CreatePackage() error {
	return nil
}

func (repo *packageRepository) GetAllPackages() ([]models.Package, error) {
	return nil, nil
}
