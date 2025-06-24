package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	DB           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(db *sql.DB) UowInterface {
	return &Uow{
		DB:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	return u.Repositories[name](u.Tx), nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow UowInterface) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx
	if err := fn(u); err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) CommitOrRollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to commit")
	}

	if err := u.Tx.Commit(); err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return fmt.Errorf("commit error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	u.Tx = nil
	return nil
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	if err := u.Tx.Rollback(); err != nil {
		return err
	}

	u.Tx = nil
	return nil
}
