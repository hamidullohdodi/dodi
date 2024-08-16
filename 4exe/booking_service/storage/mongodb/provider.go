package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p *BookingRepo) CreateProvider(ctx context.Context, req *pb.CreateProviderRequest) (*pb.CreateProviderResponse, error) {
	collection := p.db.Collection("provider")

	provider := bson.M{
		"user_id":      req.GetUserId(),
		"company_name": req.GetCompanyName(),
		"service_id":   req.GetServiceIds(),
		"location": bson.M{
			"city":    req.GetLocation().GetCity(),
			"country": req.GetLocation().GetCountry(),
		},
	}
	resp, err := collection.InsertOne(ctx, provider)
	if err != nil {
		return nil, err
	}
	oid, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert objectid to hex: %v", resp.InsertedID)
	}
	return &pb.CreateProviderResponse{
		UserId:       req.GetUserId(),
		CompanyName:  req.GetCompanyName(),
		ServiceIds:   req.GetServiceIds(),
		Id:           oid.Hex(),
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		Availability: req.GetAvailability(),
		Location: &pb.Location{
			City:    req.GetLocation().GetCity(),
			Country: req.GetLocation().GetCountry(),
		},
	}, nil
}

func (p *BookingRepo) GetProvider(ctx context.Context, req *pb.GetProviderRequest) (*pb.GetProviderResponse, error) {
	// Check if the provided ID is a valid ObjectID
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid provider ID: %s", req.Id)
	}

	collection := p.db.Collection("provider")

	// MongoDB filter
	filter := bson.M{"_id": oid}

	var provider struct {
		UserId      string        `bson:"user_id"`
		CompanyName string        `bson:"company_name"`
		ServiceIds  []interface{} `bson:"service_id"`
		Location    struct {
			City    string `bson:"city"`
			Country string `bson:"country"`
		} `bson:"location"`
		Availability []struct {
			StartTime string `bson:"start_time"`
			EndTime   string `bson:"end_time"`
		} `bson:"availability"`
		Id        primitive.ObjectID `bson:"_id"`
		CreatedAt string             `bson:"created_at"`
	}

	err = collection.FindOne(ctx, filter).Decode(&provider)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("provider not found")
		}
		return nil, err
	}

	// Prepare response
	availability := make([]*pb.Availability, len(provider.Availability))
	for i, avail := range provider.Availability {
		availability[i] = &pb.Availability{
			StartTime: avail.StartTime,
			EndTime:   avail.EndTime,
		}
	}

	// Converting service IDs to string
	serviceIds := make([]*pb.ServiceId, len(provider.ServiceIds))
	for i, sid := range provider.ServiceIds {
		switch v := sid.(type) {
		case primitive.ObjectID:
			serviceIds[i] = &pb.ServiceId{
				Id: v.Hex(),
			}
		case string:
			serviceIds[i] = &pb.ServiceId{
				Id: v,
			}
		case primitive.D:
			for _, elem := range v {
				if elem.Key == "id" {
					if id, ok := elem.Value.(string); ok {
						serviceIds[i] = &pb.ServiceId{
							Id: id,
						}
					} else if oid, ok := elem.Value.(primitive.ObjectID); ok {
						serviceIds[i] = &pb.ServiceId{
							Id: oid.Hex(),
						}
					} else {
						return nil, fmt.Errorf("unexpected type for service_id: %T", elem.Value)
					}
				}
			}
		default:
			return nil, fmt.Errorf("unexpected type for service_id: %T", sid)
		}
	}

	return &pb.GetProviderResponse{
		UserId:       provider.UserId,
		CompanyName:  provider.CompanyName,
		ServiceIds:   serviceIds,
		Id:           provider.Id.Hex(),
		CreatedAt:    provider.CreatedAt,
		Availability: availability,
		Location: &pb.Location{
			City:    provider.Location.City,
			Country: provider.Location.Country,
		},
	}, nil
}


