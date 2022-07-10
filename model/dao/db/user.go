package db

import (
	"database/sql"
	"github.com/pkg/errors"
	"test-wrap-error/library/resource"
)

type UserDao struct {
	*BaseDbDao
}

type User struct {
	ID    int
	Phone int64
	Name  string
}

func NewUserDao() *UserDao {
	return &UserDao{
		newDbDao(resource.MySQLTestDB1),
	}
}

func (dao *UserDao) CheckUserExist(id int) (bool, error) {
	_, err := dao.BaseSelectOne("select * from user where id=? limit 1", id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return true, err
}

func (dao *UserDao) GetUserInfo(id int) (*User, error) {
	data, err := dao.BaseSelectOne("select * from user where id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	phone := data["phone"].(int64)
	name := string(data["name"].([]byte))
	user := &User{
		ID:    id,
		Phone: phone,
		Name:  name,
	}
	return user, err
}
