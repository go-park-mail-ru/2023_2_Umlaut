package repository

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/pashagolub/pgxmock/v3"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	id          = 1
	m           = "Test message"
	b           = false
	t           = time.Now()
	testMessage = core.Message{
		SenderId:  &id,
		DialogId:  &id,
		Text:      &m,
		IsRead:    &b,
		CreatedAt: &t,
	}
)

func TestMessagePostgres_CreateMessage(t *testing.T) {
	mock, mockErr := pgxmock.NewPool()
	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}
	defer mock.Close()

	messageRepo := NewMessagePostgres(mock)

	mock.ExpectQuery(`INSERT INTO "message"`).
		WithArgs(testMessage.SenderId, testMessage.RecipientId, testMessage.DialogId, testMessage.Text).
		WillReturnRows(pgxmock.NewRows([]string{"id", "dialog_id", "sender_id", "recipient_id", "message_text", "is_read", "created_at"}).
			AddRow(&id, &id, &id, &id, &m, &b, nil))

	newMessage, err := messageRepo.CreateMessage(context.Background(), testMessage)

	assert.NoError(t, err)
	assert.NotNil(t, newMessage)
}

func TestMessagePostgres_UpdateMessage(t *testing.T) {
	mock, mockErr := pgxmock.NewPool()
	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}
	defer mock.Close()

	messageRepo := NewMessagePostgres(mock)

	mock.ExpectQuery(`UPDATE "message"`).
		WithArgs(testMessage.Text, testMessage.IsRead).
		WillReturnRows(pgxmock.NewRows([]string{"id", "dialog_id", "sender_id", "recipient_id", "message_text", "is_read", "created_at"}).
			AddRow(&id, &id, &id, &id, &m, &b, nil))

	updatedMessage, err := messageRepo.UpdateMessage(context.Background(), testMessage)

	assert.NoError(t, err)
	assert.NotNil(t, updatedMessage)
}

func TestMessagePostgres_GetDialogMessages(t *testing.T) {
	mock, mockErr := pgxmock.NewPool()
	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}
	defer mock.Close()

	messageRepo := NewMessagePostgres(mock)

	dialogID := 1
	testMessages := []core.Message{testMessage, testMessage}

	mock.ExpectQuery(`SELECT`).
		WithArgs(dialogID, dialogID, dialogID, dialogID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "dialog_id", "sender_id", "recipient_id", "message_text", "is_read", "created_at"}).
			AddRow(testMessages[0].Id, testMessages[0].DialogId, testMessages[0].SenderId, testMessages[0].RecipientId, testMessages[0].Text, testMessages[0].IsRead, testMessages[0].CreatedAt).
			AddRow(testMessages[1].Id, testMessages[1].DialogId, testMessages[1].SenderId, testMessages[1].RecipientId, testMessages[1].Text, testMessages[1].IsRead, testMessages[1].CreatedAt))

	messages, err := messageRepo.GetDialogMessages(context.Background(), dialogID, dialogID)

	assert.NoError(t, err)
	assert.NotNil(t, messages)
}
