package mockjson

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserJson() {
	gofakeit.Seed(0)

	// Slice เก็บ user
	var users []User

	for i := 1; i <= 100; i++ {
		now := time.Now()
		user := User{
			ID:        i,
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Summary:   gofakeit.LoremIpsumWord(),
			CreatedAt: now,
			UpdatedAt: now,
		}
		users = append(users, user)
	}

	// สร้างไฟล์ users.json
	file, err := os.Create("./mock_json/users.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// เขียน JSON ลงไฟล์
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // ทำให้สวยงาม
	if err := encoder.Encode(users); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Successfully created users.json with 100 users")
}
