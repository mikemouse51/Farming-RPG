package screens

import (
	"github.com/phestek/farming-rpg/eng"
	"github.com/phestek/farming-rpg/entities"
	"github.com/phestek/farming-rpg/world"
)

type TestScreen struct {
	entities  []entities.Entitier
	testChunk world.Chunk
}

func (screen *TestScreen) Init(window *eng.Window) {
	screen.entities = append(screen.entities, entities.NewTeazel(window, eng.Vector2f{300, 200}))
	screen.testChunk.TilesBatch = eng.NewSpriteBatch(window, eng.Vector2i{512, 512})
	screen.testChunk.Tiles = [256]int {
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}
	screen.testChunk.Redraw()
}

func (screen *TestScreen) Dispose() {
	screen.entities = nil
}

func (screen *TestScreen) Update(delta float32) {
	for i := 0; i < len(screen.entities); i++ {
		screen.entities[i].Update(delta)
	}
}

func (screen *TestScreen) Draw(window *eng.Window) {
	window.Clear(eng.Color{15, 15, 15, 255})
	window.Draw(&screen.testChunk.TilesBatch)
	for i := 0; i < len(screen.entities); i++ {
		screen.entities[i].Draw(window)
	}
	window.Display()
}
