package main
import(
	"fmt"
	"net/http"
	_"reflect"
	// "database/sql"
	"login"
	"os"
	"io"
)
func main(){
	//注册处理器，             stripPrefix将URL.path的带有/template/的字段给去掉，交给第二个参数处理
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
	route()
	err :=http.ListenAndServe(":999",nil)
	if err!=nil{
		panic("服务建立错误")
	}
}
func route(){
	http.HandleFunc("/login",login.Login)
	http.HandleFunc("/regist",login.Insert)
	http.HandleFunc("/readCookie",login.ReadCookie)
	http.HandleFunc("/undata",upFile)
}
func upFile(w http.ResponseWriter,r *http.Request){
		// 从请求里读取文件;相关链接
		// FormFile https://studygolang.com/static/pkgdoc/pkg/net_http.htm#Request.FormFile
		file,fileHead,_:=r.FormFile("file")
		fmt.Println(fileHead.Size)
		// 创建同类型同名字文件  相关链接
		// Create https://studygolang.com/static/pkgdoc/pkg/os.htm#Create
		f, _ := os.Create("./src/upFile/"+fileHead.Filename)
		fmt.Println(f.Name)
		// 拷贝上传文件的二进制数据到自己新建的文件里
		// https://studygolang.com/static/pkgdoc/pkg/io.htm#Copy
		io.Copy(f, file)
		defer f.Close()
		defer file.Close()
		fmt.Println("upload success")
}