package main

import (
	"fmt"
	"poo_taller15/engine"
	"poo_taller15/model"
)

func main() {
	calc := engine.NewCalculadorEngine()
	rep := engine.NewReporteEngine(calc)

	// Competencias de prueba comunes
	compInvestigacion := model.Competencia{ID: "C1", Nombre: "Artículos Científicos", Tipo: "investigacion"}
	compGestion := model.Competencia{ID: "C2", Nombre: "Dirección de Escuela", Tipo: "gestion"}

	// --- ESCENARIO 1: Ciencias + Investigación (Aplica R1) ---
	d1 := model.Docente{ID: "D1", Nombre: "Dr. Alberto", Departamento: "Ciencias", Cargo: "Profesor"}
	e1 := model.Evaluacion{ID: "E1", DocenteID: d1.ID, CompetenciaID: compInvestigacion.ID, PuntajeBase: 80.0}
	fmt.Println("--- ESCENARIO 1 ---")
	fmt.Println(rep.GenerarReporteJustificado(e1, d1, compInvestigacion, 5))

	// --- ESCENARIO 2: Director + Gestión (Aplica R2) ---
	d2 := model.Docente{ID: "D2", Nombre: "Ing. María", Departamento: "Ingeniería", Cargo: "Director"}
	e2 := model.Evaluacion{ID: "E2", DocenteID: d2.ID, CompetenciaID: compGestion.ID, PuntajeBase: 90.0}
	fmt.Println("--- ESCENARIO 2 ---")
	fmt.Println(rep.GenerarReporteJustificado(e2, d2, compGestion, 4))

	// --- ESCENARIO 3: Ciencias + Director + Gestión (Aplica R2, NO R1) ---
	d3 := model.Docente{ID: "D3", Nombre: "Dr. Carlos", Departamento: "Ciencias", Cargo: "Director"}
	e3 := model.Evaluacion{ID: "E3", DocenteID: d3.ID, CompetenciaID: compGestion.ID, PuntajeBase: 75.0}
	fmt.Println("--- ESCENARIO 3 ---")
	fmt.Println(rep.GenerarReporteJustificado(e3, d3, compGestion, 6))

	// --- ESCENARIO 4: Ciencias + Investigación + >10 años (Aplica R1 y R3) ---
	d4 := model.Docente{ID: "D4", Nombre: "Dra. Elena", Departamento: "Ciencias", Cargo: "Profesor"}
	e4 := model.Evaluacion{ID: "E4", DocenteID: d4.ID, CompetenciaID: compInvestigacion.ID, PuntajeBase: 100.0}
	fmt.Println("--- ESCENARIO 4 ---")
	fmt.Println(rep.GenerarReporteJustificado(e4, d4, compInvestigacion, 12))
}
