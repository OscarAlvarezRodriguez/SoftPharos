# Frontend - SoftPharos

Aplicación web desarrollada con Vue 3, Vite, Vuetify 3 y Pinia.

## 📋 Requisitos

- Node.js 20.19.0+ o 22.12.0+
- npm 9+

## 🚀 Comandos

```bash
# Instalar dependencias
npm install

# Desarrollo (puerto 5173)
npm run dev

# Build para producción
npm run build

# Preview producción
npm run preview

# Tests
npm run test:unit

# Linter
npm run lint
```

## 📁 Estructura

```
src/
├── assets/
│   └── styles/
│       └── theme.css       # CSS global de temas
├── components/
│   ├── layout/
│   │   ├── AppNavbar.vue
│   │   └── AppFooter.vue
│   └── home/
│       ├── HomeHero.vue
│       ├── HomeStats.vue
│       └── HomeFeatures.vue
├── views/
│   ├── HomeView.vue
│   ├── LoginView.vue
│   └── RegisterView.vue
├── stores/
│   └── theme.js            # Store de temas (Pinia)
├── plugins/
│   └── vuetify.js          # Config Vuetify
├── router/
│   └── index.js
├── App.vue
└── main.js
```

## 🔗 API Backend

Por defecto conecta a `http://localhost:8080`
