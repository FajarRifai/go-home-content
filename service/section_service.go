package services

import (
	"go-home-content/models"
	"go-home-content/repository"
)

type SectionService struct {
	Repo *repository.SectionRepository
}

func (s *SectionService) CreateSection(section models.CreateSection) (int64, error) {
	return s.Repo.CreateSection(section)
}

func (s *SectionService) GetSections() ([]models.ListSection, error) {
	return s.Repo.GetSections()
}

func (s *SectionService) GetSectionById(id int) (models.Section, error) {
	return s.Repo.GetSectionById(id)
}

func (s *SectionService) GetSectionDetailById(id int) ([]string, error) {
	return s.Repo.GetSectionDetailById(id)
}
