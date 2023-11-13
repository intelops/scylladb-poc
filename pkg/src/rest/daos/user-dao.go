package daos

import (
	"log"

	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/daos/clients/scylla"
	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/models"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

const (
	TABLENAME = "users"
)

type UserDao struct {
	DbManager *scylla.Manager
	DBsession gocqlx.Session
}

func NewUserDao() (*UserDao, error) {
	DbManager, err := scylla.NewManager()
	if err != nil {
		log.Fatalf("error while loading scylla manager: %s", err)
		return nil, err
	}
	session, err := DbManager.Connect()
	if err != nil {
		log.Fatalf("error while connecting to scylla: %s", err)
		return nil, err

	}
	err = DbManager.CreateTable(&session, models.CreateUserTable, TABLENAME)
	if err != nil {
		log.Fatalf("unable to create %s table due to: %s", TABLENAME, err)
		return nil, err
	}
	return &UserDao{DbManager: DbManager, DBsession: session}, nil
}

func (u *UserDao) CreateUser(m *models.User) (*models.User, error) {
	stmt, names := qb.Insert(TABLENAME).Columns("id", "firstname", "lastname", "password", "email", "phone").ToCql()
	log.Println(m)
	err := gocqlx.Session.Query(u.DBsession, stmt, names).BindStruct(m).ExecRelease()

	if err != nil {
		log.Fatal("failed to insert data: ", err)
		return nil, err
	}
	log.Println("data inserted successfully")
	return m, nil
}

func (u *UserDao) GetUser(id string) (*models.User, error) {
	stmt, names := qb.Select(TABLENAME).Where(qb.Eq("id")).ToCql()
	m := &models.User{}
	err := gocqlx.Session.Query(u.DBsession, stmt, names).Bind(id).Get(m)
	if err != nil {
		log.Fatal("failed to get data: ", err)
		return nil, err
	}
	log.Println("data fetched successfully")
	return m, nil
}

func (u *UserDao) GetUsers() ([]*models.User, error) {
	stmt, names := qb.Select(TABLENAME).ToCql()
	var m []*models.User
	err := gocqlx.Session.Query(u.DBsession, stmt, names).Select(&m)
	if err != nil {
		log.Fatal("failed to get data: ", err)
		return nil, err
	}
	log.Println("data fetched successfully")
	return m, nil
}

func (u *UserDao) DeleteUser(id string) error {
	stmt, names := qb.Delete(TABLENAME).Where(qb.Eq("id")).ToCql()
	err := gocqlx.Session.Query(u.DBsession, stmt, names).Bind(id).ExecRelease()
	if err != nil {
		log.Fatal("failed to delete data: ", err)
		return err
	}
	log.Println("data deleted successfully")
	return nil
}

func (u *UserDao) UpdateUser(m *models.User) (*models.User, error) {
	stmt, names := qb.Update(TABLENAME).Set("firstname", "lastname", "password", "email", "phone").Where(qb.Eq("id")).ToCql()
	err := gocqlx.Session.Query(u.DBsession, stmt, names).BindStruct(m).ExecRelease()
	if err != nil {
		log.Fatal("failed to update data: ", err)
		return nil, err
	}
	log.Println("data updated successfully")
	return m, nil

}