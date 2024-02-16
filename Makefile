all: one two three four five
# correr tests en go.
one:
	go test ./...
# generar un reporte de coverage con el nombre “coverage.out”.
two:
	go test -cover -coverprofile=coverage.out ./...
# obtener información del coverage con un template de html.
three:
	go tool cover -html=coverage.out
# obtener información del coverage total del proyecto.
four:
	go tool cover -func=coverage.out 
# ejecutar el linter.
five:
	staticcheck ./...
