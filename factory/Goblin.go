package factory

type Goblin struct{}

type GoblinCreator struct{
	EnemyCreator
}

func NewGoblinCreator() *GoblinCreator{
	g := &GoblinCreator{}
	g.self = g
	return g
}

func (g *GoblinCreator) createEnemy() Enemy{
	return &Goblin{}
}

func (g *Goblin) Name() string {
	return "Goblin"
}
func (g *Goblin) Attack() int{
	return 10
}
func (g *Goblin) HP() int{
	return 100
}