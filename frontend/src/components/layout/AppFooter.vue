<template>
  <v-footer class="app-footer">
    <v-container>
      <v-row>
        <v-col cols="12" md="4">
          <h4 class="gradient-text text-subtitle-1 font-weight-bold mb-2 footer-heading">Universidad</h4>
          <p class="text-body-2 footer-text">Universidad Nacional de Colombia</p>
          <p class="text-body-2 footer-text">Sede Bogotá</p>
          <p class="text-body-2 footer-text">Facultad de ingeniería</p>
        </v-col>

        <v-col cols="12" md="4" class="text-center text-md-center">
          <img
              src="@/assets/icons/Logotipo_unal.svg"
              alt="Logotipo Universidad Nacional de Colombia"
              class="unal-logo"
          >
          <p class="text-caption footer-text">
            &copy; {{ currentYear }} SoftPharos. Proyecto académico UNAL.
          </p>
        </v-col>

        <v-col cols="12" md="4" class="text-right">
          <h4 class="gradient-text text-subtitle-1 font-weight-bold mb-2 footer-heading">Contacto</h4>
          <p class="text-body-2 footer-text">Oscar Eduardo Alvarez Rodriguez</p>
          <p class="text-body-2 footer-text">oalvarezr@unal.edu.co</p>
        </v-col>
      </v-row>

      <v-divider class="my-4 footer-divider"></v-divider>
    </v-container>
  </v-footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped>
.unal-logo {
  display: block;
  max-width: 100px;
  width: 100%;
  height: auto;
  margin-inline: auto;
}
</style>
