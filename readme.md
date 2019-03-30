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
log.level = 1
log.file = log.txt
log.files = ["info.log","error.log","debug.log"]
log.maxdays = 1
log.asyn = true

database.user.username = "user123"
database.user.password = "123456"
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
	User User `json:"user"`
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