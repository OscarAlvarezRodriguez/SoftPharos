<template>
  <v-container class="fill-height">
    <v-row justify="center" align="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <v-card class="not-found-card pa-6 pa-md-8 text-center">
          <div class="not-found-code text-h2 text-md-h1 font-weight-bold mb-2">
            404
          </div>

          <div class="text-h6 text-md-h5 font-weight-bold mb-3" style="color: var(--color-text);">
            Página no encontrada
          </div>

          <p class="text-body-1 mb-6 not-found-subtitle">
            La ruta que intentas acceder no existe en SoftPharos.
            Es posible que el enlace esté roto o que la página haya sido movida.
          </p>

          <div class="d-flex flex-column flex-sm-row justify-center ga-3 mb-6">
            <v-btn
                color="primary"
                elevation="0"
                to="/"
                prepend-icon="mdi-home-outline"
                size="large"
            >
              Volver al inicio
            </v-btn>

            <v-btn
                variant="outlined"
                color="primary"
                elevation="0"
                size="large"
                @click="goBack"
            >
              Ir a la página anterior
            </v-btn>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push({ name: 'home' })
  }
}
</script>

<style scoped>
.not-found-card {
  background-color: var(--color-background);
  border: 1px solid var(--color-border);
}

.not-found-code {
  color: var(--color-primary);
  letter-spacing: 0.08em;
}

.not-found-subtitle {
  color: var(--color-text-secondary);
  line-height: 1.7;
}
</style>
