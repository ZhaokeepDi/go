package login
import(
	"useSql"
	// "encodeJson"
	// "route"
	_"mysql"
	_"encoding/json"
	// "net/http"
	"fmt"
)
func Updata(){
	stmt:=useSql.Insert("INSERT INTO userinfo VALUES ('小小小',123456)")
	fmt.Println(stmt)
} 