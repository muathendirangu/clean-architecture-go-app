package interfaces

type DBHandler interface {
	Execute(query string)
	Query(query string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DBRepository struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}

type DBUserRepository DBRepository
type DBOrderRepository DBRepository
type DBCustomerRepository DBRepository
type DBItemRepository DBRepository

func NewDBUserRepository(dbHandlers map[string]DBHandler) *DBUserRepository {
	dbUserRepository := new(DBUserRepository)
	dbUserRepository.dbHandlers = dbHandlers
	dbUserRepository.dbHandler = dbHandlers["DBUserRepository"]
	return dbUserRepository
}
