package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/testdata"
)

func TestCreatePaymentAccount(t *testing.T) {

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	s, err := CreatePaymentAccount(svc)

	assert.NoError(t, err)
	assert.NotEmpty(t, *s)

}