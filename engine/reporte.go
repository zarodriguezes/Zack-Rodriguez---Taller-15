package engine

import (
	"fmt"
	"poo_taller15/model"
)

func GenerarReporte(e model.Evaluacion, puntajeFinal float64, reglas []string) string {

	reporte := fmt.Sprintf("Evaluación ID: %s\n", e.ID)
	reporte += fmt.Sprintf("Puntaje base: %.2f\n", e.PuntajeBase)
	reporte += "Reglas aplicadas:\n"

	for _, r := range reglas {
		reporte += "- " + r + "\n"
	}

	reporte += fmt.Sprintf("\nPuntaje final: %.2f\n", puntajeFinal)

	return reporte
}
