package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Define a struct to represent the MongoDB document structure
type System []struct {
	ID struct {
		Oid string `bson:"$oid"`
	} `bson:"_id"`
	ID64 struct {
		NumberLong string `bson:"$numberLong"`
	} `bson:"id64"`
	Name   string `bson:"name"`
	Coords struct {
		X float64 `bson:"x"`
		Y float64 `bson:"y"`
		Z float64 `bson:"z"`
	} `bson:"coords"`
	Allegiance         string `bson:"allegiance"`
	Government         string `bson:"government"`
	PrimaryEconomy     string `bson:"primaryEconomy"`
	SecondaryEconomy   string `bson:"secondaryEconomy"`
	Security           string `bson:"security"`
	Population         int    `bson:"population"`
	BodyCount          int    `bson:"bodyCount"`
	ControllingFaction struct {
		Name       string `bson:"name"`
		Government string `bson:"government"`
		Allegiance string `bson:"allegiance"`
	} `bson:"controllingFaction"`
	Factions []struct {
		Name       string  `bson:"name"`
		Allegiance string  `bson:"allegiance"`
		Government string  `bson:"government"`
		Influence  float64 `bson:"influence"`
		State      string  `bson:"state"`
	} `bson:"factions"`
	ControllingPower          string   `bson:"controllingPower"`
	PowerState                string   `bson:"powerState"`
	PowerStateControlProgress float64  `bson:"powerStateControlProgress"`
	PowerStateReinforcement   int      `bson:"powerStateReinforcement"`
	PowerStateUndermining     int      `bson:"powerStateUndermining"`
	Powers                    []string `bson:"powers"`
	Timestamps                struct {
		ControllingPower time.Time `bson:"controllingPower"`
		PowerState       time.Time `bson:"powerState"`
		Powers           time.Time `bson:"powers"`
	} `bson:"timestamps"`
	Date   string `bson:"date"`
	Bodies []struct {
		ID64 struct {
			NumberLong string `bson:"$numberLong"`
		} `bson:"id64"`
		BodyID              int     `bson:"bodyId"`
		Name                string  `bson:"name"`
		Type                string  `bson:"type"`
		OrbitalPeriod       float64 `bson:"orbitalPeriod"`
		SemiMajorAxis       float64 `bson:"semiMajorAxis"`
		OrbitalEccentricity float64 `bson:"orbitalEccentricity"`
		OrbitalInclination  float64 `bson:"orbitalInclination"`
		ArgOfPeriapsis      float64 `bson:"argOfPeriapsis"`
		MeanAnomaly         float64 `bson:"meanAnomaly"`
		AscendingNode       float64 `bson:"ascendingNode"`
		Timestamps          struct {
			MeanAnomaly time.Time `bson:"meanAnomaly"`
		} `bson:"timestamps"`
		Stations                      []interface{} `bson:"stations"`
		UpdateTime                    string        `bson:"updateTime"`
		SubType                       string        `bson:"subType,omitempty"`
		DistanceToArrival             int           `bson:"distanceToArrival,omitempty"`
		MainStar                      bool          `bson:"mainStar,omitempty"`
		Age                           int           `bson:"age,omitempty"`
		SpectralClass                 string        `bson:"spectralClass,omitempty"`
		Luminosity                    string        `bson:"luminosity,omitempty"`
		AbsoluteMagnitude             float64       `bson:"absoluteMagnitude,omitempty"`
		SolarMasses                   float64       `bson:"solarMasses,omitempty"`
		SolarRadius                   float64       `bson:"solarRadius,omitempty"`
		SurfaceTemperature            int           `bson:"surfaceTemperature,omitempty"`
		RotationalPeriod              float64       `bson:"rotationalPeriod,omitempty"`
		RotationalPeriodTidallyLocked bool          `bson:"rotationalPeriodTidallyLocked,omitempty"`
		AxialTilt                     int           `bson:"axialTilt,omitempty"`
		Parents                       []struct {
			Null int `bson:"Null"`
		} `bson:"parents,omitempty"`
		Belts []struct {
			Name string `bson:"name"`
			Type string `bson:"type"`
			Mass struct {
				NumberLong string `bson:"$numberLong"`
			} `bson:"mass"`
			InnerRadius int `bson:"innerRadius"`
			OuterRadius int `bson:"outerRadius"`
		} `bson:"belts,omitempty"`
		IsLandable       bool        `bson:"isLandable,omitempty"`
		Gravity          float64     `bson:"gravity,omitempty"`
		EarthMasses      float64     `bson:"earthMasses,omitempty"`
		Radius           float64     `bson:"radius,omitempty"`
		SurfacePressure  int         `bson:"surfacePressure,omitempty"`
		VolcanismType    string      `bson:"volcanismType,omitempty"`
		AtmosphereType   interface{} `bson:"atmosphereType,omitempty"`
		SolidComposition struct {
			Ice   int `bson:"Ice"`
			Metal int `bson:"Metal"`
			Rock  int `bson:"Rock"`
		} `bson:"solidComposition,omitempty"`
		TerraformingState string `bson:"terraformingState,omitempty"`
		Materials         struct {
			Cadmium    float64 `bson:"Cadmium"`
			Chromium   float64 `bson:"Chromium"`
			Iron       float64 `bson:"Iron"`
			Manganese  float64 `bson:"Manganese"`
			Mercury    float64 `bson:"Mercury"`
			Nickel     float64 `bson:"Nickel"`
			Technetium float64 `bson:"Technetium"`
			Zirconium  float64 `bson:"Zirconium"`
		} `bson:"materials,omitempty"`
		Signals struct {
			Signals struct {
				SAASignalTypeHuman int `bson:"$SAA_SignalType_Human;"`
			} `bson:"signals"`
			Genuses    []interface{} `bson:"genuses"`
			UpdateTime string        `bson:"updateTime"`
		} `bson:"signals,omitempty"`
		AtmosphereComposition struct {
			SulphurDioxide int `bson:"Sulphur dioxide"`
		} `bson:"atmosphereComposition,omitempty"`
		ReserveLevel string `bson:"reserveLevel,omitempty"`
		Rings        []struct {
			Name string `bson:"name"`
			Type string `bson:"type"`
			Mass struct {
				NumberLong string `bson:"$numberLong"`
			} `bson:"mass"`
			InnerRadius int `bson:"innerRadius"`
			OuterRadius int `bson:"outerRadius"`
			ID64        struct {
				NumberLong string `bson:"$numberLong"`
			} `bson:"id64"`
			Signals struct {
				Signals struct {
					Monazite     int `bson:"Monazite"`
					Painite      int `bson:"Painite"`
					Platinum     int `bson:"Platinum"`
					Rhodplumsite int `bson:"Rhodplumsite"`
					Serendibite  int `bson:"Serendibite"`
				} `bson:"signals"`
				UpdateTime string `bson:"updateTime"`
			} `bson:"signals,omitempty"`
		} `bson:"rings,omitempty"`
	} `bson:"bodies"`
	Stations []struct {
		Name string `bson:"name"`
		ID   struct {
			NumberLong string `bson:"$numberLong"`
		} `bson:"id"`
		UpdateTime              string      `bson:"updateTime"`
		ControllingFaction      string      `bson:"controllingFaction"`
		ControllingFactionState interface{} `bson:"controllingFactionState"`
		DistanceToArrival       float64     `bson:"distanceToArrival"`
		PrimaryEconomy          string      `bson:"primaryEconomy"`
		Economies               struct {
			Industrial int `bson:"Industrial"`
			Refinery   int `bson:"Refinery"`
		} `bson:"economies"`
		Government  string   `bson:"government"`
		Services    []string `bson:"services"`
		Type        string   `bson:"type"`
		LandingPads struct {
			Large  int `bson:"large"`
			Medium int `bson:"medium"`
			Small  int `bson:"small"`
		} `bson:"landingPads"`
		Market struct {
			Commodities []struct {
				Name        string `bson:"name"`
				Symbol      string `bson:"symbol"`
				Category    string `bson:"category"`
				CommodityID int    `bson:"commodityId"`
				Demand      int    `bson:"demand"`
				Supply      int    `bson:"supply"`
				BuyPrice    int    `bson:"buyPrice"`
				SellPrice   int    `bson:"sellPrice"`
			} `bson:"commodities"`
			ProhibitedCommodities []interface{} `bson:"prohibitedCommodities"`
			UpdateTime            string        `bson:"updateTime"`
		} `bson:"market"`
		Shipyard struct {
			Ships []struct {
				Name   string `bson:"name"`
				Symbol string `bson:"symbol"`
				ShipID int    `bson:"shipId"`
			} `bson:"ships"`
			UpdateTime string `bson:"updateTime"`
		} `bson:"shipyard,omitempty"`
		Outfitting struct {
			Modules []struct {
				Name     string `bson:"name"`
				Symbol   string `bson:"symbol"`
				ModuleID int    `bson:"moduleId"`
				Class    int    `bson:"class"`
				Rating   string `bson:"rating"`
				Category string `bson:"category"`
				Ship     string `bson:"ship,omitempty"`
			} `bson:"modules"`
			UpdateTime string `bson:"updateTime"`
		} `bson:"outfitting,omitempty"`
	} `bson:"stations"`
}

