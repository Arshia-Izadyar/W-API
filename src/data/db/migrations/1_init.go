package migrations

import (
	"database/sql"
	"time"
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
	createPropertyCategory(database)
	createCarType(database)
	createGearbox(database)
	createColor(database)
	createYear(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	tables = addNewTable(database, models.Country{}, tables)
	tables = addNewTable(database, models.City{}, tables)
	tables = addNewTable(database, models.File{}, tables)
	tables = addNewTable(database, models.PersianYear{}, tables)
	// Property
	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	// Car
	tables = addNewTable(database, models.Company{}, tables)
	tables = addNewTable(database, models.Gearbox{}, tables)
	tables = addNewTable(database, models.Color{}, tables)
	tables = addNewTable(database, models.CarType{}, tables)

	tables = addNewTable(database, models.CarModel{}, tables)
	tables = addNewTable(database, models.CarModelColor{}, tables)
	tables = addNewTable(database, models.CarModelYear{}, tables)
	tables = addNewTable(database, models.CarModelFile{}, tables)
	tables = addNewTable(database, models.CarModelPrice{}, tables)
	tables = addNewTable(database, models.CarModelProperty{}, tables)
	tables = addNewTable(database, models.CarModelComment{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(err, logging.Postgres, logging.Insert, "cant add tables", nil)
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

func createPropertyCategory(database *gorm.DB) {
	count := 0

	database.
		Model(&models.PropertyCategory{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})                     // بدنه
		database.Create(&models.PropertyCategory{Name: "Engine"})                   // موتور
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})               // پیشرانه
		database.Create(&models.PropertyCategory{Name: "Suspension"})               // تعلیق
		database.Create(&models.PropertyCategory{Name: "Equipment"})                // تجهیزات
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})   // سیستم های پشتیبانی راننده
		database.Create(&models.PropertyCategory{Name: "Lights"})                   // چراغ ها
		database.Create(&models.PropertyCategory{Name: "Multimedia"})               // چند رسانه ای
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})         // تجهیزات ایمنی
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"}) // صندلی و فرمان
		database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})      // پنجره و آینه
	}
	createProperty(database, "Body")
	createProperty(database, "Engine")
	createProperty(database, "Drivetrain")
	createProperty(database, "Suspension")
	createProperty(database, "Comfort")
	createProperty(database, "Driver support systems")
	createProperty(database, "Lights")
	createProperty(database, "Multimedia")
	createProperty(database, "Safety equipment")
	createProperty(database, "Seats and steering wheel")
	createProperty(database, "Windows and mirrors")

}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.
		Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select("count(*)").
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}
	var props *[]models.Property
	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)

	case "Engine":
		props = getEngineProperties(catModel.Id)

	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)

	case "Suspension":
		props = getSuspensionProperties(catModel.Id)

	case "Comfort":
		props = getComfortProperties(catModel.Id)

	case "Driver support systems":
		props = getDriverSupportSystemProperties(catModel.Id)

	case "Lights":
		props = getLightsProperties(catModel.Id)

	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)

	case "Safety equipment":
		props = getSafetyEquipmentProperties(catModel.Id)

	case "Seats and steering wheel":
		props = getSeatsProperties(catModel.Id)

	case "Windows and mirrors":
		props = getWindowsProperties(catModel.Id)

	default:
		props = &([]models.Property{})
	}

	for _, prop := range *props {
		database.Create(&prop)
	}
}
func createCarType(database *gorm.DB) {
	count := 0
	database.
		Model(&models.CarType{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.CarType{Name: "Crossover"})
		database.Create(&models.CarType{Name: "Sedan"})
		database.Create(&models.CarType{Name: "Sports"})
		database.Create(&models.CarType{Name: "Coupe"})
		database.Create(&models.CarType{Name: "Hatchback"})
	}
}

func createGearbox(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Gearbox{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Gearbox{Name: "Manual"})
		database.Create(&models.Gearbox{Name: "Automatic"})
	}
}

func createColor(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Color{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Color{Name: "Black", Hex: "#000000"})
		database.Create(&models.Color{Name: "White", Hex: "#ffffff"})
		database.Create(&models.Color{Name: "Blue", Hex: "#0000ff"})
	}
}

func createYear(database *gorm.DB) {
	count := 0
	database.
		Model(&models.PersianYear{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {

		database.Create(&models.PersianYear{
			PersianTitle: "1402",
			Year:         1402,
			StartAt:      time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2024, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1401",
			Year:         1401,
			StartAt:      time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1400",
			Year:         1400,
			StartAt:      time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1399",
			Year:         1399,
			StartAt:      time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1398",
			Year:         1398,
			StartAt:      time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1398",
			Year:         1398,
			StartAt:      time.Date(2018, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})
	}
}

func Down_1() {

}
