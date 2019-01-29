package gotest

import "testing"
import "sms/models"

func TestGetDB(t *testing.T) {
	_, err := models.GetDB("CPMS")
	if err != nil {
		t.Error(err)
	}
}
