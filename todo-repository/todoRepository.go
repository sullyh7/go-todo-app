package TodoRepository

import (
	"fmt"
	"sully/todo-app/model"

	"gorm.io/gorm"
)

func FindAll(db *gorm.DB) ([]model.TodoItem, error) {
	var items []model.TodoItem
	if err := db.Find(&items).Error; err != nil {
		return []model.TodoItem{}, err
	}
	return items, nil
}

func FindById(db *gorm.DB, id int) (model.TodoItem, error) {
	var item model.TodoItem
	if err := db.Where("id = ?", id).First(&item).Error; err != nil {
		return model.TodoItem{}, err
	}
	return item, nil
}

func EditById(db *gorm.DB, id int, updatedItem *model.TodoItem) error {
	if err := db.Model(&model.TodoItem{}).Where("id = ?", id).Updates(updatedItem).Error; err != nil {
		return err
	}
	return nil
}

func Save(db *gorm.DB, item *model.TodoItem) error {
	if err := db.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func Complete(db *gorm.DB, id int) error {
	var updatedItem model.TodoItem
	if err := db.Where("id = ?", id).First(&updatedItem).Error; err != nil {
		return err
	}
	if updatedItem.Done {
		return fmt.Errorf("item already complete")
	}
	updatedItem.Done = true
	if err := db.Model(&model.TodoItem{}).Where("id = ?", id).Updates(updatedItem).Error; err != nil {
		return err
	}
	return nil
}

func DeleteById(db *gorm.DB, id string) error {
	if err := db.Where("id = ?", id).Delete(&model.TodoItem{}).Error; err != nil {
		return err
	}
	return nil
}
