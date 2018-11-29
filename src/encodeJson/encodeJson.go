package encodeJson
import(
	"encoding/json"
	"outErr"
)
func Encode(data interface{}) string{
	obj,err:=json.Marshal(data)
	outErr.OutErr(err);
	return string(obj)
}