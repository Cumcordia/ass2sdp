package abstractfactory

type IceSkeleton struct{}
type IceTerrain struct{}
type IceSoundtrack struct{}

func (IceSkeleton) Name() string {
	return "ice skeleton"
}
func (IceSkeleton) Attack() int {
	return 25
}
func (IceTerrain) Type() string {
	return "frozen ice world"
}
func (IceSoundtrack) Track() string {
	return "frozen_music.ogg"
}
