package controllers

import (
	"BookStore/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) BookController {
	return BookController{DB: db}
}

func (ac *BookController) GetBooks(ctx *gin.Context) {
	// ordered by cost in descending, ascending orders by param price_order
	var books []models.Book
	order := ctx.Query("price")
	var query strings.Builder
	query.WriteString("SELECT * FROM books")
	title := ctx.Query("title")
	if title != "" {
		query.WriteString(" WHERE title LIKE '%")
		query.WriteString(title)
		query.WriteString("%'")
	}
	if order == "asc" {
		query.WriteString(" ORDER BY price ASC")
	} else if order == "desc" {
		query.WriteString(" ORDER BY price DESC")
	}
	// search by title
	ac.DB.Raw(query.String()).Scan(&books)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success", "data": books,
	})
}

func (ac *BookController) GetBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")
	result := ac.DB.Raw("SELECT * FROM books WHERE id = ?", id).Scan(&book)
	// check if it is empty
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "fail", "message": "Book not found with id: " + id,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success", "data": book,
	})
}

func (ac *BookController) CreateBook(ctx *gin.Context) {
	var payload *models.Book
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", "message": err.Error(),
		})
		return
	}
	newBook := models.Book{
		Title:       payload.Title,
		Description: payload.Description,
		Price:       payload.Price,
	}
	result := ac.DB.Raw(
		"INSERT INTO books (title, description, price) VALUES (?, ?, ?) RETURNING *",
		newBook.Title, newBook.Description, newBook.Price,
	).Scan(&newBook)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "error", "message": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success", "data": newBook,
	})
}
func (ac *BookController) UpdateBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")
	result := ac.DB.Raw("SELECT * FROM books WHERE id = ?", id).Scan(&book)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "fail", "message": "Book not found",
		})
		return
	}
	var payload *models.Book
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", "message": err.Error(),
		})
		return
	}

	var updates []string
	var params []interface{}
	if payload.Title != "" {
		updates = append(updates, "title = ?")
		params = append(params, payload.Title)
	}
	if payload.Description != "" {
		updates = append(updates, "description = ?")
		params = append(params, payload.Description)
	}
	if payload.Price != 0 {
		updates = append(updates, "price = ?")
		params = append(params, payload.Price)
	}

	if len(updates) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", "message": "No fields to update",
		})
		return
	}

	query := "UPDATE books SET " + strings.Join(updates, ", ") + " WHERE id = ? RETURNING *"
	params = append(params, id)
	result = ac.DB.Raw(query, params...).Scan(&book)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "error", "message": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success", "books": book,
	})
}

func (ac *BookController) DeleteBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")
	result := ac.DB.Raw("SELECT * FROM books WHERE id = ?", id).Scan(&book)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "fail", "message": "Book not found",
		})
		return
	}
	result = ac.DB.Delete(&book, id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "error", "message": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success", "message": "Book deleted successfully",
	})
}
