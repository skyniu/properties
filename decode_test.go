package properties

import (
	"testing"
	"fmt"
	"regexp"
)
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



func TestUnMarshal(t *testing.T) {
	p:= NewMustProperties([]byte(data))
	var c Conf
	fmt.Println(p.Unmarshal(&c))
	fmt.Printf("%+v\n",c)

	p.Set("name",1)
	p.Set("kong","haha")

}

func TestMarshal(t *testing.T) {
	r:=regexp.MustCompile("((.+),|(\".+\"),)*")
	fmt.Println(r.MatchString("1234,5678,\"ddfg\"",))
	sub:=r.FindAllSubmatch([]byte("1234,5678,\"ddfg\""),-1)
	for _,v:=range sub{
		fmt.Println(len(v))
		for i:=0;i< len(v);i++{
			fmt.Println(string(v[i]))
		}
	}
}

func TestMarshalInterface(t *testing.T) {
	splitArray()
}