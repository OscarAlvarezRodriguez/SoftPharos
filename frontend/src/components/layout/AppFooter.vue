<template>
  <v-footer class="app-footer">
    <v-container>
      <v-row>
        <v-col cols="12" md="4">
          <h3 class="gradient-text text-subtitle-1 font-weight-bold mb-2">SoftPharos</h3>
          <p class="text-body-2 footer-text">
            Iluminando el camino del desarrollo de software iterativo
          </p>
        </v-col>

        <v-col cols="12" md="4">
          <h4 class="text-subtitle-2 font-weight-bold mb-2 footer-heading">Universidad</h4>
          <p class="text-body-2 mb-1 footer-text">Universidad Nacional de Colombia</p>
          <p class="text-body-2 footer-text">Sede Bogotá</p>
        </v-col>

        <v-col cols="12" md="4">
          <h4 class="text-subtitle-2 font-weight-bold mb-2 footer-heading">Curso</h4>
          <p class="text-body-2 mb-1 footer-text">Ingeniería de Software I</p>
          <p class="text-body-2 footer-text">{{ currentYear }}</p>
        </v-col>
      </v-row>

      <v-divider class="my-4 footer-divider"></v-divider>

      <v-row>
        <v-col cols="12" class="text-center">
          <p class="text-caption footer-text">
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
