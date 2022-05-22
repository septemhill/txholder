package application

import (
	storage "github.com/septemhill/txholder/storage/application"
)

type application struct {
	name    string
	appId   string
	version string
}

func (a *application) updateName(name string) {
	a.name = name
}

func (a *application) updateVersion(version string) {
	a.version = version
}

func (a *application) unmarshalToStorage() *storage.Application {
	return unmarshalToStorage(a)
}

func unmarshalToStorage(a *application) *storage.Application {
	return &storage.Application{
		Name:    a.name,
		AppId:   a.appId,
		Version: a.version,
	}
}

func marshalToModel(a *storage.Application) *application {
	return &application{
		name:    a.Name,
		appId:   a.AppId,
		version: a.Version,
	}
}
