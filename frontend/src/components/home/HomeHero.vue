<template>
  <div class="hero-background py-16">
    <v-container>
      <v-row justify="center">
        <v-col cols="12" class="text-center">
          <h1 class="text-h2 text-sm-h1 font-weight-bold mb-4" style="color: var(--color-text);">
            Documenta el <span class="gradient-text">proceso</span>,<br>
            no solo el resultado
          </h1>

          <p class="text-h6 text-sm-h5 mb-8 mx-auto" style="color: var(--color-text-secondary); max-width: 800px;">
            SoftPharos es la plataforma educativa que permite a los estudiantes visualizar
            y documentar el desarrollo iterativo de sus proyectos de software, aprendiendo
            de cada decisión.
          </p>

          <div class="d-flex flex-column flex-sm-row justify-center ga-4 mb-12">
            <v-btn
              to="/register"
              color="primary"
              size="x-large"
              prepend-icon="mdi-account-plus"
            >
              Comenzar Ahora
            </v-btn>

            <v-btn
              to="/login"
              variant="outlined"
              color="primary"
              size="x-large"
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
