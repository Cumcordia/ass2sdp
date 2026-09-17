package factory

type Skeleton struct{}

type SkeletonCreator struct{
	EnemyCreator
}

func NewSkeletonCreator() *SkeletonCreator{
	s := &SkeletonCreator{}
	s.self = s
	return s
}

func (s *SkeletonCreator) createEnemy() Enemy{
	return &Skeleton{}
}

func (s *Skeleton) Name() string{
	return "Skeleton"
}
func (s *Skeleton) Attack() int{
	return 10
}
func (s *Skeleton) HP() int{
	return 150
}