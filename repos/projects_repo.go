package repos

import (
	"database/sql"
	"microservice/app/core"
	"microservice/domain"
)

type ProjectsRepo struct {
	log core.Logger
	db  *sql.DB
}

func NewProjectsRepo(log core.Logger, db *sql.DB) *ProjectsRepo {
	return &ProjectsRepo{log: log, db: db}
}

func (r *ProjectsRepo) FetchByUserId(userId int32) ([]*domain.Project, error) {
	query := `select 
    			id, 
    			name,
    			"desc", 
    			color,
    			created_at,
    			updated_at,
    			deleted_at
			from projects
			where user_id=$1 and deleted_at is null`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var result []*domain.Project
	for rows.Next() {
		item := &domain.Project{
			UserId: userId,
		}
		err := rows.Scan(&item.Id,
			&item.Name,
			&item.Desc,
			&item.Color,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

func (r *ProjectsRepo) FetchByUserIdTrashed(userId int32) ([]*domain.Project, error) {
	query := `select 
    			id, 
    			name,
    			"desc", 
    			color,
    			created_at,
    			updated_at,
    			deleted_at
			from projects
			where user_id=$1 and deleted_at is not null`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var result []*domain.Project
	for rows.Next() {
		item := &domain.Project{
			UserId: userId,
		}
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Desc,
			&item.Color,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

func (r *ProjectsRepo) FetchById(id int32) (*domain.Project, error) {
	var item = &domain.Project{
		Id: id,
	}
	query := `select 
    			user_id,
    			name,
    			"desc", 
    			color,
    			created_at,
    			updated_at,
    			deleted_at
			from projects
			where id=$1
			limit 1`

	err := r.db.QueryRow(query, id).Scan(
		&item.UserId,
		&item.Name,
		&item.Desc,
		&item.Color,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.DeletedAt)
	switch err {
	case nil:
		return item, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (r *ProjectsRepo) Insert(item *domain.Project) error {
	query := `INSERT INTO projects (user_id, name, "desc", color) VALUES ($1, $2, $3, $4) returning id;`
	err := r.db.QueryRow(query, item.UserId, item.Name, item.Desc, item.Color).Scan(&item.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProjectsRepo) Update(item domain.Project) error {
	query := `UPDATE projects 
				SET name=$2, "desc"=$3, color=$4, updated_at=now()
				WHERE id=$1`
	_, err := r.db.Exec(query, item.Id, item.Name, item.Desc, item.Color)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProjectsRepo) Remove(id int32) error {
	query := `UPDATE projects 
				SET deleted_at=now()
				WHERE id=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
