package domain

import "time"

type Project struct {
	Id     int32 `json:"id"`
	UserId int32 `json:"user_id"`

	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Color string `json:"color"`

	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ProjectsRepository interface {
	FetchByUserId(int32) ([]*Project, error)
	FetchByUserIdTrashed(int32) ([]*Project, error)
	FetchById(int32) (*Project, error)

	Insert(*Project) error
	Update(Project) error
	Remove(int32) error
}

type ProjectsInteractor interface {
	Active(userId int32, trashed bool) (ProjectListResponse, error)
	Info(userId int32, projectId int32) (ProjectInfoResponse, error)
	Trashed(userId int32) (ProjectListResponse, error)
	Create(project Project) (IdResponse, error)
	Remove(userId, projectId int32) (StatusResponse, error)
}

// Responses (only for UseCase layer)

type ProjectListResponse struct {
	StatusCode string
	Projects   []*Project
}

type ProjectInfoResponse struct {
	StatusCode string
	Project    *Project
}
