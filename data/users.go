package data

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Email string `bson:"email" json:"email"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Admin	bool	`bson:"admin" json:"admin"`
}