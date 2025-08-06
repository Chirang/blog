package controllers

import (
    "net/http"
    "blog/v2/models"
    "blog/v2/dto"
    "blog/v2/repository"

    "github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
    var input dto.CreatePostInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post := models.Post{Title: input.Title, Content: input.Content}
    if err := repository.CreatePost(models.DB, &post); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": post})
}

func FindPosts(c *gin.Context) {
    posts, err := repository.GetAllPosts(models.DB)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(c *gin.Context) {
    post, err := repository.GetPostByID(models.DB, c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
    post, err := repository.GetPostByID(models.DB, c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    var input dto.UpdatePostInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updated := models.Post{Title: input.Title, Content: input.Content}
    if err := repository.UpdatePost(models.DB, post, updated); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
    post, err := repository.GetPostByID(models.DB, c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    if err := repository.DeletePost(models.DB, post); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": "success"})
}
