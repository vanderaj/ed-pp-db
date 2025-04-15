# ed-pp-db

Create a CSV file of systems controlled by a Power, with exploited, fortified, and stronghold systems for the back office spreadsheet.

## Loading the data

Download the galaxy_stations.json.gz file from Spansh:

- [https://downloads.spansh.co.uk/galaxy_stations.json.gz](https://downloads.spansh.co.uk/galaxy_stations.json.gz)

It's currently 2 GB in size.

## Decompress the file

Either use the desktop or gzip to decompress the file

```bash
gzip -d galaxy_stations.json.gz
```

This will produce a 12 GB file.

## Data Load

Load into MongoDB using Mongo Compass or MongoDB shell. This will take quite a while.

```bash
mongoimport --db galaxyStations --collection galaxyStations --file galaxy_stations.json --jsonArray
```

## Create Index

You'll definitely want to create indexes on the collection to speed up queries.

```bash
use galaxyStations
db.galaxyStations.createIndex({ "controllingPower": 1 }, { "powerState": 1 })
```

## Build the app

```bash
go build
```

## Run the app

```bash
./ed-pp-db -power "Pranav Antal" -output output.csv
```

Right now the arguments don't work, but will in a future release.

The output will be a CSV file with the following columns:
- systemName: The name of the system
- powerState: The state of the power (exploited, fortified, stronghold)
- PowerStateControlProgress: the reinforced percentage for the system (0-100%)
- Strategy: A default reinforcement strategy for the system
- Notes: If an anarchy, default text around bounty hunting in an anarchy system
- Stronghold: The nearest stronghold or fortified system (fortified only if no stronghold within range)
- population: The population of the system
- economies: The economies of the system
- Mining: The mining level of the system (N/A through Pristine)
- L-pad: If there is a space L-pad to land on
- Resource location: if there are any noted hotspots
- PMF: the controlling PMF
- distance: The distance to the nearest stronghold or fortified system (fortified only if no stronghold within range)
