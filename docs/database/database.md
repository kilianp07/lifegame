 # database Package Documentation

The `database` package is responsible for handling all database-related functionalities in the application. It utilizes the GORM ORM framework and SQLite driver to interact with the database.

## Import Statements

```go
import (
	"github.com/kilianp07/lifegame/config"
	"github.com/kilianp07/lifegame/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
```

The package imports the required configurations, models, and GORM packages for database interactions.

## Controller Struct

```go
type Controller struct {
	DB *gorm.DB
}
```

The `Controller` struct holds a pointer to the `gorm.DB` instance. This instance is used to interact with the database.

## Functions

### NewController

```go
func NewController(conf config.DBConf) (*Controller, error) {
	// ...
}
```

The `NewController` function initializes and returns a new instance of the `Controller` struct with a connection to the database. It takes a configuration object as an argument containing the database path. This function creates the database connection using the provided config and initializes GORM with the given DBPath.

### Init

```go
func (c *Controller) Init() error {
	// ...
}
```

The `Init` method performs any necessary database initialization, such as migrations, by calling the `migrate` method.

### migrate

```go
func (c *Controller) migrate() error {
	// ...
}
```

The `migrate` method utilizes GORM's AutoMigrate functionality to create or update the database schema based on the model definitions.

### CreateUser

```go
func (c *Controller) CreateUser(user *model.User) error {
	// ...
}
```

The `CreateUser` method creates a new user record in the database by using the `gorm.DB` instance to create the `model.User` struct.

### Close

```go
func (c *Controller) Close() error {
	// ...
}
```

The `Close` method gracefully closes the database connection and returns any errors that might have occurred during the closure process.

