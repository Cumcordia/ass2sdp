package abstractfactory

type DesertFactory struct{}

func (DesertFactory) CreateEnemy() Enemy{
	return DesertSlime{}
}
func (DesertFactory) CreateSoundtrack() Soundtrack{
	return DesertSoundtrack{}
}
func (DesertFactory) CreateWorld() Terrain{
	return &DesertTerrain{}
}