# Variables
APP_NAME=app-paises

.PHONY: setup run docker-build

BINARY_HEY=$(shell which hey)

# Instala dependencias y limpia código
setup:
	go mod tidy

# Ejecuta la app localmente
run:
	@go fmt ./...
	@go vet ./...
	@go run main.go

# Prueba de estrés inteligente
stress:
ifdef BINARY_HEY
	hey -n 1000 -c 10 http://localhost:8080/
else
	@echo "⚠️  'hey' no instalado. Instalando ahora..."
	go install github.com/rakyll/hey@latest
	@echo "Ejecuta 'make stress' de nuevo."
endif

# Test de integración y unidad
test:
	go test -v ./...

# Crea la imagen de Docker
docker-build:
	docker build -t $(APP_NAME) .

# Comando "Mágico" para Railway/Render (Simulado)
# La mayoría de estas plataformas detectan el Dockerfile automáticamente al hacer push
deploy:
	git add .
	git commit -m "🚀 Deploy automático: Mejoras de estabilidad"
	git push origin main
