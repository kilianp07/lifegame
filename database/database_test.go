package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/kilianp07/lifegame/config"
	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	var (
		err        error
		controller = &Controller{}
		path       string
	)

	// Test with invalid path
	conf := config.DBConf{
		DBPath: "../test/test.db",
	}
	_, err = NewController(conf)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	path, err = os.Getwd()
	if err != nil {
		t.Errorf("Unable to get current path: %v", err)
	}
	// Test with valid path
	conf = config.DBConf{
		DBPath: fmt.Sprintf("%s/test.db", path),
	}

	// Delete created database at the end of the test
	defer os.Remove(conf.DBPath)

	controller, err = NewController(conf)
	assert.NoError(t, err)
	assert.NotNil(t, controller)
	assert.NotNil(t, controller.DB)
}

func TestController_Init(t *testing.T) {
	// Create a temporary test database
	tempDBPath := "../tests/temp.db"
	defer os.Remove(tempDBPath)

	// Create a new controller
	conf := config.DBConf{
		DBPath: tempDBPath,
	}
	controller, err := NewController(conf)
	assert.NoError(t, err)
	assert.NotNil(t, controller)

	// Initialize the controller
	err = controller.Init()
	assert.NoError(t, err)

	// Verify that the migration was successful
	// Add your verification logic here

	// Clean up the controller
	err = controller.Close()
	assert.NoError(t, err)
}
