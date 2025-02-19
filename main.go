package main

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Activity struct {
	ID           int       `json:"id"`
	Title        string    `json:"title" validate:"required"`
	Category     string    `json:"category" validate:"required"`
	ActivityDate time.Time `json:"activity_date"`
	Status       string    `json:"status" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description" validate:"required"`
}

func initDB() (*sql.DB, error) {
	dns := "postgres://postgres.paexclsqvwfcswxyqsoz:2QVULt6xiEVLBbfW@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	app := fiber.New()
	validate := validator.New()

	// GET /activities
	app.Get("/activities", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM activities")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		var activities []Activity
		for rows.Next() {
			var activity Activity
			err := rows.Scan(
				&activity.ID,
				&activity.Title,
				&activity.Category,
				&activity.ActivityDate,
				&activity.Status,
				&activity.CreatedAt,
				&activity.Description)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			activities = append(activities, activity)
		}

		return c.Status(fiber.StatusOK).JSON(activities)
	})

	// POST /activities
	app.Post("/activities", func(c *fiber.Ctx) error {
		var activity Activity
		err := c.BodyParser(&activity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		err = validate.Struct(&activity)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Set activity_date to current time if it's zero
		if activity.ActivityDate.IsZero() {
			activity.ActivityDate = time.Now()
		}

		sqlStatement := `INSERT INTO activities (title, category, activity_date, status, description) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, category, activity_date, status, created_at, description`
		err = db.QueryRow(sqlStatement, activity.Title, activity.Category, activity.ActivityDate, activity.Status, activity.Description).Scan(&activity.ID, &activity.Title, &activity.Category, &activity.ActivityDate, &activity.Status, &activity.CreatedAt, &activity.Description)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status": "success",
			"data":   activity,
		})
	})

	// PUT /activities/:id
	app.Put("/activities/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		var activity Activity
		err = c.BodyParser(&activity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		err = validate.Struct(&activity)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		sqlStatement := `UPDATE activities SET title = $1, category = $2, activity_date = $3, status = $4, description = $5 WHERE id = $6 RETURNING id, title, category, activity_date, status, created_at, description`
		err = db.QueryRow(sqlStatement, activity.Title, activity.Category, activity.ActivityDate, activity.Status, activity.Description, id).Scan(&activity.ID, &activity.Title, &activity.Category, &activity.ActivityDate, &activity.Status, &activity.CreatedAt, &activity.Description)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"data":   activity,
		})
	})

	// DELETE /activities/:id
	app.Delete("/activities/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		sqlStatement := `DELETE FROM activities WHERE id = $1`
		_, err = db.Exec(sqlStatement, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Activity deleted successfully",
		})
	})

	app.Listen(":8000")
}
