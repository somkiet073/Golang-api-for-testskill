package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Project = project
type Project struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255);" json:"title"`
	Description string    `gorm:"type:varchar(255);" json:"description"`
	User        User      `json:"user"`
	UserID      uint32    `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare = prepare
func (p *Project) Prepare() {
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// =======================================================================

// CreateProject  createProject
func (p *Project) CreateProject(db *gorm.DB) (*Project, error) {
	err := db.Debug().Create(&p).Error
	if err != nil {
		return &Project{}, err
	}
	return p, nil
}

// FindAllProject = findAllProject
func (p *Project) FindAllProject(db *gorm.DB) (*[]Project, error) {
	var project []Project
	err := db.Debug().Model(&Project{}).Limit(100).Find(&project).Error
	if err != nil {
		return &[]Project{}, err
	}
	return &project, nil
}

// FindProjectByID = findProjectByID
func (p *Project) FindProjectByID(db *gorm.DB, pid uint64) (*Project, error) {
	err := db.Debug().Model(&Project{}).Where("id=?", pid).Take(&p).Error
	if err != nil {
		return &Project{}, err
	}

	return p, nil
}

// UpdateProject = updateProject
func (p *Project) UpdateProject(db *gorm.DB) (*Project, error) {

	var err error

	var projects = Project{
		Title:       p.Title,
		Description: p.Description,
		UpdatedAt:   p.UpdatedAt,
	}

	err = db.Debug().Model(&Project{}).Where("id=?", p.ID).Updates(projects).Error

	if err != nil {
		return &Project{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&Project{}).Where("id=?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Project{}, nil
		}
	}

	return p, nil
}

// DeleteProject = deleteProject
func (p *Project) DeleteProject(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Project{}).Where("id = ? and user_id = ? ", pid, uid).Take(&Project{}).Delete(&Project{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Project no found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
