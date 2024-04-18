package repository

import (
	"context"
	"log"
	"os"
	"user_management_service/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowerRepository struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

func NewFollowerRepository(logger *log.Logger) (*FollowerRepository, error) {
	uri := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USERNAME")
	pass := os.Getenv("NEO4J_PASS")
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	return &FollowerRepository{
		driver: driver,
		logger: logger,
	}, nil
}

func (f *FollowerRepository) CheckConnection() {
	ctx := context.Background()
	err := f.driver.VerifyConnectivity(ctx)
	if err != nil {
		f.logger.Panic(err)
		return
	}

	f.logger.Printf(`Neo4J server address: %s`, f.driver.Target().Host)
}

func (f *FollowerRepository) CloseDriverConnection(ctx context.Context) {
	f.driver.Close(ctx)
}

func (f *FollowerRepository) WriteFollower(follower *model.Follower) error {
	ctx := context.Background()
	session := f.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	f.logger.Print("")
	f.logger.Print("Repository check:")
	f.logger.Print(follower.FollowerId, follower.FollowedId)
	f.logger.Print("")

	savedFollowing, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (a:Person), (b:Person) WHERE a.id = $aId  AND b.id = $bId CREATE (a) -[r:FOLLOWS]-> (b) RETURN type(r)",
				map[string]any{"aId": follower.FollowerId, "bId": follower.FollowedId})

			f.logger.Print("Napravljen je upit")
			if err != nil {
				f.logger.Print("Error ovde nastao")
				return nil, err
			}

			if result.Next(ctx) {
				f.logger.Print("Uspesno je upisano")
				return result.Record().Values[0], nil
			}

			f.logger.Print("Rezultat ima neki err")
			f.logger.Print(result.Err().Error())
			f.logger.Print("Rezultat(deo bez errora)")
			f.logger.Print(result.Record().Values[0])
			return nil, result.Err()
		})

	if err != nil {
		f.logger.Println("Error inserting follow: ", err)
		return err
	}

	f.logger.Println(savedFollowing.(string))
	return nil
}
