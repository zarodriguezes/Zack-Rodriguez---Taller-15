package model

import "time"

type EstadoEvaluacion string

const (
	EvaluacionPendiente  EstadoEvaluacion = "pendiente"
	EvaluacionEnProceso  EstadoEvaluacion = "en_proceso"
	EvaluacionFinalizada EstadoEvaluacion = "finalizada"
)

type Evaluacion struct {
	ID           string
	DocenteID    string
	PeriodoID    string
	Fecha        time.Time
	Estado       EstadoEvaluacion
	PuntajeBase  float64 // Puntaje sin reglas
	PuntajeTotal float64 // Puntaje con reglas aplicadas
}
