<template>
  <div class="hero-background py-12 py-md-16">
    <v-container>
      <v-row justify="center">
        <v-col cols="12" md="10" lg="8" class="text-center">
          <h1 class="text-h3 text-md-h2 font-weight-bold mb-6" style="color: var(--color-text); line-height: 1.3;">
            La ingeniería de software <span class="gradient-text">NO</span> es programar, es diseñar soluciones.
          </h1>

          <p class="text-body-1 text-md-h6 mb-10 mx-auto" style="color: var(--color-text-secondary); max-width: 700px; line-height: 1.7; font-weight: 400;">
            SoftPharos es la plataforma educativa que permite a los estudiantes visualizar
            y documentar el desarrollo iterativo de sus proyectos de software, aprendiendo
            de cada decisión.
          </p>

          <div class="d-flex flex-column flex-sm-row justify-center ga-3 mb-8">
            <v-btn
              to="/register"
              color="primary"
              size="large"
              elevation="0"
              prepend-icon="mdi-account-plus"
            >
              Comenzar Ahora
            </v-btn>

            <v-btn
              to="/login"
              variant="outlined"
              color="primary"
              size="large"
              elevation="0"
            >
              Ya tengo cuenta
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
</script>
