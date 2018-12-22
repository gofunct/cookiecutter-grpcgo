package store

import (
"errors"
	validator "gopkg.in/go-playground/validator.v9"
	"os"
"time"
"golang.org/x/crypto/bcrypt"
)

var (
	valid = validator.New()
	ErrAccountNotFound = errors.New("account not found")
	ErrEmailExists     = errors.New("email already exists")
	ErrNoDatabase      = errors.New("no database connection details")
	ErrNoPasswordGiven = errors.New("a password is required")
)

type Database interface {
	List(count int32, token string) ([]*Account, string, error)
	ReadByID(ID string) (*Account, error)
	ReadByEmail(email string) (*Account, error)
	Create(a *Account, password string) error
	Update(a *Account) error
	Delete(ID string) error
	Confirm(token string) (*Account, error)
	GeneratePasswordToken(email string) (*Account, error)
	UpdatePassword(string, string) (*Account, error)
	Migrate() error
	Truncate() error
	Close() error
}

message Account {
string uuid = 1  [(gogoproto.moretags) = "db:\"uuid\""];
string name = 2 [(gogoproto.moretags) = "db:\"name\""];
string email = 3 [(gogoproto.moretags) = "db:\"email\""];
bool confirmed_and_active = 4 [(gogoproto.moretags) = "db:\"confirmed_and_active\""];
google.protobuf.Timestamp member_since = 5 [(gogoproto.moretags) = "db:\"member_since\"", (gogoproto.stdtime) = true, (gogoproto.nullable) = false]; // API format becomes RFC 3339
bool support = 6 [(gogoproto.moretags) = "db:\"support\""];
string phonenumber = 7 [(gogoproto.moretags) = "db:\"phonenumber\""];
string photo_url = 8 [(gogoproto.moretags) = "db:\"photo_url\""];
string password = 9 [(gogoproto.moretags) = "db:\"password\""];
}


type Account struct {
	UUID               string `db:"uuid"`
	Name               string `db:"name"`
	Email              string `validate:"required"`
	HashedPassword     string `db:"hashed_password"`
	ConfirmationToken  string
	PasswordResetToken string
	Metadata           map[string]string
	CreatedAt          time.Time `db:"created_at"`
}

func (a *Account) Valid() error {
	return valid.Struct(a)
}

func (a *Account) HashPassword(password string) error {
	if password == "" {
		return ErrNoPasswordGiven
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.HashedPassword = string(hash[:])

	return nil
}

func (a *Account) ComparePasswordToHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.HashedPassword), []byte(password))
}

func EmailExists(db Database, a *Account) error {
	a, err := db.ReadByEmail(a.Email)
	if err != nil && err != ErrAccountNotFound {
		return err
	}

	if a != nil {
		return ErrEmailExists
	}

	return nil
}

func DatabaseFromEnv() Database {
	var conn Database
	var err error

	pg := os.Getenv("POSTGRESQL_URL")
	if pg != "" {
		pgConn := &PostgreSQL{}
		err = pgConn.Connect(pg)
		conn = pgConn
	}

	if conn == nil {
		panic(ErrNoDatabase)
	}

	if err != nil {
		panic(err)
	}

	return conn
}