func main() {
	fmt.Println("Connecting to MongoDB!")

	// MongoDB connection URI
	uri := "mongodb://localhost:27017"
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Access the collection
	collection := client.Database("galaxyStations").Collection("galaxyStations")

	// Define the filter for querying documents (optional)
	filter := bson.D{
		{"controllingPower", "Pranav Antal"},
		{"powerState",
			bson.D{
				{"$in",
					bson.A{
						"Exploited",
						"Fortified",
						"Stronghold",
					},
				},
			},
		},
	}
	// Execute the query
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Unmarshal the documents into a slice of structs
	var systems []System
	for cursor.Next(context.TODO()) {
		var system System
		if err := cursor.Decode(&system); err != nil {
			log.Printf("Error decoding document: %v", err)
			continue
		}
		systems = append(systems, system)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}

	// Print the results
	fmt.Println("Retrieved documents:")
	for _, system := range systems {
		fmt.Printf("Name: %s, Population: %d, Economies: %s\n", system[0].Name, system[0].Population, system[0].PrimaryEconomy)
	}

	// Look up the record in Mongo DB

	// System name, population, economies

	// Nearest stronghold system (if any) and distance

	// Any bodies with mining - what level is it?

	// L pad stations (M pad only, Planetary, Yes)

	// Megaship strategy

	// Industrial HVT/donations strategy

	// Bounty hunting (installation || megaship || rings) strategy

	// Anarchy data raid strategy

	// Power USS farming strategy

	// Power Commodities strategy

	// Append system to CSV file
}
