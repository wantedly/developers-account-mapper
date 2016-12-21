package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/wantedly/developers-account-mapper/models"
)

const (
	accountMapTable = "DevelopersAccountMap"
)

type DynamoDB struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {
	db := dynamodb.New(session.New(&aws.Config{}))

	return &DynamoDB{
		db: db,
	}
}

func (d *DynamoDB) ListUsers() ([]*models.User, error) {
	resp, err := d.db.Scan(&dynamodb.ScanInput{
		TableName: aws.String(accountMapTable),
	})
	if err != nil {
		return nil, err
	}

	var users []*models.User

	for _, item := range resp.Items {
		users = append(users, models.NewUser(*item["LoginName"].S, *item["GitHubUsername"].S, *item["SlackUsername"].S, ""))
	}

	return users, nil
}

func (d *DynamoDB) AddUser(user *models.User) (error) {
	_, err := d.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(accountMapTable),
		Item: map[string]*dynamodb.AttributeValue{
			"LoginName": {
				S: aws.String(user.LoginName),
			},
			"GitHubUsername": {
				S: aws.String(user.GitHubUsername),
			},
			"SlackUsername": {
				S: aws.String(user.SlackUsername),
			},
			"SlackUserId": {
				S: aws.String(user.SlackUserId),
			},
		},
	})
	return err
}
