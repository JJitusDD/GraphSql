package persistence

import (
	"context"
	"fmt"
	"project-test/internal/domain/model/entities"
	"project-test/internal/domain/repository"
	"project-test/pkg/agent/postgres"
	"strings"

	"github.com/sirupsen/logrus"
)

type User struct {
	pg  *postgres.Postgres
	log *logrus.Logger
}

func NewUser(pg *postgres.Postgres, l *logrus.Logger) repository.IUser {
	return &User{
		pg:  pg,
		log: l,
	}
}

func (al *User) Create(ctx context.Context, in *entities.User) (ou *entities.User, err error) {
	tx := al.pg.DB.Where(&entities.User{})
	al.log.WithFields(logrus.Fields{
		"in": in,
	}).Info("input create bank statement logs")

	err = tx.Create(&in).Scan(&ou).Error
	if err != nil {
		return nil, err
	}

	return
}

func (r *User) Update(ctx context.Context, id string, in *entities.User) (ou *entities.User, err error) {
	tx := r.pg.DB.Model(&entities.User{})
	_, err = r.FindById(ctx, id)
	if err != nil {
		return
	}

	err = tx.Where("id = ?", id).Updates(&in).Error
	if err != nil {
		return
	}
	return in, nil
}

func (r *User) FindById(ctx context.Context, id string) (o *entities.User, e error) {
	tx := r.pg.DB.Model(&entities.User{})
	e = tx.Where("id = ?", id).First(&o).Error
	if e != nil {
		return nil, e
	}
	return
}

func (r *User) Find(ctx context.Context, keysAndValues ...string) (ou []*entities.User, err error) {
	db := r.pg.DB.Model(&entities.User{}).WithContext(ctx)
	if len(keysAndValues) > 0 {
		for i := 0; i < len(keysAndValues); i = i + 3 {
			if keysAndValues[i+1] == "in" {
				stringArray := strings.ReplaceAll(keysAndValues[i+2], `(`, ``)
				stringArray = strings.ReplaceAll(stringArray, `)`, ``)
				whereIns := strings.Split(stringArray, ",")
				db = db.Where(fmt.Sprintf("%s %s ?", keysAndValues[i], keysAndValues[i+1]), whereIns)

			} else {
				db = db.Where(fmt.Sprintf("%s %s ?", keysAndValues[i], keysAndValues[i+1]), keysAndValues[i+2])
			}
		}
	}

	err = db.Find(&ou).Order("created_date asc").Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *User) Delete(ctx context.Context, id string) (e error) {
	tx := r.pg.DB.Model(&entities.User{})
	_, e = r.FindById(ctx, id)
	if e != nil {
		return
	}

	e = tx.Where("id = ?", id).Delete(&entities.User{}).Error
	if e != nil {
		return
	}

	return nil
}
