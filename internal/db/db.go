package db

import(
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true", 
		user, 
		pass, 
		host, 
		port, 
		name
	)
	db, err := gorm.Open(
		mysql.Open(dsn), 
		&gorm.Config{}
	)
	if err != nil {
		panic("Error connecting database")
	}
	return db
}