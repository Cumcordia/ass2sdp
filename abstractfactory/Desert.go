package abstractfactory

type DesertSlime struct{}
type DesertTerrain struct{}
type DesertSoundtrack struct{}

func (DesertSlime) Name() string {
	return "desert slime"
}
func (DesertSlime) Attack() int {
	return 30
}
func (DesertTerrain) Type() string {
	return "dry deserts"
}
func (DesertSoundtrack) Track() string {
	return "desert_music.ogg"
}
