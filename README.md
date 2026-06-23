# Justificación de Diseño - Sistema de Evaluación Docente (Patrón Strategy)

1. ¿Por qué elegiste el patrón Strategy en lugar de Chain of Responsibility o Visitor?
Frente a Chain of Responsibility (CoR): CoR está diseñado para pasar una solicitud a lo largo de una cadena de manejadores donde, por lo general, *un solo* manejador la procesa o detiene el flujo si hay un error. En este sistema necesitamos lo contrario: evaluar muchas reglas independientes sobre un mismo docente y acumular de forma secuencial todos los porcentajes sobre el puntaje final.
Frente a Visitor: Visitor se usa para añadir operaciones sobre estructuras de objetos complejas o árboles sin modificar sus clases. Aquí la estructura de datos no es jerárquica ni compleja, sino un conjunto plano de metadatos del docente que requieren algoritmos matemáticos variables.

2. ¿Cómo manejaste el caso donde dos reglas tienen la misma prioridad?
Se resolvió mediante el uso de sort.Slice en el motor de cálculo (`engine/calculador.go`). 
 Esto garantiza que si dos estrategias comparten exactamente el mismo número de prioridad, el motor respetará estrictamente el orden físico en el que fueron registradas originalmente dentro de la lista del constructor.

3. ¿Por qué usaste composición en lugar de agregación para ciertas relaciones?
En la arquitectura de strategy.EvalContext, se utilizó composición para contener las estructuras de Docente, Competencia y Evaluacion. 
El ciclo de vida de este objeto contexto está completamente acoplado a la transacción del cálculo: nace cuando se va a calcular el puntaje y se destruye inmediatamente después por el Garbage Collector de Go. El contexto "es dueño" de esa foto instantánea de los datos mientras dura la evaluación.

4. ¿Cómo extenderías el sistema para soportar reglas que dependen del historial del docente?
Modificaría la estructura del contexto de evaluación (`strategy.EvalContext`) para inyectar una colección de registros pasados:

type EvalContext struct {
    // ... campos actuales
    HistorialEvaluaciones []model.EvaluacionHistorica
}