package model

import (
	"strconv"
	"strings"
)

type App struct {
	Rank     int
	ID       string
	Name     string
	BundleID string
	Rights   string
	Artist   string
}

func (app *App) String() string {
	rank := strconv.Itoa(app.Rank)
	id := app.ID
	appName := app.Name
	artist := app.Artist

	return rank + ": " + appName + "(" + id + ")" + " [" + artist + "]"
}

func NewApp(content interface{}, rank int) *App {
	m := content.(map[string]interface{})

	return &App{
		Rank:   rank,
		Artist: m["artistName"].(string),
		Name:   m["name"].(string),
		ID:     m["id"].(string),
		Rights: m["copyright"].(string),
	}
}

func (app *App) Contains(keyword string) bool {
	key := strings.ToLower(keyword)
	artist := strings.ToLower(app.Artist)
	name := strings.ToLower(app.Name)
	bundleID := strings.ToLower(app.BundleID)
	rights := strings.ToLower(app.Rights)

	if strings.Contains(artist, key) {
		return true
	}

	if strings.Contains(name, key) {
		return true
	}

	if strings.Contains(bundleID, key) {
		return true
	}

	if strings.Contains(rights, key) {
		return true
	}

	return false
}
