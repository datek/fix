package examples

type UserStore interface {
	CreateUser(name string) error
	DeleteUser(name string) error
}

func NewUserStore(db DB) UserStore {
	return &userStore{db}
}

type userStore struct {
	db DB
}

func (store *userStore) CreateUser(name string) error {
	return store.db.Execute("INSERT " + name)
}

func (store *userStore) DeleteUser(name string) error {
	return store.db.Execute("DELETE " + name)
}

type DB interface {
	Execute(statement string) error
}

func NewMockDB(execute func(statement string) error) DB {
	return &MockDb{ExecuteMock: execute}
}

type MockDb struct {
	ExecuteMock func(statement string) error
}

func (db *MockDb) Execute(statement string) error {
	return db.ExecuteMock(statement)
}
