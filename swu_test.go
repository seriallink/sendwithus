package swu

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	test = New(os.Getenv("SWU_KEY"))
}

var test *Api

func TestTemplates(t *testing.T) {
	_, err := test.Emails()
	assert.NoError(t, err)
}

func TestGetTemplate(t *testing.T) {
	_, err := test.GetTemplate(os.Getenv("SWU_TEMPLATE_ID"))
	assert.NoError(t, err)
}

func TestGetTemplateVersion(t *testing.T) {
	_, err := test.GetTemplateVersion(os.Getenv("SWU_TEMPLATE_ID"), os.Getenv("SWU_VERSION_ID"))
	assert.NoError(t, err)
}

func TestCreateTemplate(t *testing.T) {
	_, err := test.CreateTemplate(&Version{
		Name:    "Test",
		Subject: "Test",
		Text:    "ALOHA",
	})
	assert.NoError(t, err)
}

func TestCreateTemplateVersion(t *testing.T) {
	_, err := test.CreateTemplateVersion(os.Getenv("SWU_TEMPLATE_ID"), &Version{
		Name:    "Test",
		Subject: "Test",
		Text:    "Hello Template Version!",
	})
	assert.NoError(t, err)
}

func TestUpdateTemplateVersion(t *testing.T) {
	_, err := test.UpdateTemplateVersion(os.Getenv("SWU_TEMPLATE_ID"), os.Getenv("SWU_VERSION_ID"),
		&Version{
			Name:    "Test",
			Subject: "Test",
			Text:    "Hello New Template Version!",
		})
	assert.NoError(t, err)
}

func TestSend(t *testing.T) {
	email := &Email{
		Id: os.Getenv("SWU_TEMPLATE_ID"),
		Sender: &Sender{
			Recipient: &Recipient{
				Name:    "NoReply",
				Address: os.Getenv("SWU_SENDER_EMAIL"),
			},
		},
		Recipient: &Recipient{
			Name:    "John Doe",
			Address: os.Getenv("SWU_CUSTOMER_EMAIL"),
		},
		EmailData: map[string]interface{}{
			"first_name": "John",
			"last_name":  "Doe",
		},
	}
	log, err := test.Send(email)
	assert.NoError(t, err)
	assert.True(t, log.Success)
}

func TestGetLog(t *testing.T) {
	logId := os.Getenv("SWU_LOG_ID")
	log, err := test.GetLog(logId)
	assert.NoError(t, err)
	assert.Equal(t, logId, log.Id)
}

func TestGetLogs(t *testing.T) {
	logs, err := test.GetLogs(&LogQuery{Count: 10})
	assert.NoError(t, err)
	assert.NotZero(t, len(logs))
}

func TestGetCustomerLogs(t *testing.T) {
	resp, err := test.GetCustomerLogs(os.Getenv("SWU_CUSTOMER_EMAIL"), &LogQuery{Count: 10})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestResendLog(t *testing.T) {
	log, err := test.ResendLog(os.Getenv("SWU_LOG_ID"))
	assert.NoError(t, err)
	assert.True(t, log.Success)
}

func TestGetCustomer(t *testing.T) {
	email := os.Getenv("SWU_CUSTOMER_EMAIL")
	resp, err := test.GetCustomer(email)
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	if err == nil {
		assert.Equal(t, email, resp.Customer.Email)
	}
}

func TestSaveCustomer(t *testing.T) {
	customer := &Customer{
		Email:  os.Getenv("SWU_CUSTOMER_EMAIL"),
		Locale: "pt-BR",
	}
	resp, err := test.SaveCustomer(customer)
	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestDeleteCustomer(t *testing.T) {
	resp, err := test.DeleteCustomer(os.Getenv("SWU_CUSTOMER_EMAIL"))
	assert.NoError(t, err)
	assert.True(t, resp.Success)
}
