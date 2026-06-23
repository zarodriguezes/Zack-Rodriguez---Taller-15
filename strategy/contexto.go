package strategy

import "poo_taller15/model"

type EvalContext struct {
	Docente         model.Docente
	Competencia     model.Competencia
	Evaluacion      model.Evaluacion
	Antiguedad      int
	ReglasAplicadas map[string]bool
}
