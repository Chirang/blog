package dto

type CreatePostInput struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

type UpdatePostInput struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}
