package main

import (
	"context"
	pb "github.com/farit2000/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Specification struct {
	Capacity  int32
	MaxWeight int32
}

func MarshalSpecification(spec *pb.Specification) *Specification{
	return &Specification{
		Capacity: spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalSpecification(spec *Specification) *pb.Specification{
	return &pb.Specification{
		Capacity: spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

type Vessel struct {
	ID        string
	Capacity  int32
	Name      string
	Available bool
	OwnerID   string
	MaxWeight int32
}

func MarshalVessel(ves *pb.Vessel) *Vessel {
	return &Vessel{
		ID: ves.Id,
		Capacity: ves.Capacity,
		MaxWeight: ves.MaxWeight,
		Name: ves.Name,
		Available: ves.Available,
		OwnerID: ves.OwnerId,
	}
}

func UnmarshalVessel(ves *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id: ves.ID,
		Capacity: ves.Capacity,
		MaxWeight: ves.MaxWeight,
		Name: ves.Name,
		Available: ves.Available,
		OwnerId: ves.OwnerID,
	}
}

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repository *MongoRepository) Create(ctx context.Context, ves *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, ves)
	return err
}
