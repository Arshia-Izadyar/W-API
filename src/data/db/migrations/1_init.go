package migrations

import (
	"database/sql"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/pkg/logging"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.LoadCfg())

func Up_1() {
	database := db.GetDB()
	createTables(database)
	createCountries(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	// 1
	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)

	// 2
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		panic(err)
	}

	createDefaultInfo(database)

	logger.Info(logging.Postgres, logging.Insert, "added tables", nil)
}

func addNewTable(db *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultInfo(db *gorm.DB) {
	adminRole := models.Role{Name: "admin"}
	createIfNotExist(db, &adminRole)
	defaultRole := models.Role{Name: "default"}
	createIfNotExist(db, &defaultRole)

	u := models.User{
		BaseModel:   models.BaseModel{},
		UserName:    constants.AdminRoleName,
		FirstName:   sql.NullString{Valid: true, String: "Test"},
		LastName:    sql.NullString{Valid: true, String: "last"},
		PhoneNumber: sql.NullString{Valid: true, String: "09108624707"},
		Email:       sql.NullString{Valid: true, String: "arshia@gmail.com"},
		Password:    "",
		Enable:      false,
	}
	hashed_pass, _ := bcrypt.GenerateFromPassword([]byte("a123"), bcrypt.MinCost)

	u.Password = string(hashed_pass)

	CreateAdminUser(db, &u, adminRole.Id)
}

func createIfNotExist(db *gorm.DB, r *models.Role) {
	exists := 0
	db.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		db.Create(r)
	}
}

func CreateAdminUser(db *gorm.DB, u *models.User, roleId int) {
	exists := 0
	db.
		Model(&models.User{}).
		Select("1").
		Where("user_name = ?", u.UserName).
		First(&exists)
	if exists == 0 {
		db.Create(u)
		userRole := models.UserRole{UserId: u.Id, RoleId: roleId}
		db.Create(&userRole)
	}
}

func createCountries(db *gorm.DB) {
	count := 0
	db.Model(models.Country{}).Select("count(*)").Find(&count)
	if count == 0 {
		iran := models.Country{Name: "Iran"}
		db.Create(&iran)

		usa := models.Country{Name: "USA"}
		db.Create(&usa)

		germany := models.Country{Name: "Germany"}
		db.Create(&germany)

		// Create cities and set the CountryId
		db.Create(&models.City{Name: "tehran", CountryId: iran.Id})
		db.Create(&models.City{Name: "shiraz", CountryId: iran.Id})
		db.Create(&models.City{Name: "ghazvin", CountryId: iran.Id})
		db.Create(&models.City{Name: "ahvaz", CountryId: iran.Id})
		db.Create(&models.City{Name: "kerman", CountryId: iran.Id})

		db.Create(&models.City{Name: "NY", CountryId: usa.Id})
		db.Create(&models.City{Name: "Ws", CountryId: usa.Id})
		db.Create(&models.City{Name: "Tx", CountryId: usa.Id})

		db.Create(&models.City{Name: "Berlin", CountryId: germany.Id})
		db.Create(&models.City{Name: "deF", CountryId: germany.Id})
		db.Create(&models.City{Name: "deA", CountryId: germany.Id})
	}
}

func Down_1() {

}
