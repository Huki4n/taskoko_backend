package repository

import (
	"awesomeProject/ent"
	"awesomeProject/ent/project"
	"context"
	"fmt"
	"log"
)

type ProjectRepository struct {
	client *ent.Client
}

func NewProjectRepository(client *ent.Client) *ProjectRepository {
	return &ProjectRepository{client: client}
}

// CreateProject логика взаимодействия с бд для создания проекта
func (r *ProjectRepository) CreateProject(ctx context.Context, projectName string, userID int) (*ent.Project, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	createProject, err := tx.Project.Create().
		SetName(projectName).
		Save(ctx)

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Printf("Create project rollback error: %v", err)
			return nil, err
		}
		return nil, err
	}

	_, err = tx.ProjectColumn.CreateBulk(
		tx.ProjectColumn.Create().SetName("To Do").SetColor("#9d66d5").SetOrder(1).SetProject(createProject),
		tx.ProjectColumn.Create().SetName("In Work").SetColor("#3361ff").SetOrder(2).SetProject(createProject),
		tx.ProjectColumn.Create().SetName("Review").SetColor("#ffcd43").SetOrder(3).SetProject(createProject),
		tx.ProjectColumn.Create().SetName("Done").SetColor("#98db7c").SetOrder(4).SetProject(createProject),
	).Save(ctx)

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Printf("Create columns rollback error: %v", err)
			return nil, err
		}
		return nil, err
	}

	_, err = tx.Project.Update().
		Where(project.ID(createProject.ID)).
		AddUserIDs(userID).
		Save(ctx)

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Printf("Project update rollback error: %v", err)
			return nil, err
		}
		return nil, err
	}

	if commitErr := tx.Commit(); commitErr != nil {
		log.Printf("transaction commit failed: %v", commitErr)
		return nil, commitErr
	}

	return createProject, nil
}

// GetProject логика взаимодействия с бд для получения проекта по id
func (r *ProjectRepository) GetProject(ctx context.Context, projectID int) (*ent.Project, error) {
	// Ищем проект по ID
	getProject, err := r.client.Project.Query().
		Where(project.ID(projectID)).
		WithColumns().
		WithUsers(
			func(q *ent.UserQuery) {
				q.Select("id", "name", "email")
			}).
		Only(ctx)

	if err != nil {
		log.Printf("Get project query error: %v", err)
		return nil, err
	}

	return getProject, nil
}

// GetProjects логика взаимодействия с бд для получения всех проектов юзера
func (r *ProjectRepository) GetProjects(ctx context.Context) ([]*ent.Project, error) {
	projects, err := r.client.Project.Query().WithUsers(
		func(q *ent.UserQuery) {
			q.Select("id", "name", "email")
		}).All(ctx)

	if err != nil {
		log.Printf("Get projects query error: %v", err)
		return nil, err
	}

	return projects, nil
}

// DeleteProject логика взаимодействия с бд для получения всех проектов юзера
func (r *ProjectRepository) DeleteProject(ctx context.Context, projectID int) (string, error) {
	_, err := r.client.Project.Query().Where(project.ID(projectID)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", fmt.Errorf("project with ID %d not found", projectID)
		}

		log.Printf("Error while querying project: %v", err)
		return "", err
	}

	// Если проект найден, удаляем его
	_, err = r.client.Project.Delete().Where(project.ID(projectID)).Exec(ctx)
	if err != nil {
		log.Printf("Error while deleting project: %v", err)
		return "", err
	}

	return "Project has been deleted", nil
}
