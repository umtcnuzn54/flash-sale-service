package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause" 
)

type Product struct {
	gorm.Model
	Name  string
	Stock int
	Price float64
}

func main() {
	dsn := "host=localhost user=admin password=password dbname=flashsaledb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı hatası:", err)
	}

	db.AutoMigrate(&Product{})

	var count int64
	db.Model(&Product{}).Count(&count)
	if count == 0 {
		db.Create(&Product{Name: "iPhone 15", Stock: 100, Price: 50000})
	} else {
		db.Model(&Product{}).Where("id = ?", 1).Update("stock", 100)
	}
	fmt.Println("📦 Stok 100 olarak ayarlandı, saldırıya hazır!")

	app := fiber.New()

	app.Get("/buy", func(c *fiber.Ctx) error {
		
		err := db.Transaction(func(tx *gorm.DB) error {
			var product Product

			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, 1).Error; err != nil {
				return err
			}

			if product.Stock <= 0 {
				return fmt.Errorf("Stok bitti")
			}

			time.Sleep(10 * time.Millisecond)

			product.Stock = product.Stock - 1
			return tx.Save(&product).Error
		})

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Stok tükendi!"})
		}

		return c.JSON(fiber.Map{"message": "Başarılı!"})
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		var product Product
		db.First(&product, 1)
		return c.JSON(product)
	})

	app.Listen(":3000")
}
