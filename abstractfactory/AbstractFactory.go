package abstractfactory

type WorldFactory interface{
	CreateEnemy() Enemy
	CreateTerrain() Terrain
	CreateSoundtrack() Soundtrack
}