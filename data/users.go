package data

type User struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Email string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Admin	bool	`bson:"admin"`
}