package main

import (
	"fmt"
	"test-wrap-error/library/resource"
	"test-wrap-error/model/dao/db"
)

func init() {
	resource.InitMysql()
}

// dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层？
// 可以在 dao 层根据不同的查询场景，来判断是否抛给上层
// 如 检查用户 id 是否存在，则可以在 dao 层处理这个错误
// 通过用户 id 获取信息，没有查到的情况下，可以 Wrap 这个 error，抛给上层

func main() {
	userDao := db.NewUserDao()
	userID := 1
	exist, err := userDao.CheckUserExist(userID)
	if err != nil {
		fmt.Printf("check user exist err, %+v\n", err)
	}

	if !exist {
		fmt.Printf("user %d not exist\n", userID)
		return
	}
	fmt.Printf("user %d exist\n", userID)

	userInfo, err := userDao.GetUserInfo(userID)
	if err != nil {
		fmt.Printf("get user info err, %+v\n", err)
	}
	fmt.Printf("user info, id: %d, name: %s, phone: %d\n", userInfo.ID, userInfo.Name, userInfo.Phone)
}
