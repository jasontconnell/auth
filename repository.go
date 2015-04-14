package shared

import (
	"repository"
	"shared/data"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"crypt"
)

type UserRepository struct {
	repository.Repository
}

func NewUserRepository(config repository.Configuration) *UserRepository {
	repo := new(UserRepository)
	repo.Initialize(config.DatabaseServer, config.Database, config.DatabaseUser, config.DatabasePassword)
	return repo
}

func (repo *UserRepository) LoginUser(username, password string) (*data.User,error) {
	key := "random array of characters"
	collection := repo.OpenCollection("users")
	user := data.User{}
		
	if err := collection.Find(bson.M{ "username": username }).One(&user); err == nil {
		decrypted := crypt.Decrypt(key, user.Password)
		if password != decrypted {
			return nil, errors.New("Invalid credentials")
		} else {
			return &user, nil
		}
	} else {
		return nil, err
	}
}