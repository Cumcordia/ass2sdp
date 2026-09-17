package abstractfactory

type WorldFactory interface{
	CreateEnemy() Enemy
	CreateWorld() Terrain
	CreateSoundtrack() Soundtrack
}