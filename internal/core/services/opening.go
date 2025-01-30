package services

import (
	"fmt"

	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/ports"
)

type OpeningService struct {
	OpeningRepository ports.OpeningRepository
}

func NewOpeningService(repo ports.OpeningRepository) *OpeningService {
	return &OpeningService{
		OpeningRepository: repo,
	}
}
func (o *OpeningService) GetOpeningByName(name string) []domain.Opening {
	fmt.Println("got here in service")
	openings, err := o.OpeningRepository.GetOpeningByName(name)
	if err != nil {
		fmt.Println("error here")
		return nil
	}
	return openings
}
