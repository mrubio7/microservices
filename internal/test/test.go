package testutil

import (
	"context"
	"fmt"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

var (
	testContainer testcontainers.Container
	dbConnections map[string]*gorm.DB
)

// SetupTestDBs initializes the PostgreSQL container and creates the necessary schemas.
func SetupTestDBs(schemas []string) {
	dbConnections = make(map[string]*gorm.DB)

	ctx := context.Background()

	// Initialize the PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	var err error
	testContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %v", err)
	}

	// Create connections for each schema
	for _, schema := range schemas {
		cfg := config.LoadTestDatabaseConfig(ctx, testContainer, schema)
		database := database.NewDatabase(cfg)

		if database == nil {
			log.Fatalf("Failed to initialize database for schema '%s'", schema)
		}

		dbConnections[schema] = database.GetDB()
	}
}

func CleanupTestDBs() {
	if testContainer != nil {
		ctx := context.Background()
		if err := testContainer.Terminate(ctx); err != nil {
			log.Printf("Failed to terminate test container: %v", err)
		}
	}
}

func GetTestDB(schema string) *gorm.DB {
	if db, exists := dbConnections[schema]; exists {
		clearTables(db, schema)
		return db
	}
	log.Fatalf("Schema '%s' not initialized", schema)
	return nil
}

func clearTables(db *gorm.DB, schema string) {
	var tableNames []string
	query := `
		SELECT tablename
		FROM pg_tables
		WHERE schemaname = $1;
	`
	if err := db.Raw(query, schema).Scan(&tableNames).Error; err != nil {
		log.Fatalf("Failed to fetch table names for schema '%s': %v", schema, err)
	}

	for _, tableName := range tableNames {
		truncateQuery := fmt.Sprintf("TRUNCATE TABLE %s.%s CASCADE;", schema, tableName)
		if err := db.Exec(truncateQuery).Error; err != nil {
			log.Fatalf("Failed to truncate table '%s.%s': %v", schema, tableName, err)
		}
	}
}
