# go-auth-api-sample

# Live Reload using Air
## Installation - Ubuntu

1. 
go get -u github.com/cosmtrek/air
or
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

2. Update your alias to alias air='$(go env GOPATH)/bin/air'
3. Execute 'air init' to create the .air.toml file

# Database
Uses GORM as ORM
Refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details on the conn string setting
example: dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"