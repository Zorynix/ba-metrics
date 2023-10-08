package services

import (
	"awesomeProject2/models"
	"context"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/mssola/useragent"
	"github.com/oschwald/geoip2-golang"
	"github.com/rs/zerolog/log"
)

func (pg *Postgres) TakeLink(id uuid.UUID) (models.Link, error) {
	rows, err := pg.db.Query(context.Background(), `
		SELECT id, url FROM ba_schema.links 
		WHERE id = $1`, id)
	if err != nil {
		return models.Link{}, err
	}

	link, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Link])
	if err != nil {
		return models.Link{}, err
	}

	return link, nil
}

func TakeFingerprints(c *fiber.Ctx) models.Fingerprints {

	ip := c.IP()
	userAgent := c.Get("User-Agent")
	referer := c.Get("Referer")

	ua := useragent.New(userAgent)

	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	geo_ip := net.ParseIP(ip)
	record, err := db.City(geo_ip)
	if err != nil {
		log.Fatal()
	}

	browser, _ := ua.Browser()
	localization := ua.Localization()
	model := ua.Model()
	platform := ua.Platform()
	os := ua.OS()
	city := record.City.Names["ru"]
	country := record.Country.Names["ru"]
	timezone := record.Location.TimeZone

	return models.Fingerprints{Ip: ip, City: city, Country: country, Timezone: timezone,
		Referer: referer, Browser: browser, Localization: localization,
		Model: model, Platform: platform, OS: os}
}

func (ch *Clickhouse) RecordMetrics(linkId uuid.UUID, fingerprints models.Fingerprints) error {

	err := ch.db.Exec(context.Background(), `
	INSERT INTO 
		ba_metrics.metrics (created_at, link_id, ip, city, country, timezone, referer, browser, localization, model, platform, os)
		VALUES
			(now(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`, linkId, fingerprints.Ip, fingerprints.City, fingerprints.Country, fingerprints.Timezone, fingerprints.Referer, fingerprints.Browser,
		fingerprints.Localization, fingerprints.Model, fingerprints.Platform, fingerprints.OS)
	if err != nil {
		return err
	}
	return nil
}
