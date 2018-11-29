package useSql
import(
	"database/sql"
	_"mysql"
	"outErr"
) 
func getSql() *sql.DB{
	db,err:=sql.Open("mysql","root:000000@/test?charset=utf8")
	outErr.OutErr(err);
	return db
}
func Select(SelectSql string) *sql.Rows{
	db:=getSql()
	rows,err:=db.Query(SelectSql)
	outErr.OutErr(err);
	return rows
}
func Insert(InsertSql string) *sql.Stmt{
	db:=getSql()
	stmt,err:=db.Prepare(InsertSql)
	outErr.OutErr(err);
	return stmt
}