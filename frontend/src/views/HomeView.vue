<template>
  <v-container class="fill-height">
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <v-card elevation="4">
          <v-card-title class="text-h4 text-center pa-6">
            Hola Mundo desde Vuetify
          </v-card-title>

          <v-card-text class="text-center">
            <p class="text-h6 mb-4">SoftPharos</p>
            <p class="text-body-1 mb-6">
              Plataforma educativa para visualizar y documentar el proceso
              de desarrollo de proyectos de software
            </p>

            <v-btn
              color="primary"
              size="large"
              prepend-icon="mdi-rocket-launch"
              @click="handleClick"
            >
              Comenzar
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
function handleClick() {
  console.log('¡Botón Vuetify funcionando!')
}
</script>
