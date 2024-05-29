package config

type User struct {
	ID       int    `json:"ID,-"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int8   `json:"role"`
	Name     string `json:"name"`
}

type Cartridge struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Model string `json:"Model"`
}

type Computer struct {
	ID             int    `json:"ID"`
	Name           string `json:"Name"`
	Processor      string `json:"Processor"`
	HDDVolume      int    `json:"HDDVolume"`
	SSDName        string `json:"SSDName"`
	SSDVolume      int    `json:"SSDVolume"`
	SSDFromFactory bool   `json:"SSDFromFactory"`
	RAMVolume      int    `json:"RAMVolume"`
	RAMFromFactory bool   `json:"RAMFromFactory"`
	MonitorID      int    `json:"MonitorName"`
	KeyboardID     int    `json:"KeyboardName"`
	MouseID        int    `json:"MouseName"`
	ContractID     int    `json:"ContractName"`
	RespPersonID   int    `json:"RespPersonID"`
}

type Contract struct {
	ID            int    `json:"ID"`
	Name          string `json:"Name"`
	Firm          string `json:"Firm"`
	Date          string `json:"Date"`
	ServerAddress string `json:"ServerAddress"`
}

type Keyboard struct {
	ID           int    `json:"ID"`
	Name         string `json:"Name"`
	Model        string `json:"Model"`
	InvNumber    string `json:"InvNumber"`
	ContractID   int    `json:"ContractID"`
	RespPersonID int    `json:"RespPersonID"`
}

type Monitor struct {
	ID           int    `json:"ID"`
	Name         string `json:"Name"`
	Model        string `json:"Model"`
	InvNumber    string `json:"InvNumber"`
	SerNumber    string `json:"SerNumber"`
	Size         int    `json:"Size"`
	ContractID   int    `json:"ContractID"`
	RespPersonID int    `json:"RespPersonID"`
}

type Mouse struct {
	ID           int    `json:"ID"`
	Name         string `json:"Name"`
	Model        string `json:"Model"`
	InvNumber    string `json:"InvNumber"`
	ContractID   int    `json:"ContractID"`
	RespPersonID int    `json:"RespPersonID"`
}

type Movement struct {
	ID         int    `json:"ID"`
	ProductID  int    `json:"ProductID"`
	Quantity   int    `json:"Quantity"`
	ForWhoID   int    `json:"ForWhoID"`
	FromWhomID int    `json:"FromWhomID"`
	Date       string `json:"Date"`
}

type Printer struct {
	ID           int    `json:"ID"`
	Name         string `json:"Name"`
	Model        string `json:"Model"`
	InvNumber    string `json:"InvNumber"`
	SerNumber    string `json:"SerNumber"`
	CartridgeID  int    `json:"CartridgeID"`
	ContractID   int    `json:"ContractID"`
	RespPersonID int    `json:"RespPersonID"`
}

type Product struct {
	ID            int    `json:"ID"`
	Name          string `json:"Name"`
	Model         string `json:"Model"`
	ProductCardID string `json:"ProductCardID"`
	RespPersonID  int    `json:"RespPersonID"`
	ContractID    int    `json:"ContractID"`
}

type ProductCard struct {
	ID         int `json:"ID"`
	ComputerID int `json:"ComputerID"`
	PrinterID  int `json:"PrinterID"`
}

type Repair struct {
	ID            int    `json:"ID"`
	PrinterID     int    `json:"PrinterID"`
	ComputerID    int    `json:"ComputerID"`
	CartridgeID   int    `json:"CartridgeID"`
	MonitorID     int    `json:"MonitorID"`
	Reason        int    `json:"Reason"`
	TransferDate  string `json:"TransferDate"`
	ReceivingDate string `json:"ReceivingDate"`
}
type RespPerson struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	JobTitle   string `json:"JobTitle"`
	Department string `json:"Department"`
}

type DBConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}
