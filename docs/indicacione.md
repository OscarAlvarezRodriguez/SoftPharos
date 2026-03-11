# Guía Técnica: Primera Aplicación Full-Stack
## Vue.js + Django + PostgreSQL

---

## Índice

1. [Introducción](#1-introducción)
2. [HTML, CSS y Frameworks: ¿Por qué Vue.js?](#2-html-css-y-frameworks-por-qué-vuejs)
3. [Requisitos e Instalación](#3-requisitos-e-instalación)
4. [Arquitectura del Sistema](#4-arquitectura-del-sistema)
5. [Parte 1: Base de Datos (PostgreSQL)](#5-parte-1-base-de-datos-postgresql)
6. [Parte 2: Backend (Django)](#6-parte-2-backend-django)
7. [Parte 3: Frontend (Vue.js)](#7-parte-3-frontend-vuejs)
8. [Parte 4: Integración y Pruebas](#8-parte-4-integración-y-pruebas)
9. [Comandos de Referencia](#9-comandos-de-referencia)
10. [Evidencias Requeridas (Screenshots)](#10-evidencias-requeridas-screenshots)

---

## 1. Introducción

### Objetivo del Ejercicio

Construir un **sistema de login mínimo** que integre tres tecnologías fundamentales del desarrollo web:

```
┌─────────────┐     POST /api/login    ┌─────────────┐      Query       ┌─────────────┐
│   Vue.js    │ ────────────────────►  │   Django    │ ───────────────► │ PostgreSQL  │
│  (Frontend) │ ◄────────────────────  │  (Backend)  │ ◄─────────────── │ (Database)  │
└─────────────┘     JSON Response      └─────────────┘     Results      └─────────────┘
```

### Flujo Funcional

1. Usuario ingresa correo y contraseña en pantalla de login
2. Frontend envía petición `POST /api/login` al backend
3. Backend busca el usuario por correo en la base de datos
4. Backend aplica hash SHA512 al password recibido y compara con el almacenado
5. Si coincide, devuelve `{ success: true, secret_phrase: "..." }`
6. Frontend redirige a segunda pantalla mostrando la frase secreta

---

## 2. HTML, CSS y Frameworks: ¿Por qué Vue.js?

### HTML y CSS: Los Fundamentos

| Tecnología | Función | Ejemplos |
|------------|---------|----------|
| **HTML** | Define la **estructura** de la página | `<div>`, `<form>`, `<input>`, `<button>` |
| **CSS** | Define la **apariencia visual** | colores, tamaños, espaciado, sombras |
| **JavaScript** | Define el **comportamiento** | eventos, validaciones, llamadas a API |

### Comparativa: JavaScript Nativo vs Vue.js

| Tarea | JavaScript Nativo | Vue.js |
|-------|-------------------|--------|
| **Capturar valor de input** | `document.getElementById('email').value` | `v-model="email"` → automático |
| **Mostrar variable en HTML** | `element.textContent = nombre` | `{{ nombre }}` → automático |
| **Actualizar UI al cambiar dato** | Manipular DOM manualmente | Reactivo → se actualiza solo |
| **Escuchar click en botón** | `btn.addEventListener('click', fn)` | `@click="fn"` |
| **Mostrar/ocultar elemento** | `element.style.display = 'none'` | `v-if="condicion"` |
| **Iterar lista en HTML** | Crear elementos con `createElement()` | `v-for="item in lista"` |
| **Manejar estado** | Variables globales dispersas | `data()` organizado por componente |
| **Navegar entre pantallas** | `window.location` o manipular history | `router.push('/ruta')` |
| **Reutilizar código UI** | Copiar/pegar HTML | Importar componente |

### ¿Por qué usar un Framework?

| Aspecto | Sin Framework | Con Vue.js |
|---------|---------------|------------|
| **Líneas de código** | Muchas | Menos |
| **Mantenibilidad** | Difícil | Fácil |
| **Escalabilidad** | Compleja | Simple |
| **Errores comunes** | Frecuentes | Reducidos |
| **Curva de aprendizaje inicial** | Baja | Media |
| **Productividad a largo plazo** | Baja | Alta |

### Ejemplo Comparativo

**Sin Framework (JavaScript Nativo):**
```
1. Crear elementos HTML manualmente
2. Agregar event listeners con addEventListener()
3. Manipular DOM con document.querySelector()
4. Actualizar texto con element.textContent = valor
5. Manejar navegación manipulando window.location
```

**Con Vue.js:**
```
1. Declarar datos en data() → Se muestran automáticamente
2. Usar v-model → Input sincronizado con variable
3. Usar @click → Evento conectado a método
4. Usar v-if → Mostrar/ocultar según condición
5. Usar router.push() → Navegar a otra vista
```

### Anatomía de un Componente Vue

Un archivo `.vue` tiene **tres secciones**:

```
┌─────────────────────────────────────┐
│  <template>                         │  ← HTML del componente
│    Estructura visual                │
│  </template>                        │
├─────────────────────────────────────┤
│  <script>                           │  ← JavaScript/Lógica
│    - Datos (data)                   │
│    - Métodos (methods)              │
│    - Ciclo de vida (created, etc)   │
│  </script>                          │
├─────────────────────────────────────┤
│  <style scoped>                     │  ← CSS (solo afecta este componente)
│    Estilos visuales                 │
│  </style>                           │
└─────────────────────────────────────┘
```

### Recomendaciones de Estilo para el Ejercicio

Para el formulario de login, aplicar:

- **Contenedor centrado:** Usar `display: flex` con `justify-content: center`
- **Fondo atractivo:** Gradiente con `background: linear-gradient(...)`
- **Tarjeta del formulario:** Fondo blanco, `border-radius`, `box-shadow`
- **Inputs claros:** Bordes suaves, padding generoso, efecto focus
- **Botón llamativo:** Color contrastante, efecto hover con `transform`
- **Mensajes de error:** Fondo rojo claro, texto rojo oscuro

---

## 3. Requisitos e Instalación

### Software Necesario

| Software | Versión Mínima | Propósito |
|----------|----------------|-----------|
| Python | 3.10+ | Backend con Django |
| Node.js | 18+ | Runtime para Vue.js |
| npm | 9+ | Gestor de paquetes JS |
| PostgreSQL | 14+ | Base de datos |
| Git | 2.30+ | Control de versiones |

### Verificación de Instalaciones

Ejecutar en terminal:

```bash
python3 --version
node --version
npm --version
psql --version
git --version
```

### Instalación por Sistema Operativo

**macOS:**
```bash
brew install python node postgresql@14 git
brew services start postgresql@14
```

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install python3 python3-pip python3-venv nodejs npm postgresql git
sudo systemctl start postgresql
```

**Windows:**
- Descargar instaladores oficiales de cada software
- O usar Chocolatey: `choco install python nodejs postgresql git`

---

## 4. Arquitectura del Sistema

### Arquitectura por Capas - Backend

```
┌────────────────────────────────────────────────────────┐
│                      BACKEND                           │
├────────────────────────────────────────────────────────┤
│                                                        │
│   CONTROLLER (Views)                                   │
│   └── Recibe HTTP Request, valida, devuelve Response   │
│                    │                                   │
│                    ▼                                   │
│   SERVICE                                              │
│   └── Lógica de negocio, orquesta operaciones          │
│                    │                                   │
│                    ▼                                   │
│   REPOSITORY                                           │
│   └── Acceso a datos, queries a PostgreSQL             │
│                    │                                   │
│                    ▼                                   │
│   DOMAIN (Models)                                      │
│   └── Definición de entidades/tablas                   │
│                                                        │
└────────────────────────────────────────────────────────┘
```

### Arquitectura por Capas - Frontend

```
┌────────────────────────────────────────────────────────┐
│                      FRONTEND                          │
├────────────────────────────────────────────────────────┤
│                                                        │
│   VIEWS                                                │
│   └── Páginas completas (LoginView, SecretView)        │
│                    │                                   │
│                    ▼                                   │
│   COMPONENTS                                           │
│   └── Elementos reutilizables (LoginForm, etc)         │
│                    │                                   │
│                    ▼                                   │
│   SERVICES (API)                                       │
│   └── Comunicación HTTP con el backend                 │
│                                                        │
└────────────────────────────────────────────────────────┘
```

### Estructura de Carpetas

```
ingesoft1-login/
├── backend/
│   ├── manage.py
│   ├── requirements.txt
│   ├── config/
│   │   ├── settings.py      ← Configuración Django + BD
│   │   └── urls.py          ← Rutas principales
│   └── auth_app/
│       ├── controllers/
│       │   └── auth_controller.py
│       ├── services/
│       │   └── auth_service.py
│       ├── repositories/
│       │   └── user_repository.py
│       ├── domain/
│       │   └── models.py
│       └── urls.py
│
└── frontend/
    ├── package.json
    └── src/
        ├── main.js
        ├── App.vue
        ├── router/
        │   └── index.js
        ├── views/
        │   ├── LoginView.vue
        │   └── SecretView.vue
        ├── components/
        │   ├── LoginForm.vue
        │   └── SecretDisplay.vue
        └── services/
            └── authService.js
```

---

## 5. Parte 1: Base de Datos (PostgreSQL)

### Paso 5.1: Acceder a PostgreSQL

Conectarse como superusuario para crear la base de datos.

```bash
# Linux/macOS
sudo -u postgres psql

# Windows
psql -U postgres
```

### Paso 5.2: Crear Base de Datos y Usuario

Dentro del cliente `psql`, ejecutar comandos SQL para:

1. **Crear un usuario** llamado `ingesoft1_user` con una contraseña
2. **Crear una base de datos** llamada `ingesoft1_db`
3. **Asignar permisos** al usuario sobre la base de datos
4. **Verificar** la conexión con `\conninfo`

### Paso 5.3: Definir la Tabla

La tabla `ingesoft1_users` debe tener:

| Columna | Tipo | Restricciones |
|---------|------|---------------|
| `id` | Integer | Primary Key, Auto-increment |
| `correo` | String (255) | Unique, Not Null |
| `password_hash` | String (128) | Not Null |
| `secret_phrase` | String (255) | Not Null |

> **Nota:** La tabla se creará automáticamente con las migraciones de Django.

---

## 6. Parte 2: Backend (Django)

### Paso 6.1: Crear Proyecto

1. Crear carpeta `backend/`
2. Crear y activar entorno virtual de Python
3. Crear archivo `requirements.txt` con dependencias:
    - Django
    - djangorestframework
    - psycopg2-binary (conector PostgreSQL)
    - django-cors-headers
4. Instalar dependencias con `pip install -r requirements.txt`
5. Crear proyecto Django con `django-admin startproject config .`
6. Crear aplicación con `python manage.py startapp auth_app`

### Paso 6.2: Configurar Django (settings.py)

Modificar `config/settings.py` para:

1. **Agregar apps instaladas:**
    - `rest_framework`
    - `corsheaders`
    - `auth_app`

2. **Configurar middleware CORS** (debe ir primero en la lista)

3. **Configurar base de datos PostgreSQL:**
    - ENGINE: postgresql
    - NAME: ingesoft1_db
    - USER: ingesoft1_user
    - PASSWORD: (la que definiste)
    - HOST: localhost
    - PORT: 5432

4. **Configurar CORS** para permitir peticiones desde `http://localhost:5173`

### Paso 6.3: Capa Domain (models.py)

Ubicación: `auth_app/domain/models.py`

**Qué debe hacer:**
- Definir clase `Ingesoft1User` que herede de `models.Model`
- Definir campo `correo` como EmailField, único
- Definir campo `password_hash` como CharField (128 caracteres)
- Definir campo `secret_phrase` como CharField (255 caracteres)
- Configurar `Meta.db_table = 'ingesoft1_users'`

### Paso 6.4: Capa Repository (user_repository.py)

Ubicación: `auth_app/repositories/user_repository.py`

**Qué debe hacer:**
- Crear clase `UserRepository`
- Método `find_by_correo(correo)`:
    - Buscar usuario por correo usando el ORM de Django
    - Retornar el usuario si existe, o `None` si no existe
- Método `create_user(correo, password_hash, secret_phrase)`:
    - Crear y guardar un nuevo usuario en la base de datos

### Paso 6.5: Capa Service (auth_service.py)

Ubicación: `auth_app/services/auth_service.py`

**Qué debe hacer:**
- Crear clase `AuthService`
- Método estático `hash_password(password)`:
    - Importar `hashlib`
    - Aplicar SHA512 al password
    - Retornar el hash en formato hexadecimal
- Método `authenticate(correo, password)`:
    - Usar el repository para buscar el usuario por correo
    - Si no existe, retornar `{ success: False, message: "Usuario no encontrado" }`
    - Si existe, aplicar hash al password recibido
    - Comparar con el password_hash almacenado
    - Si no coincide, retornar `{ success: False, message: "Contraseña incorrecta" }`
    - Si coincide, retornar `{ success: True, secret_phrase: "..." }`

### Paso 6.6: Capa Controller (auth_controller.py)

Ubicación: `auth_app/controllers/auth_controller.py`

**Qué debe hacer:**
- Importar decorador `@api_view` de rest_framework
- Crear función `login(request)` decorada con `@api_view(['POST'])`
- Extraer `correo` y `password` del `request.data`
- Validar que ambos campos existan (retornar 400 si faltan)
- Instanciar `AuthService` y llamar `authenticate()`
- Retornar respuesta JSON con código 200 si éxito, 401 si falla

### Paso 6.7: Configurar URLs

1. **auth_app/urls.py:**
    - Definir ruta `path('login', login)` apuntando al controller

2. **config/urls.py:**
    - Incluir las URLs de auth_app bajo el prefijo `api/`
    - Resultado: `POST /api/login`

### Paso 6.8: Migraciones

```bash
python manage.py makemigrations
python manage.py migrate
```

### Paso 6.9: Crear Usuario de Prueba

Crear un comando de Django o usar el shell para insertar un usuario de prueba:
- Correo: `estudiante@universidad.edu.co`
- Password: `MiPassword123` (almacenar su hash SHA512)
- Frase secreta: `¡Felicitaciones! Has completado el ejercicio de integración Full-Stack.`

### Paso 6.10: Probar Backend

```bash
python manage.py runserver
```

Probar con curl o Postman:
```bash
curl -X POST http://localhost:8000/api/login \
  -H "Content-Type: application/json" \
  -d '{"correo": "estudiante@universidad.edu.co", "password": "MiPassword123"}'
```

---

## 7. Parte 3: Frontend (Vue.js)

### Paso 7.1: Crear Proyecto Vue

```bash
npm create vite@latest frontend -- --template vue
cd frontend
npm install
npm install vue-router@4 axios
```

### Paso 7.2: Crear Estructura de Carpetas

Dentro de `src/`, crear:
- `views/` - Para las vistas (páginas)
- `components/` - Para componentes reutilizables
- `services/` - Para comunicación con API
- `router/` - Para configuración de rutas

### Paso 7.3: Capa Services (authService.js)

Ubicación: `src/services/authService.js`

**Qué debe hacer:**
- Importar `axios`
- Crear instancia de axios con baseURL `http://localhost:8000/api`
- Exportar objeto con método `login(correo, password)`:
    - Hacer POST a `/login` con los datos
    - Manejar respuesta exitosa y errores
    - Retornar el objeto de respuesta del servidor

### Paso 7.4: Componente LoginForm.vue

Ubicación: `src/components/LoginForm.vue`

**Template (HTML) - Qué debe contener:**
- Contenedor principal centrado
- Formulario con `@submit.prevent` para evitar recarga
- Input para correo con `v-model` vinculado a variable
- Input para password con `v-model` vinculado a variable
- Div condicional (`v-if`) para mostrar mensajes de error
- Botón submit que muestre "Ingresando..." durante carga

**Script (JS) - Qué debe hacer:**
- Importar el servicio de autenticación
- Definir datos: `correo`, `password`, `loading`, `errorMessage`
- Método `handleSubmit()`:
    - Activar estado de carga
    - Llamar al servicio de login
    - Si éxito, emitir evento `login-success` con la frase secreta
    - Si error, mostrar mensaje de error
    - Desactivar estado de carga

**Style (CSS) - Qué aplicar:**
- Contenedor: centrado vertical y horizontal, fondo gradiente
- Formulario: fondo blanco, padding, border-radius, sombra
- Inputs: bordes suaves, padding, efecto focus con color
- Botón: gradiente, color blanco, efecto hover
- Error: fondo rojo claro, texto rojo, padding

### Paso 7.5: Componente SecretDisplay.vue

Ubicación: `src/components/SecretDisplay.vue`

**Template - Qué debe contener:**
- Contenedor con fondo verde (éxito)
- Tarjeta blanca centrada
- Ícono de check (puede ser emoji ✓ o SVG)
- Título "¡Login Exitoso!"
- Caja destacada con la frase secreta (recibida como prop)
- Botón para cerrar sesión

**Script - Qué debe hacer:**
- Definir prop `secretPhrase` (requerida, tipo String)
- Método `handleLogout()` que emita evento `logout`

**Style - Qué aplicar:**
- Fondo gradiente verde
- Tarjeta centrada con sombra
- Frase secreta en caja con borde punteado
- Botón de logout con estilo secundario

### Paso 7.6: Vista LoginView.vue

Ubicación: `src/views/LoginView.vue`

**Qué debe hacer:**
- Importar y usar componente `LoginForm`
- Escuchar evento `login-success`
- Al recibir la frase secreta, navegar a la ruta `/secret` pasando la frase como parámetro

### Paso 7.7: Vista SecretView.vue

Ubicación: `src/views/SecretView.vue`

**Qué debe hacer:**
- Importar y usar componente `SecretDisplay`
- En `created()`, obtener la frase secreta de los parámetros de ruta
- Si no hay frase secreta, redirigir al login
- Escuchar evento `logout` y redirigir al login

### Paso 7.8: Configurar Router (router/index.js)

**Qué debe hacer:**
- Importar `createRouter` y `createWebHistory` de vue-router
- Importar las vistas
- Definir rutas:
    - `/` → LoginView
    - `/secret/:secretPhrase?` → SecretView
- Crear y exportar el router

### Paso 7.9: Configurar App.vue y main.js

**App.vue:**
- Solo contener `<router-view />` para renderizar la vista activa
- Agregar estilos globales básicos (reset de margin/padding)

**main.js:**
- Importar createApp de Vue
- Importar App y router
- Crear app, usar router, montar en #app

### Paso 7.10: Ejecutar Frontend

```bash
npm run dev
```

Acceder a `http://localhost:5173`

---

## 8. Parte 4: Integración y Pruebas

### Verificación del Sistema Completo

1. **Terminal 1:** Backend corriendo (`python manage.py runserver`)
2. **Terminal 2:** Frontend corriendo (`npm run dev`)
3. **PostgreSQL:** Servicio activo

### Casos de Prueba

| # | Acción | Resultado Esperado |
|---|--------|-------------------|
| 1 | Abrir `localhost:5173` | Ver formulario de login |
| 2 | Enviar formulario vacío | Validación HTML5 impide envío |
| 3 | Correo inexistente | Mensaje "Usuario no encontrado" |
| 4 | Password incorrecto | Mensaje "Contraseña incorrecta" |
| 5 | Credenciales correctas | Redirección a pantalla con frase |
| 6 | Click "Cerrar Sesión" | Volver al formulario de login |

### Verificar en DevTools del Navegador

1. Abrir DevTools (F12)
2. Ir a pestaña **Network**
3. Realizar login
4. Verificar:
    - Petición POST a `/api/login`
    - Headers Content-Type: application/json
    - Payload con correo y password
    - Respuesta JSON del servidor

---

## 9. Comandos de Referencia

### PostgreSQL

```bash
psql -U usuario -d basedatos -h localhost   # Conectar
\l                                           # Listar bases de datos
\dt                                          # Listar tablas
\d nombre_tabla                              # Describir tabla
\q                                           # Salir
```

### Django

```bash
source venv/bin/activate          # Activar entorno (Linux/macOS)
venv\Scripts\activate             # Activar entorno (Windows)
python manage.py runserver        # Iniciar servidor
python manage.py makemigrations   # Crear migraciones
python manage.py migrate          # Aplicar migraciones
python manage.py shell            # Shell interactivo
```

### Vue.js / npm

```bash
npm create vite@latest nombre -- --template vue   # Crear proyecto
npm install                                        # Instalar dependencias
npm install paquete                               # Agregar dependencia
npm run dev                                       # Servidor desarrollo
npm run build                                     # Build producción
```

### cURL (Pruebas de API)

```bash
curl -X POST http://localhost:8000/api/login \
  -H "Content-Type: application/json" \
  -d '{"correo": "email@test.com", "password": "123456"}'
```

---

## 10. Evidencias Requeridas (Screenshots)

Documentar el progreso en un **Google Docs** con las siguientes capturas de pantalla.

### Formato del Documento

```
EVIDENCIAS - EJERCICIO FULL-STACK
Nombre: [Tu nombre]
Fecha: [Fecha de entrega]
Curso: Ingeniería de Software I
```

### Screenshots Obligatorios

| # | Descripción | Qué debe verse en la captura |
|---|-------------|------------------------------|
| **1** | PostgreSQL instalado | Terminal con `psql --version` |
| **2** | Base de datos creada | Resultado de `\conninfo` en psql mostrando conexión a `ingesoft1_db` |
| **3** | Tabla creada | Resultado de `\d ingesoft1_users` mostrando columnas |
| **4** | Usuario de prueba | Query `SELECT correo, secret_phrase FROM ingesoft1_users;` |
| **5** | Estructura backend | Explorador de archivos mostrando carpetas controllers/, services/, repositories/, domain/ |
| **6** | Django corriendo | Terminal con servidor iniciado en puerto 8000 |
| **7** | API - Login exitoso | Respuesta curl/Postman con `success: true` |
| **8** | API - Login fallido | Respuesta curl/Postman con `success: false` |
| **9** | Estructura frontend | Explorador mostrando carpetas views/, components/, services/ |
| **10** | Vue corriendo | Terminal con servidor Vite en puerto 5173 |
| **11** | Pantalla login | Navegador mostrando formulario vacío |
| **12** | Error de login | Navegador mostrando mensaje de error |
| **13** | Login exitoso | Navegador mostrando pantalla con frase secreta |
| **14** | DevTools - Request | Panel Network mostrando petición POST /api/login |
| **15** | DevTools - Response | Panel Network mostrando respuesta JSON del servidor |

### Sección Final del Documento
```
TIEMPO TOTAL INVERTIDO:
[Horas aproximadas]
```

**¡Éxito en tu ejercicio de integración Full-Stack!**
