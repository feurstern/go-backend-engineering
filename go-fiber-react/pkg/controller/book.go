package controller

import (
	"fmt"
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	db := database.DBConnection

	title := c.FormValue("title")
	author := c.FormValue("author")
	description := c.FormValue("description")
	categoryIdStr := c.FormValue("category_id")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Inavlid category id",
		})
	}

	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success ": false,
			"message":  "Invalid request!",
		})
	}

	book := model.Book{
		Title:       title,
		Author:      author,
		Description: description,
		Category_id: categoryId,
	}

	if err := db.Create(&book).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert the data!",
		})
	}

	files := form.File["images"]

	var images []model.BookCover

	for _, x := range files {
		filename := fmt.Sprintf("asset/images/%s", x.Filename)

		if err := c.SaveFile(x, filename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Internal server error",
			})
		}

		images = append(images, model.BookCover{
			BookId: book.ID,
			Name:   x.Filename,
			Path:   filename,
		})

	}

	if len(images) > 0 {

		if err := db.Create(&images).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    book,
	})

}

func Create(c *fiber.Ctx) error {
	db := database.DBConnection

	title := c.FormValue("title")
	author := c.FormValue("author")
	description := c.FormValue("description")
	categoryIdStr := c.FormValue("category_id")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Inval.id body request",
		})
	}

	book := model.Book{
		Title:       title,
		Category_id: categoryId,
		Description: description,
		Author:      author,
	}

	if err := db.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Inval.id body request",
		})
	}

	files := form.File["images"]
	var bookCover []model.BookCover

	for _, x := range files {
		filename := fmt.Sprintf("/asset/image/%s", x.Filename)
		if err := c.SaveFile(x, filename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to store the data",
			})
		}

		bookCover = append(bookCover, model.BookCover{
			BookId: book.ID,
			Name:   x.Filename,
			Path:   filename,
		})
	}

	if len(bookCover) > 0 {
		if err := db.Create(&bookCover).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to inser the book cover!",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    book,
	})

}
