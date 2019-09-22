package repository

import (
	"database/sql"
	"fmt"
	"os"

	model "github.com/Gustibimo/fave/api/model"
	"github.com/labstack/gommon/log"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

var db *sql.DB

type postgresMerchantRepository struct {
	Conn *sql.DB
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewPostgresMerchantRepository(Conn *sql.DB) MerchantRepository {
	return &postgresMerchantRepository{Conn}
}

func (m *postgresMerchantRepository) fetch(query string, args ...interface{}) ([]*model.Merchants, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Merchants, 0)
	for rows.Next() {
		t := new(model.Merchants)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Address,
			&t.City,
			&t.Category,
			&t.Rating,
			&t.Logo,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *postgresMerchantRepository) Fetch(cursor string, num int64) ([]*model.Merchants, error) {

	query := `SELECT id,name,address, city, category, rating, logo
  						FROM merchants WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, num)

}

func (m *postgresMerchantRepository) GetByID(id int64) (*model.Merchants, error) {
	query := `SELECT id,name,address, city, category, rating, logo
	FROM merchants WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &model.Merchants{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, model.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *postgresMerchantRepository) Store(a *model.Merchants) (int64, error) {

	query := `INSERT  merchants SET name=? , address=? , city=?, category=? , rating=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {

		return 0, err
	}

	log.Debug("Name: ", a.Name)
	res, err := stmt.Exec(a.Name, a.Address, a.City, a.Category, a.Rating)
	if err != nil {

		return 0, err
	}
	return res.LastInsertId()
}

func (m *postgresMerchantRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM merchants WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return false, err
	}
	res, err := stmt.Exec(id)
	if err != nil {

		return false, err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		log.Error(err)
		return false, err
	}

	return true, nil
}
func (m *postgresMerchantRepository) Update(mr *model.Merchants) (*model.Merchants, error) {
	query := `UPDATE merchants set name=?, address=?, city=?, category=? WHERE ID = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, nil
	}

	res, err := stmt.Exec(mr.Name, mr.Address, mr.City, mr.Category, mr.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		log.Error(err)
		return nil, err
	}

	return mr, nil
}
