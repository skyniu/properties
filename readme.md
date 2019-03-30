# useage

## properties 


## usage
- install

````text
go get github.com/skyniu/properties
````
- example
````go
var data = `
enable = true
host = 127.0.0.1
port = ":8082"
#log --------------------
log.level = 1
log.file = log.txt
log.files = ["info.log","error.log","debug.log",warn.log]
log.maxdays = 1
log.asyn = true
# database ---------------------------
database.user[0].username = "user123-0"
database.user[1].username = "user123-1"
database.user[0].password = "123456-0"
database.user[1].password = "123456-1"
database.address = "10.1.12.23:3306"

`
type Conf struct {
	Enable bool `json:"enable"`
	Host string `json:"host"`
	Port string `json:"port"`
	Log Log `json:"log"`
	DataBase DataBase `json:"database"`
}

type Log struct {
	Level int `json:"level"`
	File string `json:"file"`
	Files []string `json:"files"`
	Maxdays int `json:"maxdays"`
	Asyn bool `json:"asyn"`
}

type DataBase struct {
	User []User `json:"user"`
	Address string `json:"address"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}



func main() {
	p:= NewMustProperties([]byte(data))
	var c Conf
	fmt.Println(p.Unmarshal(&c))
	fmt.Printf("%+v\n",c)
}
````