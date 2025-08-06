package repository

import (
    "blog/v2/models"
    "gorm.io/gorm"
)

func GetAllPosts(db *gorm.DB) ([]models.Post, error) {
    var posts []models.Post
    err := db.Find(&posts).Error
    return posts, err
}

func GetPostByID(db *gorm.DB, id string) (*models.Post, error) {
    var post models.Post
    err := db.Where("id = ?", id).First(&post).Error
    return &post, err
}

func CreatePost(db *gorm.DB, post *models.Post) error {
    return db.Create(post).Error
}

func UpdatePost(db *gorm.DB, post *models.Post, updated models.Post) error {
    return db.Model(post).Updates(updated).Error
}

func DeletePost(db *gorm.DB, post *models.Post) error {
    return db.Delete(post).Error
}
