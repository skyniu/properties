package properties

import (
	"strings"
	"github.com/pkg/errors"
)

func isExplain(s string)bool  {
	if strings.HasPrefix(s,"#"){
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

