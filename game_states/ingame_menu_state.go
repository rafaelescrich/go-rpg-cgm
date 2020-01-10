package game_states

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/steelx/go-rpg-cgm/gui"
	"github.com/steelx/go-rpg-cgm/state_machine"
)

type InGameMenuState struct {
	Stack        *gui.StateStack
	StateMachine *state_machine.StateMachine
}

func InGameMenuStateCreate(stack *gui.StateStack, win *pixelgl.Window) *InGameMenuState {
	igm := &InGameMenuState{
		Stack: stack,
	}

	igm.StateMachine = state_machine.Create(map[string]func() state_machine.State{
		"frontmenu": func() state_machine.State {
			return FrontMenuStateCreate(igm, win)
		},
		"items": func() state_machine.State {
			return ItemsMenuStateCreate(igm, win)
		},
		"magic": func() state_machine.State {
			//return MagicMenuStateCreate(this)
			return state_machine.Create(map[string]func() state_machine.State{})
		},
		"equip": func() state_machine.State {
			//return EquipMenuStateCreate(this)
			return state_machine.Create(map[string]func() state_machine.State{})
		},
		"status": func() state_machine.State {
			//return StatusMenuStateCreate(this)
			return state_machine.Create(map[string]func() state_machine.State{})
		},
	})

	igm.StateMachine.Change("frontmenu", nil)

	return igm
}

func (igm *InGameMenuState) Update(dt float64) bool {
	igm.StateMachine.Update(dt)
	//fmt.Println("ingame_menu_state", reflect.DeepEqual(igm.Stack.Top(), igm)) // temp
	//if reflect.DeepEqual(igm.Stack.Top(), igm) {
	//	igm.StateMachine.Update(dt)
	//}
	return true
}
func (igm InGameMenuState) Render(win *pixelgl.Window) {
	igm.StateMachine.Render(win)

	//temp camera matrix
	cam := pixel.IM.Scaled(pixel.V(0, 0), 1.0).Moved(win.Bounds().Center())
	win.SetMatrix(cam)
}

func (igm InGameMenuState) Enter()                          {}
func (igm InGameMenuState) Exit()                           {}
func (igm InGameMenuState) HandleInput(win *pixelgl.Window) {}
