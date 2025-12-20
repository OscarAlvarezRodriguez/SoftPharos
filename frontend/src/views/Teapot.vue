<template>
  <v-container class="fill-height">
    <v-row justify="center" align="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <v-card class="teapot-card pa-6 pa-md-8 text-center">
          <div class="teapot-code text-h2 text-md-h1 font-weight-bold mb-2">
            418
          </div>

          <div class="text-h6 text-md-h5 font-weight-bold mb-3 teapot-title">
            Soy una tetera ☕
          </div>

          <div class="mb-6 teapot-icon">
            <v-icon
                icon="mdi-tea-outline"
                size="64"
            />
          </div>

          <p class="text-body-1 mb-6 teapot-subtitle">
            Este endpoint no sirve café, ni responde como una API normal.
            Solo soy una humilde tetera HTTP que se niega educadamente a procesar tu petición.
          </p>

          <div class="d-flex flex-column flex-sm-row justify-center ga-3 mb-4">
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
              Intentar otra ruta
            </v-btn>
          </div>

          <p class="text-caption teapot-hint">
            HTTP 418 – I&apos;m a teapot.
            Easter egg para quienes exploran más allá de las rutas esperadas.
          </p>
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
.teapot-card {
  background-color: var(--color-background);
  border: 1px solid var(--color-border);
}

.teapot-code {
  color: var(--color-primary);
  letter-spacing: 0.08em;
}

.teapot-title {
  color: var(--color-text);
}

.teapot-subtitle {
  color: var(--color-text-secondary);
  line-height: 1.7;
}

.teapot-hint {
  color: var(--color-text-secondary);
  white-space: pre-line;
}

.teapot-icon {
  color: var(--color-primary);
}
</style>
