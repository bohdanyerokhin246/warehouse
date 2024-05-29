package init

import "site/internal/config"

var DBconn = config.DBConnection{Host: "localhost", Port: "5432", User: "postgres", Password: "Admin@@246", DbName: "site", SSLMode: "disable"}
var Cartridge = config.Cartridge{}
var Computer = config.Computer{}