func (p *BookingRepo) ListProviders(ctx context.Context, req *pb.ListProvidersRequest) (*pb.ListProvidersResponse, error) {
    collection := p.db.Collection("provider")
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var providers []*pb.Provider
    for cursor.Next(ctx) {
        var provider struct {
            Id          primitive.ObjectID `bson:"_id"`
            UserId      string             `bson:"user_id"`
            CompanyName string             `bson:"company_name"`
            ServiceIds  []struct {
                Id string `bson:"id"`
            } `bson:"service_id"` // This now matches your MongoDB schema
            Location struct {
                City    string `bson:"city"`
                Country string `bson:"country"`
            } `bson:"location"`
            Availability []struct {
                StartTime string `bson:"start_time"`
                EndTime   string `bson:"end_time"`
            } `bson:"availability"`
        }
        err := cursor.Decode(&provider)
        if err != nil {
            return nil, err
        }

        // Convert service IDs
        serviceIds := make([]*pb.ServiceId, len(provider.ServiceIds))
        for i, sid := range provider.ServiceIds {
            serviceIds[i] = &pb.ServiceId{Id: sid.Id}
        }

        // Convert availability
        availability := make([]*pb.Availability, len(provider.Availability))
        for i, avail := range provider.Availability {
            availability[i] = &pb.Availability{
                StartTime: avail.StartTime,
                EndTime:   avail.EndTime,
            }
        }

        // Add provider to the list
        providers = append(providers, &pb.Provider{
            Id:           provider.Id.Hex(),
            UserId:       provider.UserId,
            CompanyName:  provider.CompanyName,
            ServiceIds:   serviceIds,
            Availability: availability,
            Location: &pb.Location{
                City:    provider.Location.City,
                Country: provider.Location.Country,
            },
        })
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return &pb.ListProvidersResponse{
        Providers: providers,
    }, nil
}



func (p *BookingRepo) UpdateProvider(ctx context.Context, req *pb.UpdateProviderRequest) (*pb.UpdateProviderResponse, error) {
	collection := p.db.Collection("provider")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"user_id":      req.GetUserId(),
			"company_name": req.GetCompanyName(),
			"id":           req.GetId(),
			"location": bson.M{
				"city":    req.GetLocation().GetCity(),
				"country": req.GetLocation().GetCountry(),
			},
		},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProviderResponse{
		UserId:      req.GetUserId(),
		CompanyName: req.GetCompanyName(),
		Id:          req.GetId(),
		Location: &pb.Location{
			City:    req.GetLocation().GetCity(),
			Country: req.GetLocation().GetCountry(),
		},
	}, nil
}

func (p *BookingRepo) DeleteProvider(ctx context.Context, req *pb.DeleteProviderRequest) (*pb.DeleteProviderResponse, error) {
	collection := p.db.Collection("provider")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProviderResponse{
		Message: "Provider deleted successfully",
	}, nil
}

func (p *BookingRepo) SearchProviders(ctx context.Context, req *pb.SearchProvidersRequest) (*pb.SearchProvidersResponse, error) {
	collection := p.db.Collection("provider")

	filter := bson.M{}
	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}

	if req.CompanyName != "" {
		filter["company_name"] = req.CompanyName
	}

	if req.Location != nil {
		if req.Location.City != "" || req.Location.Country != "" {
			locationFilter := bson.M{}
			if req.Location.City != "" {
				locationFilter["city"] = req.Location.City
			}
			if req.Location.Country != "" {
				locationFilter["country"] = req.Location.Country
			}
			filter["location"] = locationFilter
		}
	}

	// Log the filter to debug
	fmt.Printf("Search filter: %+v\n", filter)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var providers []*pb.Provider
	for cursor.Next(ctx) {
		var provider struct {
			UserId      string   `bson:"user_id"`
			CompanyName string   `bson:"company_name"`
			ServiceIds  []string `bson:"service_id"`
			Location    struct {
				City    string `bson:"city"`
				Country string `bson:"country"`
			} `bson:"location"`
			Id           primitive.ObjectID `bson:"_id"`
			CreatedAt    string             `bson:"created_at"`
			Availability []struct {
				StartTime string `bson:"start_time"`
				EndTime   string `bson:"end_time"`
			} `bson:"availability"`
		}
		if err := cursor.Decode(&provider); err != nil {
			return nil, err
		}

		// Convert availability
		var availability []*pb.Availability
		for _, avail := range provider.Availability {
			availability = append(availability, &pb.Availability{
				StartTime: avail.StartTime,
				EndTime:   avail.EndTime,
			})
		}

		providers = append(providers, &pb.Provider{
			UserId:       provider.UserId,
			CompanyName:  provider.CompanyName,
			ServiceIds:   convertServiceIds(provider.ServiceIds),
			Id:           provider.Id.Hex(),
			Availability: availability,
			Location: &pb.Location{
				City:    provider.Location.City,
				Country: provider.Location.Country,
			},
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Log the result to debug
	fmt.Printf("Search result: %+v\n", providers)

	return &pb.SearchProvidersResponse{
		Providers: providers,
	}, nil
}

func convertServiceIds(serviceIds []string) []*pb.ServiceId {
	converted := make([]*pb.ServiceId, len(serviceIds))
	for i, id := range serviceIds {
		converted[i] = &pb.ServiceId{Id: id}
	}
	return converted
}
