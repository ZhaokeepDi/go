package route
import(
	"net/http"
)
func Get(r *http.Request){
	if(r.Method!="GET"){panic("请求方式错误，应该为GET")}
	
}
func Post(r *http.Request){
	if(r.Method!="POST"){panic("请求方式错误，应该为POST")}
}