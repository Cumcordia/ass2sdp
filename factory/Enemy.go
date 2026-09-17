package factory

import "fmt"

type Enemy interface {
	Name() string
	Attack() int
	HP() int
}

type creator interface {
	createEnemy() Enemy
}

type EnemyCreator struct {
	self creator
}

func (e *EnemyCreator) Spawn() string {
	c := e.self.createEnemy()
	return fmt.Sprint(c.Name(), c.HP(), c.Attack())
}