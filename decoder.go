package properties

import (
	"bufio"
	"bytes"
	"strings"
	"github.com/oliveagle/jsonpath"
)

type Decoder struct {
	propMap map[string]interface{}
}

func (d *Decoder)parse(data []byte) error {
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
		err = marshalInterface(kvs[0],d.propMap,kvs[1])
		if err !=nil{
			return err
		}
	}
	return nil
}

func (d *Decoder)Map()map[string]interface{}  {
	return d.propMap
}

func (d *Decoder)Set(k string,v interface{})  {
	marshalInterface(k,d.propMap,v)
}

func (d *Decoder)Get(k string)  interface{}{
	v,_:=jsonpath.JsonPathLookup(d.propMap,"$."+k)
	return v
}

func (d *Decoder)Unmarshal(v interface{})error  {

	return nil
}
func NewDecoder(data []byte)(*Decoder,error)  {
	d:=&Decoder{propMap:map[string]interface{}{}}
	return d,d.parse(data)
}

func NewMustDecoder(data []byte)*Decoder  {
	d:=&Decoder{propMap:map[string]interface{}{}}
	err:=d.parse(data)
	if err !=nil{
		panic(err)
	}
	return d
}

