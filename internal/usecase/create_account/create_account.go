package createaccount

import (
	"github.com/PenaSoftwareCraftsman/fc-ms-wallet/internal/entity"
	"github.com/PenaSoftwareCraftsman/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientId string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	ClientGateway  gateway.ClientGateway
	AccountGateway gateway.AccountGateway
}

func NewCreateAccountUseCase(
	clientGateway gateway.ClientGateway,
	accountGateway gateway.AccountGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		ClientGateway:  clientGateway,
		AccountGateway: accountGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {

	client, err := uc.ClientGateway.Get(input.ClientId)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
