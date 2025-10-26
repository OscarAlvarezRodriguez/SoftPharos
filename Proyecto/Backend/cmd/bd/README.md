# 📊 Scripts de Base de Datos

Este directorio contiene los scripts SQL para la inicialización y población de la base de datos.

## 📁 Estructura de Archivos

```
bd/
├── init.sql       → Estructura de la BD (tablas, índices, relaciones)
├── seed.sql       → Datos esenciales para producción (roles, admin)
└── seed_dev.sql   → Datos de prueba para desarrollo (usuarios, proyectos demo)
```

## 🔄 Orden de Ejecución

1. **init.sql** - Crea todas las tablas y relaciones
2. **seed.sql** - Inserta datos esenciales (siempre se ejecuta)
3. **seed_dev.sql** - Inserta datos de desarrollo (opcional)

## 🚀 Uso

### Ejecución Automática (Recomendado)

El script `setup.sh` en la raíz del proyecto ejecuta automáticamente estos archivos:

```bash
./setup.sh
```

Esto hará:
1. ✅ Ejecutar `init.sql` si la BD está vacía
2. ✅ Ejecutar `seed.sql` automáticamente
3. ❓ Preguntar si quieres ejecutar `seed_dev.sql`

### Ejecución Manual

Si necesitas ejecutar los scripts manualmente:

```bash
# 1. Inicializar estructura
docker exec -i pg-demo-compose psql -U softpharos -d softpharos_db < init.sql

# 2. Poblar datos esenciales
docker exec -i pg-demo-compose psql -U softpharos -d softpharos_db < seed.sql

# 3. (Opcional) Poblar datos de desarrollo
docker exec -i pg-demo-compose psql -U softpharos -d softpharos_db < seed_dev.sql
```
