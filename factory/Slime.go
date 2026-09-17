package factory

type Slime struct{}

type SlimeCreator struct{
	EnemyCreator
}

func NewSlimeCreator() *SlimeCreator{
	s := &SlimeCreator{}
	s.self = s
	return s
}

func (s *SlimeCreator) createEnemy() Enemy{
	return &Slime{}
}

func (s *Slime) Name() string{
	return "Slime"
}
func (s *Slime) Attack() int{
	return 20
}
func (s *Slime) HP() int{
	return 200
}