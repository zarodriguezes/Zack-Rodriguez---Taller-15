package strategy

// --- Regla 1: Bonificación Ciencias Investigación ---
type R1CienciasInvestigacion struct{}

func (r R1CienciasInvestigacion) Nombre() string {
	return "R1: Bonificación Ciencias Investigación (+10%)"
}
func (r R1CienciasInvestigacion) Prioridad() int { return 10 }
func (r R1CienciasInvestigacion) Aplica(ctx *EvalContext) bool {
	return ctx.Docente.Departamento == "Ciencias" && ctx.Competencia.Tipo == "investigacion"
}
func (r R1CienciasInvestigacion) Calcular(ctx *EvalContext, puntaje float64) float64 {
	return puntaje * 1.10
}

// --- Regla 2: Reducción Directores Gestión ---
type R2ReduccionDirectoresGestion struct{}

func (r R2ReduccionDirectoresGestion) Nombre() string {
	return "R2: Reducción Directores Gestión (-10%)"
}
func (r R2ReduccionDirectoresGestion) Prioridad() int { return 9 }
func (r R2ReduccionDirectoresGestion) Aplica(ctx *EvalContext) bool {
	return ctx.Docente.Cargo == "Director" && ctx.Competencia.Tipo == "gestion"
}
func (r R2ReduccionDirectoresGestion) Calcular(ctx *EvalContext, puntaje float64) float64 {
	return puntaje * 0.90
}

// --- Regla 3: Bonificación Antigüedad (Depende de R1) ---
type R3Antiguedad struct{}

func (r R3Antiguedad) Nombre() string { return "R3: Bonificación Antigüedad (+5%)" }
func (r R3Antiguedad) Prioridad() int { return 1 } // Prioridad baja para ejecutarse al final
func (r R3Antiguedad) Aplica(ctx *EvalContext) bool {
	return ctx.Antiguedad > 10 && ctx.ReglasAplicadas["R1: Bonificación Ciencias Investigación (+10%)"]
}
func (r R3Antiguedad) Calcular(ctx *EvalContext, puntaje float64) float64 {
	return puntaje * 1.05
}

// --- Regla 4: Bonificación Coordinadores Académica ---
type R4CoordinadoresAcademica struct{}

func (r R4CoordinadoresAcademica) Nombre() string {
	return "R4: Bonificación Coordinadores Académica (+5%)"
}
func (r R4CoordinadoresAcademica) Prioridad() int { return 8 }
func (r R4CoordinadoresAcademica) Aplica(ctx *EvalContext) bool {
	return ctx.Docente.Cargo == "Coordinador" && ctx.Competencia.Tipo == "academica"
}
func (r R4CoordinadoresAcademica) Calcular(ctx *EvalContext, puntaje float64) float64 {
	return puntaje * 1.05
}

// --- Regla 5: Bonificación Administración Gestión ---
type R5AdministracionGestion struct{}

func (r R5AdministracionGestion) Nombre() string {
	return "R5: Bonificación Administración Gestión (+15%)"
}
func (r R5AdministracionGestion) Prioridad() int { return 7 }
func (r R5AdministracionGestion) Aplica(ctx *EvalContext) bool {
	return ctx.Docente.Departamento == "Administración" && ctx.Competencia.Tipo == "gestion"
}
func (r R5AdministracionGestion) Calcular(ctx *EvalContext, puntaje float64) float64 {
	return puntaje * 1.15
}
