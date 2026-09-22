package abstractfactory

type IceFactory struct{}

func (IceFactory) CreateEnemy() Enemy {
	return IceSkeleton{}
}
func (IceFactory) CreateSoundtrack() Soundtrack {
	return IceSoundtrack{}
}
func (IceFactory) CreateTerrain() Terrain {
	return IceTerrain{}
}
