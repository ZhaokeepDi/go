package outErr
func OutErr(err error){
	if err!=nil{
		panic(err)
	}
}