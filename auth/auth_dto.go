package auth

type AuthResponseDTO struct {
	Token string `json:"token"`
}

type CreateAccountDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (ca CreateAccountDTO) toEntity() CreateAccount {
	return CreateAccount{
		Email:     ca.Email,
		Password:  ca.Password,
		FirstName: ca.FirstName,
		LastName:  ca.LastName,
	}
}
