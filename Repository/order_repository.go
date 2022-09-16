package repository

import (
	models "order-api/Models"

	"gorm.io/gorm"
)

type InDB struct {
	DB gorm.DB
}

type OrderRepositoryAPI interface {
	CreateOrderRepository(order models.Order) (models.Order, error)
	GetOrderRepository() ([]models.Order, error)
	UpdateOrderRepository(order models.Order) (models.Order, error)
	DeleteOrderRepository(order models.Order) (models.Order, error)
	FindOrderById(id int) (models.Order, error)
}

func (idb *InDB) CreateOrderRepository(order models.Order) (models.Order, error) {
	err := idb.DB.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func ProvideRepositoryOrder(DB gorm.DB) *InDB {
	return &InDB{DB: DB}
}

func (idb *InDB) GetOrderRepository() ([]models.Order, error) {
	var order []models.Order
	err := idb.DB.Preload("Items").Find(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (idb *InDB) UpdateOrderRepository(order models.Order) (models.Order, error) {
	err := idb.DB.Updates(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (idb *InDB) FindOrderById(id int) (models.Order, error) {
	var order models.Order
	err := idb.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (idb *InDB) DeleteOrderRepository(order models.Order) (models.Order, error) {
	err := idb.DB.Delete(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil

}
