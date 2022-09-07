package adapter

import (
	"context"

	"cloud.google.com/go/bigquery"
	config "github.com/chocogem/bigquery-golang/pkg/config"
	domain "github.com/chocogem/bigquery-golang/pkg/domain"
	repository "github.com/chocogem/bigquery-golang/pkg/repository"
	"google.golang.org/api/iterator"
)

type userDatabase struct {
	bigQueryConfig domain.BigQueryConfig
}

type user struct {
	UserId     string `bigquery:"user_id"`
	UserName   string `bigquery:"user_name"`
	FirstName  string `bigquery:"first_name"`
	LastName   string `bigquery:"last_name"`
	ExpireDate string `bigquery:"expire_date"`
}

func NewUserRepository(cfg config.Config) (repository.UserRepository, error) {
	return &userDatabase{
		bigQueryConfig: domain.BigQueryConfig{ProjectId: cfg.BQProjectId, DatasetName: cfg.BQDatasetName}}, nil

}

func (c *userDatabase) FindAll() ([]domain.User, error) {

	sql := `SELECT user_id ,user_name ,first_name ,last_name , FORMAT_DATE("%y-%m-%d",expire_date) expire_date FROM ` +
		"`" + c.bigQueryConfig.ProjectId + "." + c.bigQueryConfig.DatasetName + ".users` "

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, c.bigQueryConfig.ProjectId)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	query := client.Query(sql)
	iter, err := query.Read(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, iter.TotalRows)
	var i = 0
	for {
		var row user
		err := iter.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		users[i] = domain.NewUser(row.UserId, row.UserName, row.FirstName, row.LastName, row.ExpireDate)
		i++
	}

	return users, nil
}
