//go:build integration_test

package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KScaesar/jubo-homework/backend/configs"
	"github.com/KScaesar/jubo-homework/backend/util/database"
)

func TestNewGormPgsql(t *testing.T) {
	config, err := configs.NewProjectConfig()
	assert.NoError(t, err)

	_, err = database.NewGormPgsql(config.Pgsql)
	assert.NoError(t, err)
}
