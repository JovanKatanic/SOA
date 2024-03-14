package service

import (
	"fmt"
	"stakeholders_service/model"
	"stakeholders_service/repository"
)

type AuthenticationService struct {
	UserRepository   *repository.UserRepostiroy
	PersonRepository *repository.PersonRepository
}

func (service *AuthenticationService) RegisterTourist(AccountRegistration *model.AccountRegistration) error {

	if service.UserRepository.FindByUsername(AccountRegistration.Username) != nil {
		return fmt.Errorf("ne mozes registrovati korisnika jer postoji korisnik sa tim usernamom")
	}
	user := model.User{
		Username: AccountRegistration.Username,
		Password: AccountRegistration.Password,
		Role:     2,
		IsActive: true,
	}
	err1 := service.UserRepository.CreateUser(&user)
	user2 := model.User{}
	//service.UserRepository.DatabaseConnection.Find(&user2, "\"Username\" = ?", AccountRegistration.Username)
	service.UserRepository.DatabaseConnection.Table("stakeholders.Users").Find(&user2, "\"Username\" = ?", AccountRegistration.Username)

	person := model.Person{
		UserId:  user2.Id,
		Name:    AccountRegistration.Name,
		Surname: AccountRegistration.Surname,
		Email:   AccountRegistration.Email,
	}
	err2 := service.PersonRepository.CreatePerson(&person)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}
