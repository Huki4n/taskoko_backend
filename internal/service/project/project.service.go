package service

import (
	"awesomeProject/ent"
	models "awesomeProject/internal/models/project"
	repository "awesomeProject/internal/repository/project"
	"context"
	"errors"
	"github.com/samber/lo"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// Бизнес-логика для создания проекта
func (s *ProjectService) CreateProject(ctx context.Context, request models.CreateProjectRequest, userID int) (*models.CreateProjectResponse, error) {
	if request.ProjectName == "" {
		return nil, errors.New("project name is required")
	}

	project, err := s.repo.CreateProject(ctx, request.ProjectName, userID)
	if err != nil {
		return nil, err
	}

	return &models.CreateProjectResponse{
		ID:   project.ID,
		Name: project.Name,
	}, nil
}

func (s *ProjectService) GetProject(ctx context.Context, projectId int) (*models.GetProjectResponse, error) {
	project, err := s.repo.GetProject(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &models.GetProjectResponse{
		ID:      project.ID,
		Name:    project.Name,
		Columns: project.Edges.Columns,
		Users: lo.Map(project.Edges.Users, func(user *ent.User, _ int) models.ProjectUser {
			return models.ProjectUser{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			}
		}),
	}, nil
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]models.GetProjectsResponse, error) {
	projects, err := s.repo.GetProjects(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Map(projects, func(p *ent.Project, _ int) models.GetProjectsResponse {
		return models.GetProjectsResponse{
			ID:   p.ID,
			Name: p.Name,
			Users: lo.Map(p.Edges.Users, func(user *ent.User, _ int) models.ProjectUser {
				return models.ProjectUser{
					ID:    user.ID,
					Name:  user.Name,
					Email: user.Email,
				}
			}),
		}
	}), nil
}
