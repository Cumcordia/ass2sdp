package abstractfactory

import (
	"errors"
	"fmt"
)

type WorldLoader struct{
	factory WorldFactory
}

func NewWorldLoader(factory WorldFactory) *WorldLoader{
	return &WorldLoader{factory: factory}
}

func (w *WorldLoader) LoadWorld() string{
	enemy := w.factory.CreateEnemy()
	terrain := w.factory.CreateTerrain()
	soundtrack := w.factory.CreateSoundtrack()

	return fmt.Sprint(terrain.Type(), enemy.Name(), enemy.Attack(), soundtrack.Track())
}

func GetWorldFactory(world string) (WorldFactory, error) {
	switch world {
	case "forest":
		return ForestFactory{}, nil
	case "desert":
		return DesertFactory{}, nil
	case "ice":
		return IceFactory{}, nil
	default:
		return nil, errors.New("no world selected")
	}
}