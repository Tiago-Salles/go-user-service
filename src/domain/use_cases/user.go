package use_cases


import "github.com/Tiago-Salles/go-user-service/src/domain/entities"

type UserUseCase struct{
	user entities.UserEntity
}

func (u *UserUseCase) Register () entities.UserEntity{
	u.user.Name = "Tiago"
	u.user.Email = "tiagosalles20@gmail.com"
	return u.user
}