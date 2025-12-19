<template>
  <v-app>
    <AppNavbar />

    <v-main>
      <RouterView />
    </v-main>

    <AppFooter />
  </v-app>
</template>

<script setup>
import { onMounted } from 'vue'
import { useTheme } from 'vuetify'
import { useThemeStore } from '@/stores/theme'
import AppNavbar from '@/components/layout/AppNavbar.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

const themeStore = useThemeStore()
const vuetifyTheme = useTheme()

onMounted(() => {
  themeStore.setVuetifyTheme(vuetifyTheme)
  themeStore.initTheme()
})
</script>
