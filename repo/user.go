package repo

import (
	"database/sql"
	"fmt"
	"mynozom/model"
	"mynozom/utils"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	dbs []*gorm.DB
}

func NewUserRepo(dbs []*gorm.DB) UserRepo {
	return UserRepo{
		dbs: dbs,
	}
}

func (ur *UserRepo) GetByCode(code *string) (*model.User, error) {
	row := ur.dbs[0].Raw("EXEC GetEmp @EmpCode = ?", code).Row()
	user, err := scanUserResult(row)
	if utils.CheckErr(&err) {
		return nil, err
	}

	fmt.Println("user")
	fmt.Println(user)

	return user, nil
}

func scanUserResult(row *sql.Row) (*model.User, error) {
	var user model.User
	err := row.Scan(&user.EmpName, &user.EmpPassword, &user.EmpCode, &user.SecLevel)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
