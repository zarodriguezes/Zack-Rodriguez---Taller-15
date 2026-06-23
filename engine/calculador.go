package engine

import (
	"poo_taller15/model"
	"poo_taller15/strategy"
	"sort"
)

type CalculadorEngine struct {
	Reglas []strategy.EstrategiaPonderacion
}

func NewCalculadorEngine() *CalculadorEngine {
	return &CalculadorEngine{
		Reglas: []strategy.EstrategiaPonderacion{
			strategy.R1CienciasInvestigacion{},
			strategy.R2ReduccionDirectoresGestion{},
			strategy.R3Antiguedad{},
			strategy.R4CoordinadoresAcademica{},
			strategy.R5AdministracionGestion{},
		},
	}
}

func (ce *CalculadorEngine) CalcularPuntajeConReglas(eval model.Evaluacion, doc model.Docente, comp model.Competencia, ant int) (float64, []string) {
	ctx := &strategy.EvalContext{
		Docente:         doc,
		Competencia:     comp,
		Evaluacion:      eval,
		Antiguedad:      ant,
		ReglasAplicadas: make(map[string]bool),
	}

	// Hacemos una copia de las reglas para ordenarlas sin alterar la lista original
	reglasOrdenadas := make([]strategy.EstrategiaPonderacion, len(ce.Reglas))
	copy(reglasOrdenadas, ce.Reglas)

	// Ordenamos de mayor a menor prioridad usando sort.Slice
	sort.Slice(reglasOrdenadas, func(i, j int) bool {
		return reglasOrdenadas[i].Prioridad() > reglasOrdenadas[j].Prioridad()
	})

	puntajeFinal := eval.PuntajeBase
	var aplicadas []string

	// Recorremos las reglas de forma dinámica sin usar un switch masivo
	for _, regla := range reglasOrdenadas {
		if regla.Aplica(ctx) {
			puntajeFinal = regla.Calcular(ctx, puntajeFinal)
			ctx.ReglasAplicadas[regla.Nombre()] = true
			aplicadas = append(aplicadas, regla.Nombre())
		}
	}

	return puntajeFinal, aplicadas
}
