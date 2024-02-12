 # package database

This package contains code for managing a database connection using Go. It includes two tests: `TestNewController` and `TestController_Init`. The first test verifies the error handling when initializing the controller with an invalid path, while the second test checks the successful initialization of the controller with a valid path and executes the `Init` method to set up the database.

## Imports

The package imports the following packages:

* `fmt` for formatted I/O
* `os` for working with the operating system
* `testing` for writing tests
* `github.com/kilianp07/lifegame/config` for database configuration
* `github.com/stretchr/testify/assert` for testing assertions

## Tests

### TestNewController

This test checks the behavior of the controller when it is initialized with an invalid path and a valid path:

```go
func TestNewController(t *testing.T) {
	// ...
}
```

#### Testing with Invalid Path

The test initializes a new `Controller` instance with an invalid database path, expecting an error to be returned:

```go
conf := config.DBConf{
	DBPath: "../test/test.db",
}
_, err = NewController(conf)
if err == nil {
	t.Errorf("Expected error, got nil")
}
```

#### Testing with Valid Path

The test also initializes a new `Controller` instance with a valid database path:

```go
path, err := os.Getwd()
if err != nil {
	t.Errorf("Unable to get current path: %v", err)
}
// ...
conf = config.DBConf{
	DBPath: fmt.Sprintf("%s/test.db", path),
}
// Delete created database at the end of the test
defer os.Remove(conf.DBPath)
controller, err := NewController(conf)
assert.NoError(t, err)
assert.NotNil(t, controller)
assert.NotNil(t, controller.DB)
```

### TestController_Init

This test checks the successful initialization of the `Controller` instance with a valid database path:

```go
func TestController_Init(t *testing.T) {
	// ...
}
```

#### Creating a Temporary Test Database

The test creates a new temporary database file for testing purposes and sets up the controller configuration accordingly:

```go
tempDBPath := "../tests/temp.db"
defer os.Remove(tempDBPath)
conf := config.DBConf{
	DBPath: tempDBPath,
}
controller, err := NewController(conf)
assert.NoError(t, err)
assert.NotNil(t, controller)
```

#### Initializing the Controller

The test then calls the `Init` method of the `Controller` instance to initialize the database connection:

```go
err = controller.Init()
assert.NoError(t, err)
```

#### Verifying Migration Success (To be added)

You can add your logic for verifying that the migration was successful after initializing the controller in this test case.

#### Cleaning Up the Controller

Finally, the test cleans up the controller by calling its `Close` method:

```go
err = controller.Close()
assert.NoError(t, err)
```

