package top

import (
	"strconv"
	"strings"
)

type App struct {
	Content interface{}
	Rank    int
}

func NewApp(content interface{}, rank int) *App {
	app := &App{}
	app.Content = content
	app.Rank = rank
	return app
}

func (app *App) String() string {
	rank := strconv.Itoa(app.Rank)
	bundleID := app.BundleID()
	artist := app.ArtistName()
	rights := app.Rights()

	return rank + "(" + bundleID + ")" + " [" + artist + "/" + rights + "]"
}

func (app *App) ArtistName() string {
	artist := app.Content.(map[string]interface{})["im:artist"]
	label := artist.(map[string]interface{})["label"]
	return label.(string)
}

func (app *App) AppName() string {
	name := app.Content.(map[string]interface{})["im:name"]
	label := name.(map[string]interface{})["label"]
	return label.(string)
}

func (app *App) BundleID() string {
	id := app.Content.(map[string]interface{})["id"]
	attributes := id.(map[string]interface{})["attributes"]
	bundleID := attributes.(map[string]interface{})["im:bundleId"]
	return bundleID.(string)
}

func (app *App) Rights() string {
	rights := app.Content.(map[string]interface{})["rights"]
	label := rights.(map[string]interface{})["label"]
	return label.(string)
}

func (app *App) Contains(keyword string) bool {
	key := strings.ToLower(keyword)
	artist := strings.ToLower(app.ArtistName())
	appName := strings.ToLower(app.AppName())
	bundleID := strings.ToLower(app.BundleID())
	rights := strings.ToLower(app.Rights())

	if strings.Contains(artist, key) {
		return true
	}

	if strings.Contains(appName, key) {
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
