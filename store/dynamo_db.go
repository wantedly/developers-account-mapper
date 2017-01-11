package store

import (
	"fmt"

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
		users = append(users, models.NewUser(*item["LoginName"].S, *item["GitHubUsername"].S, *item["SlackUsername"].S, *item["SlackUserId"].S))
	}

	return users, nil
}

func (d *DynamoDB) AddUser(user *models.User) error {
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

func (d *DynamoDB) GetUserByLoginName(loginName string) (*models.User, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(accountMapTable),
		Key: map[string]*dynamodb.AttributeValue{
			"LoginName": {
				S: aws.String(loginName),
			},
		},
		AttributesToGet: []*string{
			aws.String("GitHubUsername"),
			aws.String("SlackUsername"),
			aws.String("SlackUserId"),
		},
		ConsistentRead: aws.Bool(true),
	}

	resp, err := d.db.GetItem(params)

	if err != nil {
		return nil, err
	}


	if len(resp.Item) != 1 {
		return nil, fmt.Errorf("%s is not registered yet", loginName)
	}

	user := &models.User{
		LoginName:      loginName,
		GitHubUsername: *resp.Item["GitHubUsername"].S,
		SlackUserId:    *resp.Item["SlackUserId"].S,
		SlackUsername:  *resp.Item["SlackUsername"].S,
	}

	return user, nil
}
