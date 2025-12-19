<template>
  <v-footer style="background-color: var(--color-background); border-top: 1px solid var(--color-border);">
    <v-container>
      <v-row>
        <v-col cols="12" md="4">
          <h3 class="gradient-text text-subtitle-1 font-weight-bold mb-2">SoftPharos</h3>
          <p class="text-body-2" style="color: var(--color-text-secondary);">
            Iluminando el camino del desarrollo de software iterativo
          </p>
        </v-col>

        <v-col cols="12" md="4">
          <h4 class="text-subtitle-2 font-weight-bold mb-2" style="color: var(--color-text);">Universidad</h4>
          <p class="text-body-2 mb-1" style="color: var(--color-text-secondary);">Universidad Nacional de Colombia</p>
          <p class="text-body-2" style="color: var(--color-text-secondary);">Sede Bogotá</p>
        </v-col>

        <v-col cols="12" md="4">
          <h4 class="text-subtitle-2 font-weight-bold mb-2" style="color: var(--color-text);">Curso</h4>
          <p class="text-body-2 mb-1" style="color: var(--color-text-secondary);">Ingeniería de Software I</p>
          <p class="text-body-2" style="color: var(--color-text-secondary);">{{ currentYear }}</p>
        </v-col>
      </v-row>

      <v-divider class="my-4" style="border-color: var(--color-border);"></v-divider>

      <v-row>
        <v-col cols="12" class="text-center">
          <p class="text-caption" style="color: var(--color-text-secondary);">
            &copy; {{ currentYear }} SoftPharos. Proyecto académico UNAL.
          </p>
        </v-col>
      </v-row>
    </v-container>
  </v-footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>
