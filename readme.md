## Desafío de cierre 

### Desafío
Ya está desarrollada una API que lista artículos disponibles para un vendedor determinado. Esta API no ha podido implementarse porque no cumple con los requisitos mínimos de calidad. Debe contener los tests necesarios para cubrir un 80% de coverage.

El proyecto consiste en una API que cuenta con un método llamado GetProducts y recibe por queryParam el ID (string) del vendedor (seller_id).

### Objetivos 
- Clonar y validar que la API esté en correcto funcionamiento.
- Realizar test unitario de la implementación products_map.go
- Aplicar Mock y Stub tests para desarrollar los Tests de los archivos:
    - Handler: crear el archivo products_default_test.go. Será necesario probar el handler en su totalidad.
- Alcanzar un coverage del 80% total del proyecto.
- Todos los tests ejecutados deben culminar ok.