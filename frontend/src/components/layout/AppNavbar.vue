<template>
  <v-app-bar elevation="1" style="background-color: var(--color-background); border-bottom: 1px solid var(--color-border);">
    <v-container class="d-flex align-center">
      <v-app-bar-title>
        <router-link to="/" class="d-flex align-center text-decoration-none">
          <v-avatar size="40" class="mr-2">
            <svg width="40" height="40" viewBox="0 0 40 40">
              <circle cx="20" cy="20" r="18" :fill="colors.primary" opacity="0.2"/>
              <path d="M20 8L28 16L20 24L12 16L20 8Z" :fill="colors.primary"/>
              <path d="M20 16L28 24L20 32L12 24L20 16Z" :fill="colors.accent"/>
            </svg>
          </v-avatar>
          <span class="gradient-text text-h5 font-weight-bold">SoftPharos</span>
        </router-link>
      </v-app-bar-title>

      <v-spacer></v-spacer>

      <v-btn
        :icon="themeStore.currentTheme === 'dark' ? 'mdi-white-balance-sunny' : 'mdi-weather-night'"
        @click="themeStore.toggleTheme"
        variant="text"
      ></v-btn>

      <v-btn to="/login" variant="outlined" color="primary" class="ml-2">
        Iniciar Sesión
      </v-btn>

      <v-btn to="/register" color="primary" class="ml-2">
        Registrarse
      </v-btn>
    </v-container>
  </v-app-bar>
</template>

<script setup>
import { computed } from 'vue'
import { useThemeStore } from '@/stores/theme'

const themeStore = useThemeStore()

const colors = computed(() => ({
  primary: 'var(--color-primary)',
  accent: 'var(--color-accent)'
}))
</script>
