package geoip

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/go-redis/redis/v8"
)

const ASNBlocksDataCSVPath string = "data/GeoLite2-ASN-CSV_20210420/GeoLite2-ASN-Blocks-IPv4.csv"
const CityBlocksDataCSVPath string = "data/GeoLite2-City-CSV_20210413/GeoLite2-City-Blocks-IPv4.csv"
const CityLocationDataCSVPath string = "data/GeoLite2-City-CSV_20210413/GeoLite2-City-Locations-en.csv"
const CountryLocationDataCSVPath string = "data/GeoLite2-Country-CSV_20210413/GeoLite2-Country-Locations-en.csv"

func (geo *Client) ImportGeoIPData(redisClient *redis.Client) error {
	var err error
	importMap := map[string]func([]string) interface{}{
		ASNBlocksDataCSVPath:       NewASNBlock,
		CityLocationDataCSVPath:    NewCityLocation,
		CountryLocationDataCSVPath: NewCountryLocation,
		CityBlocksDataCSVPath:      NewCityBlock,
	}
	for path, f := range importMap {
		err = geo.importGeoIPData(f, path, redisClient)
		if err != nil {
			geo.log.Error("Failed to load GeoIP export data: ", err)
			return err
		}
	}
	return err
}

func (geo *Client) importGeoIPData(deserial func([]string) interface{}, path string, redisClient *redis.Client) error {
	var err error
	dump, err := os.Open(path)
	if err != nil {
		geo.log.Error("Failed to load GeoIP data: ", err)
		return err
	}
	defer dump.Close()
	reader := bufio.NewReader(dump)
	csvReader := csv.NewReader(reader)
	numRecords := 0
	ctx := context.TODO()
	// read headers first
	_, err = csvReader.Read()
	if err != nil {
		geo.log.Error("Failed to load GeoIP data as csv: ", err)
		return err
	}
	pipe := redisClient.Pipeline()
	for {
		var row []string
		row, err = csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			geo.log.Error("Error reading GeoIP data: ", err)
			break
		}
		obj := deserial(row)
		geo.deserializeData(ctx, pipe, obj)
		if numRecords%1000 == 0 {
			_, err = pipe.Exec(ctx)
			if err != nil {
				geo.log.Error("Failed to commit GeoIP data: ", err)
			}
		}
		numRecords++
		if err != nil {
			geo.log.Error(fmt.Sprintf("[%s] Failed to import GeoIP data", path), err)
		}
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		geo.log.Error("Failed to commit GeoIP data: ", err)
	}
	geo.log.Infof("[%s] Loaded %d GeoIP records", path, numRecords)
	dump.Close()
	return nil
}

func (geo *Client) deserializeData(ctx context.Context, pipe redis.Pipeliner, obj interface{}) {
	switch v := obj.(type) {
	case *ASNBlock:
		asnBlock := obj.(*ASNBlock)
		asnBlock.ToHSet(ctx, pipe)
	case *CityBlock:
		cityBlock := obj.(*CityBlock)
		cityBlock.ToHSet(ctx, pipe)
	case *CityLocation:
		cityLocation := obj.(*CityLocation)
		cityLocation.ToHSet(ctx, pipe)
	case *CountryBlock:
		countryBlock := obj.(*CountryBlock)
		countryBlock.ToHSet(ctx, pipe)
	case *CountryLocation:
		countryLocation := obj.(*CountryLocation)
		countryLocation.ToHSet(ctx, pipe)
	default:
		geo.log.Errorf("Failed to cast deserializer: %T: ", v)
	}
}
