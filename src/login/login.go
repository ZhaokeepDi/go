package login
import(
	"useSql"
	"encodeJson"
	// "route"
	_"mysql"
	_"encoding/json"
	"net/http"
	"fmt"
)
func Login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	 rows:=useSql.Select("SELECT username,password FROM userinfo WHERE password='"+r.Form["password"][0]+"' AND username='"+r.Form["username"][0]+"'")
	 if rows.Next(){
		 reqD:=map[string]interface{}{"code":1,"msg":"登录成功"}
		 cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", MaxAge: 86400}
    	 http.SetCookie(w, &cookie)
		 fmt.Fprintf(w,encodeJson.Encode(reqD))
	 }else{
		reqD:=map[string]interface{}{"code":0,"msg":"登录失败"}
		fmt.Fprintf(w,encodeJson.Encode(reqD))
	 } 
	 for rows.Next() {
        var username string
        var password string
		err := rows.Scan(&username, &password)
		if err!=nil{
			panic(err)
		}
        fmt.Println(username)
        fmt.Println(password)
    }
	if(r.Method=="POST"){
		fmt.Println("POST请求") 
	}else{
		fmt.Println("GET请求")
	}
}
func ReadCookie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, err := r.Cookie("testcookiename")
    if err == nil {
        cookievalue := cookie.Value
        w.Write([]byte("<b>cookie的值是：" + cookievalue + "</b>\n"))
    } else {
        w.Write([]byte("<b>读取出现错误：" + err.Error() + "</b>\n"))
    }
}