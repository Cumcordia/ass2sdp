package abstractfactory

type ForestGoblin struct{}
type ForestTerrain struct{}
type ForestSoundtrack struct{}

func (ForestGoblin) Name() string {
	return "forest goblin"
}
func (ForestGoblin) Attack() int {
	return 20
}
func (ForestTerrain) Type() string {
	return "dense forest"
}
func (ForestSoundtrack) Track() string {
	return "forest_music.ogg"
}
