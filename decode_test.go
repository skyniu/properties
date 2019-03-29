package properties

import (
	"testing"
	"fmt"
)
var data = `
log.level=1
log.file=log.txt
`

func TestMarshal(t *testing.T) {
	d,_:=NewDecoder([]byte(data))
	fmt.Println(d.Map())
}
