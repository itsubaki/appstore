package ranking

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

func NewApp(content interface{}, rank int) *App {
	app := &App{Rank: rank}

	artist := content.(map[string]interface{})["im:artist"]
	artistlabel := artist.(map[string]interface{})["label"]
	app.Artist = artistlabel.(string)

	name := content.(map[string]interface{})["im:name"]
	namelabel := name.(map[string]interface{})["label"]
	app.Name = namelabel.(string)

	id := content.(map[string]interface{})["id"]
	attributes := id.(map[string]interface{})["attributes"]
	bundleID := attributes.(map[string]interface{})["im:bundleId"]
	app.BundleID = bundleID.(string)

	rights := content.(map[string]interface{})["rights"]
	rightslabel := rights.(map[string]interface{})["label"]
	app.Rights = rightslabel.(string)

	return app
}

func (app *App) String() string {
	rank := strconv.Itoa(app.Rank)
	appName := app.Name
	bundleID := app.BundleID
	artist := app.Artist

	return rank + ": " + appName + "(" + bundleID + ")" + " [" + artist + "]"
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
