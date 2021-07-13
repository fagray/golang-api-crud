module babelist.ph

replace babelist.ph/api/models/babe => ./api/models

replace babelist.ph/api/config/database => ./api/config

go 1.14

require (
	babelist.ph/api/config/database v0.0.0-00010101000000-000000000000
	babelist.ph/api/models/babe v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16 // indirect
	gorm.io/driver/mysql v1.0.4 // indirect
	gorm.io/driver/sqlite v1.1.4 // indirect
	gorm.io/gorm v1.20.12 // indirect
)
