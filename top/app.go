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

func (app *App) ToString() string {
	rank := strconv.Itoa(app.Rank)
	bundleID := app.GetBundleID()
	artist := app.GetArtistName()
	rights := app.GetRights()

	return rank + "(" + bundleID + ")" + " [" + artist + "/" + rights + "]"
}

func (app *App) GetArtistName() string {
	artist := app.Content.(map[string]interface{})["im:artist"]
	label := artist.(map[string]interface{})["label"]
	return label.(string)
}

func (app *App) GetAppName() string {
	name := app.Content.(map[string]interface{})["im:name"]
	label := name.(map[string]interface{})["label"]
	return label.(string)
}

func (app *App) GetBundleID() string {
	id := app.Content.(map[string]interface{})["id"]
	attributes := id.(map[string]interface{})["attributes"]
	bundleID := attributes.(map[string]interface{})["im:bundleId"]
	return bundleID.(string)
}

func (app *App) GetRights() string {
	rights := app.Content.(map[string]interface{})["rights"]
	label := rights.(map[string]interface{})["label"]
	return label.(string)
}

func (app *App) Contains(keyword string) bool {
	key := strings.ToLower(keyword)
	artist := strings.ToLower(app.GetArtistName())
	appName := strings.ToLower(app.GetAppName())
	bundleID := strings.ToLower(app.GetBundleID())
	rights := strings.ToLower(app.GetRights())

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
