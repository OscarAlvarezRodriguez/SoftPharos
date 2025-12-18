#!/bin/bash

# Script para ejecutar tests y verificar cobertura de código
cd /Users/oscalvarez/Documents/GIT/SoftPharos/Proyecto/Backend

echo "========================================="
echo "Ejecutando todos los tests..."
echo "========================================="

# Ejecutar tests con coverage
go test ./... -coverprofile=coverage.out -v 2>&1 | tee test_output.log

echo ""
echo "========================================="
echo "Resumen de cobertura por paquete"
echo "========================================="
go tool cover -func=coverage.out | grep -E "^(softpharos/|total:)" | tail -30

echo ""
echo "========================================="
echo "Cobertura total"
echo "========================================="
go tool cover -func=coverage.out | grep total

echo ""
echo "========================================="
echo "Generando reporte HTML de cobertura..."
echo "========================================="
go tool cover -html=coverage.out -o coverage.html
echo "Reporte generado en coverage.html"

echo ""
echo "========================================="
echo "Tests con fallos (si los hay):"
echo "========================================="
grep "FAIL" test_output.log | grep -v "coverage:" || echo "No hay tests fallidos"

echo ""
echo "========================================="
echo "Paquetes sin tests:"
echo "========================================="
grep "no test files" test_output.log | head -10
