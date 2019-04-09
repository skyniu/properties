package properties

import (
	"strings"
	"strconv"
	"regexp"
	"errors"
	"fmt"
)

func isExplain(s string)bool  {
	if strings.HasPrefix(s,"#") || s==""{
		return true
	}
	return false
}
// split the line and trim space
func split(s string)([]string ,error) {
	ss:=strings.SplitN(s,"=",2)
	if len(ss)!=2{
		return nil,errors.New("error line:"+s)
	}
	var rs = []string{strings.Trim(ss[0]," "),strings.Trim(ss[1]," ")}
	return rs,nil
}

func resolveValue(v string)interface{}  {
	if strings.HasPrefix(v,"\"")&& strings.HasSuffix(v,"\""){
		return strings.Trim(v,"\"")
	}else if strings.HasPrefix(v,"[") && strings.HasSuffix(v,"]") {

		vs:=strings.Split(v[1:len(v)-1],",")
		if len(vs)==1 && strings.Trim(vs[0]," ")==""{
			return []interface{}{}
		}
		var values []interface{}
		for _,vue:=range vs{
			values = append(values,resolveValue(strings.Trim(vue," ")))
		}
		return values
	} else{
		if i,err:=strconv.ParseFloat(v,64);err==nil{
			return i
		}
		if b,err:=strconv.ParseBool(v);err ==nil{
			return b
		}

	}
	return v
}

func splitArray(){
	r:=regexp.MustCompile("\\[((.+,)|(\".+,\"))*\\]")
	g:=r.FindAllSubmatch([]byte(`[1,2,3,4,5,]`),-1)
	fmt.Println(len(g))
	for _,v:=range g{
		for _,l:=range v{
			fmt.Print(string(l)," ")
		}
		fmt.Println()
	}

}

func splitA(s string)  {

}
