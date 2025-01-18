package shard

import (
	"context"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const SRID = 4326

var (
	ErrInvalidValue = errors.New("GeoPoint: Scan received nil value")
	ErrInvalidType  = errors.New("GeoPoint: Scan expects string value")
	ErrDecodeEWKB   = errors.New("GeoPoint: failed to decode EWKB")
	ErrEncodeEWKB   = errors.New("GeoPoint: failed to encode EWKB")
)

// GeoPoint 封裝 PostGIS 的 Point 類型
type GeoPoint struct {
	geom.Point
}

// GormDataType 指定 GORM 對應的 PostGIS 資料型態
func (g GeoPoint) GormDataType() string {
	return "GEOGRAPHY(POINT, 4326)"
}

// GormValue 將 GeoPoint 轉為 SQL 查詢語句
func (g GeoPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_SetSRID(ST_GeomFromText(?), 4326)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", g.X(), g.Y())},
	}
}

// Scan 將資料庫中的值轉換為 GeoPoint
func (g *GeoPoint) Scan(val interface{}) error {
	if val == nil {
		return ErrInvalidValue
	}

	strVal, ok := val.(string)
	if !ok {
		return ErrInvalidType
	}

	pt, err := ewkbhex.Decode(strVal)
	if err != nil {
		return ErrDecodeEWKB
	}

	point, ok := pt.(*geom.Point)
	if !ok {
		return ErrDecodeEWKB
	}

	g.Point = *point
	return nil
}

func (g GeoPoint) Value() (driver.Value, error) {
	data, err := ewkbhex.Encode(g.SetSRID(SRID), binary.BigEndian)
	if err != nil {
		return nil, ErrEncodeEWKB
	}
	return data, nil
}
