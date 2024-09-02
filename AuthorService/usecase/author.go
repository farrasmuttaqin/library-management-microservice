package usecase

import (
	"author_service/configs"
	"author_service/helpers"
	"author_service/helpers/database"
	"author_service/models"
	repo "author_service/repositories"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type AuthorUseCase struct {
	Configs     configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	AuthorRepo  repo.AuthorRepository
}

func (u *AuthorUseCase) GetAuthorById(id uint) (map[string]interface{}, error) {
	cacheKey := "author:" + strconv.FormatUint(uint64(id), 10)

	// Try to get the data from Redis cache
	content, err := u.RedisHelper.Get(cacheKey).Result()
	if err == nil {
		var author models.Author
		if err2 := json.Unmarshal([]byte(content), &author); err2 == nil && !helpers.IsEmptyStruct(author) {
			return map[string]interface{}{
				"status":  1,
				"message": "success",
				"author":  author,
			}, nil
		}
	}

	// If not in cache, get from database
	qId := map[string]interface{}{
		"id": id,
	}
	author, err2 := u.AuthorRepo.Get(qId)
	if err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to retrieve author",
		}, err2
	}

	if helpers.IsEmptyStruct(author) {
		return map[string]interface{}{
			"status":  0,
			"message": "Author not found",
		}, errors.New("author is empty")
	}

	// Cache the data
	authorJSON, _ := json.Marshal(author)
	cacheTTL := time.Duration(u.Configs.RedisMasterConfiguration().TTLSecond) * time.Second
	if err3 := u.RedisHelper.Set(cacheKey, string(authorJSON), cacheTTL); err3 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "success get data, but failed to cache data",
			"author":  author,
		}, err3
	}

	return map[string]interface{}{
		"status":  1,
		"message": "success",
		"author":  author,
	}, nil
}

func (u *AuthorUseCase) CreateAuthor(authorData models.Author) (map[string]interface{}, error) {
	// Check if the author with the same email already exists
	qValidation := map[string]interface{}{
		"email": authorData.Email,
	}
	existingAuthorEmail, _ := u.AuthorRepo.Get(qValidation)
	if !helpers.IsEmptyStruct(existingAuthorEmail) {
		return map[string]interface{}{
			"status":  0,
			"message": "Author with this email already exists",
		}, errors.New("author with this email already exists")
	}

	// Create author in the database
	err := u.AuthorRepo.Create(&authorData)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	return map[string]interface{}{
		"status":  1,
		"message": "Author created successfully",
		"author":  authorData,
	}, nil
}

func (u *AuthorUseCase) UpdateAuthor(id uint, authorData models.Author) (map[string]interface{}, error) {
	// Check
	qValidation := map[string]interface{}{
		"id": id,
	}
	exists, _ := u.AuthorRepo.Get(qValidation)
	if helpers.IsEmptyStruct(exists) {
		return map[string]interface{}{
			"status":  0,
			"message": "Author not found",
		}, errors.New("author not found")
	}

	// Check if the author with the same emailData already exists
	qValidation = map[string]interface{}{
		"email": authorData.Email,
	}
	existingAuthorEmail, _ := u.AuthorRepo.GetExceptId(id, qValidation)
	if !helpers.IsEmptyStruct(existingAuthorEmail) {
		return map[string]interface{}{
			"status":  0,
			"message": "Author with this email already exists",
		}, errors.New("author with this email already exists")
	}

	// Update author in the database
	err := u.AuthorRepo.Update(id, &authorData)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	cacheKey := "author:" + strconv.FormatUint(uint64(id), 10)
	err = u.RedisHelper.Del(cacheKey)

	return map[string]interface{}{
		"status":  1,
		"message": "Author updated successfully",
		"author":  authorData,
	}, nil
}

func (u *AuthorUseCase) DeleteAuthor(id uint) (map[string]interface{}, error) {
	// Delete author from the database
	if err := u.AuthorRepo.Delete(id); err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to delete author",
		}, err
	}

	cacheKey := "author:" + strconv.FormatUint(uint64(id), 10)
	_ = u.RedisHelper.Del(cacheKey)

	return map[string]interface{}{
		"status":  1,
		"message": "Author deleted successfully",
	}, nil
}
