package abstractfactory

type ForestFactory struct{}

func (ForestFactory) CreateEnemy() Enemy{
	return ForestGoblin{}
}
func (ForestFactory) CreateTerrain() Terrain{
	return ForestTerrain{}
}
func (ForestFactory) CreateSoundtrack() Soundtrack{
	return ForestSoundtrack{}
}