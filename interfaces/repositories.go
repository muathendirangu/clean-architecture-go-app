package interfaces

type DBHandler interface {
	Execute(query string)
	Query(query string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}
