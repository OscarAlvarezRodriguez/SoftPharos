import { ref } from 'vue'
import { defineStore } from 'pinia'
import { useTheme } from 'vuetify'

export const useThemeStore = defineStore('theme', () => {
  const currentTheme = ref(localStorage.getItem('theme') || 'light')
  let vuetifyTheme = null

  function setVuetifyTheme(theme) {
    vuetifyTheme = theme
  }

  function toggleTheme() {
    currentTheme.value = currentTheme.value === 'light' ? 'dark' : 'light'
    localStorage.setItem('theme', currentTheme.value)
    applyTheme()
  }

  function applyTheme() {
    document.documentElement.setAttribute('data-theme', currentTheme.value)
    if (vuetifyTheme) {
      vuetifyTheme.global.name.value = currentTheme.value
    }
  }

  function initTheme() {
    applyTheme()
  }

  return {
    currentTheme,
    toggleTheme,
    initTheme,
    setVuetifyTheme
  }
})
