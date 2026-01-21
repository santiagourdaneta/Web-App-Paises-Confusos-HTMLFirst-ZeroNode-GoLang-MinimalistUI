#!/bin/sh
echo "🔍 Validando código antes de guardar..."

# Formatea el código automáticamente (Formatter)
go fmt ./...

# Revisa errores lógicos y de sintaxis (Linter)
go vet ./...

# Si algo falla, detenemos el commit
if [ $? -ne 0 ]; then
 echo "❌ Error: El código no pasa las pruebas de calidad."
 exit 1
fi

echo "✅ Todo limpio. Guardando..."