package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}

type Session struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	UserId bson.ObjectId `json:"userId" bson:"userId"`
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
}


//type AccountModel struct {
//	DB *mgo.Database
//}
//
//
//func (accountModel *AccountModel) CheckUsernameAndPassword(username,password string) bool{
//	var account Account
//	err := accountModel.DB.C("account").Find(bson.M{
//		"username": username,
//	}).One(&account)
//	if err != nil{
//		return false
//	} else {
//		return bcrypt.CompareHashAndPassword([]byte(account.Password),[]byte(password)) == nil
//	}
//}

