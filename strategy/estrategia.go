package strategy

type EstrategiaPonderacion interface {
	Calcular(ctx *EvalContext, puntajeActual float64) float64
	Aplica(ctx *EvalContext) bool
	Nombre() string
	Prioridad() int
}
