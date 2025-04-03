package projectModels

import "awesomeProject/ent"

type ProjectUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Create Project Models
type CreateProjectRequest struct {
	ProjectName string `json:"projectName"`
}

type CreateProjectResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Get Project Models
type GetProjectRequest struct {
	ProjectId int `json:"projectId"`
}

type GetProjectResponse struct {
	ID      int                  `json:"id"`
	Name    string               `json:"name"`
	Users   []ProjectUser        `json:"users"`
	Columns []*ent.ProjectColumn `json:"columns"`
}

// Get Projects Models
type GetProjectsResponse struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Users []ProjectUser `json:"users"`
}
