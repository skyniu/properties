package properties

import (
	"bufio"
	"bytes"
	"strings"
	"github.com/oliveagle/jsonpath"
	"encoding/json"
	"io/ioutil"
)

type Properties struct {
	propMap map[string]interface{}
}

func (p *Properties)parse(data []byte) error {
	sc:=bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan(){
		text:=sc.Text()
		text = strings.Trim(text," ")
		//is explain ,ignore
		if isExplain(text){
			continue
		}
		kvs,err:=split(text)
		if err !=nil{
			continue
		}
		err = marshalInterface(kvs[0], p.propMap,resolveValue(kvs[1]))
		if err !=nil{
			return err
		}
	}
	return nil
}



func (p *Properties)Map()map[string]interface{}  {
	return p.propMap
}

func (p *Properties)Set(k string,v interface{})  {
	marshalInterface(k, p.propMap,v)
}

func (p *Properties)Get(k string)  interface{}{
	v,_:=jsonpath.JsonPathLookup(p.propMap,"$."+k)
	return v
}

func (p *Properties)GetString(k string)string  {
	if s,ok:=p.Get(k).(string);ok{
		return s
	}
	return ""
}

func (p *Properties)GetInt(k string,def int)int  {
	if s,ok:=p.Get(k).(float64);ok{
		return int(s)
	}
	return def
}

func (p *Properties)GetBool(k string,def bool)bool  {
	if s,ok:=p.Get(k).(bool);ok{
		return s
	}
	return def
}

func (p *Properties)GetFloat(k string,def float64)float64  {
	if s,ok:=p.Get(k).(float64);ok{
		return s
	}
	return def
}

func (p *Properties)Unmarshal(v interface{})error  {
	b,_:=json.Marshal(p.propMap)
	return json.Unmarshal(b,v)
}

func (p *Properties)Update(v interface{})error  {
	b,_:=json.Marshal(v)
	return json.Unmarshal(b, p.propMap)
}

func NewProperties(data []byte)(*Properties,error)  {
	d:=&Properties{propMap:map[string]interface{}{}}
	return d,d.parse(data)
}

func NewPropertiesFromFile(f string)(*Properties,error)  {
	b,err:= ioutil.ReadFile(f)
	if err !=nil{
		return nil,err
	}
	return NewProperties(b)
}

func NewMustProperties(data []byte)*Properties {
	d:=&Properties{propMap:map[string]interface{}{}}
	err:=d.parse(data)
	if err !=nil{
		panic(err)
	}
	return d
}

