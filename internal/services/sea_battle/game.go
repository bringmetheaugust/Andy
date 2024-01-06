package seabattle

type game struct {
	net     net
	netGrid int
}

func (g game) requestStep() step {
	var newStep step

	collumn := newStep.requestStepCollumn()
	raw := newStep.requestStepRaw()

	return step{collumn, raw}
}

func (g game) makeStep() {
	step := g.requestStep()
	g.net[step.collumn][step.raw].isChecked = true
	g.net.build()
}
