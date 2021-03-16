package integration_tests

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

func SetupSQS(t *testing.T) (queueUrl string, awsSession *session.Session) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "localstack/localstack:0.12.5",
		ExposedPorts: []string{"4566/tcp", "4571/tcp"},
		Env: map[string]string{
			"SERVICES":    "sqs",
			"DEBUG":       "1",
			"PORT_WEB_UI": "8080",
		},
		WaitingFor: wait.ForListeningPort("4566/tcp"),
	}
	localStackC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	t.Cleanup(func() {
		t.Log("terminating localstack container")
		err := localStackC.Terminate(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})

	host, err := localStackC.Host(ctx)
	assert.NoError(t, err)
	port, err := localStackC.MappedPort(ctx, "4566/tcp")
	assert.NoError(t, err)

	sess := testAWSSession(fmt.Sprintf("http://%s:%d", host, port.Int()))

	sqsClient := sqs.New(sess)
	createQueueOutput, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String("payment_intent:" + uuid.New().String()),
	})
	assert.NoError(t, err)

	return *createQueueOutput.QueueUrl, sess
}

func testAWSSession(endPoint string) *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:     aws.String("us-east-1"),
			MaxRetries: aws.Int(3),
		},
	}))
	sess.Config.Endpoint = aws.String(endPoint)
	sess.Config.Credentials = credentials.NewStaticCredentials("test", "test", "")
	return sess
}
