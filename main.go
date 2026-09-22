package main

import (
	"ass2sdp/factory"
	"ass2sdp/abstractfactory"
	"fmt"
)

func main(){
	runFactoryMethodDemo()
	runAbstractFactoryDemo()
}

func runFactoryMethodDemo() {
	creators := []factory.Spawner{
		factory.NewGoblinCreator(),
		factory.NewSkeletonCreator(),
		factory.NewSlimeCreator(),
	}

	for _, c := range creators {
		fmt.Println(c.Spawn())
	}
}
func runAbstractFactoryDemo() {
	worlds := []string{"forest", "ice", "desert"}

	for _, world := range worlds {
		factory, err := abstractfactory.GetWorldFactory(world)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		loader := abstractfactory.NewWorldLoader(factory)
		fmt.Printf("-- %s world --\n%s\n\n", world, loader.LoadWorld())
	}
}