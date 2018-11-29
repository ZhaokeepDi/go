package login
import(
	"useSql"
	// "encodeJson"
	// "route"
	_"mysql"
	"encodeJson"
	"net/http"
	"fmt"
	"outErr"
)
func Insert(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username,password:=r.Form["username"][0],r.Form["password"][0];
	seleRes:=useSql.Select("SELECT username FROM userinfo WHERE username='"+username+"'")
	if seleRes.Next(){
		res:=map[string] interface{}{"code":"0","msg":"账号重复"}
		fmt.Fprintf(w,encodeJson.Encode(res))
		return
	}
	// 写入
	stmt:=useSql.Insert("INSERT userinfo SET username=?,password=?")
	res,err:=stmt.Exec(username,password)
	outErr.OutErr(err)
	if err!=nil{
		res:=map[string] interface{}{"code":"0","msg":"注册失败"}
		fmt.Fprintf(w,encodeJson.Encode(res))
	}else{
		res:=map[string] interface{}{"code":"1","msg":"注册成功"}
		fmt.Fprintf(w,encodeJson.Encode(res))
	}
	fmt.Println(res.LastInsertId())
}