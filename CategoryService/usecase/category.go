package usecase

import (
	"category_service/configs"
	"category_service/helpers"
	"category_service/helpers/database"
	"category_service/models"
	repo "category_service/repositories"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type CategoryUseCase struct {
	Configs      configs.ConfigurationsInterface
	RedisHelper  database.RedisHelper
	CategoryRepo repo.CategoryRepository
}

func (u *CategoryUseCase) GetCategoryById(id uint) (map[string]interface{}, error) {
	cacheKey := "category:" + strconv.FormatUint(uint64(id), 10)

	// Try to get the data from Redis cache
	content, err := u.RedisHelper.Get(cacheKey).Result()
	if err == nil {
		var category models.Category
		if err2 := json.Unmarshal([]byte(content), &category); err2 == nil && !helpers.IsEmptyStruct(category) {
			return map[string]interface{}{
				"status":   1,
				"message":  "success",
				"category": category,
			}, nil
		}
	}

	// If not in cache, get from database
	qId := map[string]interface{}{
		"id": id,
	}
	category, err2 := u.CategoryRepo.Get(qId)
	if err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to retrieve category",
		}, err2
	}

	if helpers.IsEmptyStruct(category) {
		return map[string]interface{}{
			"status":  0,
			"message": "Category not found",
		}, errors.New("category is empty")
	}

	// Cache the data
	categoryJson, _ := json.Marshal(category)
	cacheTTL := time.Duration(u.Configs.RedisMasterConfiguration().TTLSecond) * time.Second
	if err3 := u.RedisHelper.Set(cacheKey, string(categoryJson), cacheTTL); err3 != nil {
		return map[string]interface{}{
			"status":   0,
			"message":  "success get data, but failed to cache data",
			"category": category,
		}, err3
	}

	return map[string]interface{}{
		"status":   1,
		"message":  "success",
		"category": category,
	}, nil
}

func (u *CategoryUseCase) CreateCategory(categoryData models.Category) (map[string]interface{}, error) {
	// Create category in the database
	err := u.CategoryRepo.Create(&categoryData)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	return map[string]interface{}{
		"status":  1,
		"message": "Category created successfully",
		"book":    categoryData,
	}, nil
}

func (u *CategoryUseCase) UpdateCategory(id uint, categoryData models.Category) (map[string]interface{}, error) {
	// Check
	qValidation := map[string]interface{}{
		"id": id,
	}
	exists, _ := u.CategoryRepo.Get(qValidation)
	if helpers.IsEmptyStruct(exists) {
		return map[string]interface{}{
			"status":  0,
			"message": "Category not found",
		}, errors.New("category not found")
	}

	// Update book in the database
	err := u.CategoryRepo.Update(id, &categoryData)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to update category",
		}, err
	}

	cacheKey := "category:" + strconv.FormatUint(uint64(id), 10)
	_ = u.RedisHelper.Del(cacheKey)

	return map[string]interface{}{
		"status":  1,
		"message": "Category updated successfully",
		"author":  categoryData,
	}, nil
}

func (u *CategoryUseCase) DeleteCategory(id uint) (map[string]interface{}, error) {
	// Delete category from the database
	if err := u.CategoryRepo.Delete(id); err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to delete category",
		}, err
	}

	cacheKey := "category:" + strconv.FormatUint(uint64(id), 10)
	_ = u.RedisHelper.Del(cacheKey)

	return map[string]interface{}{
		"status":  1,
		"message": "Category deleted successfully",
	}, nil
}
