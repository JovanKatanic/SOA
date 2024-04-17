package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowerRepository struct {
	FollowerDriver *neo4j.DriverWithContext
}
