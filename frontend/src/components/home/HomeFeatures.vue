<template>
  <div class="py-16" style="background-color: var(--color-background);">
    <v-container>
      <v-row justify="center" class="mb-12">
        <v-col cols="12" class="text-center">
          <h2 class="text-h3 font-weight-bold mb-4" style="color: var(--color-text);">
            ¿Qué hace especial a SoftPharos?
          </h2>
          <p class="text-h6" style="color: var(--color-text-secondary);">
            Una plataforma diseñada específicamente para el aprendizaje iterativo
          </p>
        </v-col>
      </v-row>

      <v-row>
        <v-col
          v-for="feature in features"
          :key="feature.title"
          cols="12"
          md="6"
          lg="4"
        >
          <v-card
            elevation="2"
            class="h-100"
            style="background-color: var(--color-background-soft); border: 1px solid var(--color-border);"
          >
            <v-card-text class="pa-6">
              <div class="text-h2 mb-4">{{ feature.icon }}</div>
              <h3 class="text-h5 font-weight-bold mb-3" style="color: var(--color-text);">
                {{ feature.title }}
              </h3>
              <p style="color: var(--color-text-secondary);">
                {{ feature.description }}
              </p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
const features = [
  {
    icon: '📊',
    title: 'Gestión de Proyectos',
    description: 'Organiza y documenta tus proyectos de software con milestones y entregas semanales'
  },
  {
    icon: '💬',
    title: 'Colaboración en Equipo',
    description: 'Comenta, discute y aprende junto a tu equipo en cada etapa del desarrollo'
  },
  {
    icon: '✍️',
    title: 'Feedback Docente',
    description: 'Recibe retroalimentación continua de tus profesores para mejorar tu proceso'
  },
  {
    icon: '📈',
    title: 'Visualiza tu Evolución',
    description: 'Observa cómo tu proyecto y habilidades crecen semana a semana'
  },
  {
    icon: '🔄',
    title: 'Proceso Iterativo',
    description: 'Aprende que el desarrollo es un proceso continuo de mejora y adaptación'
  },
  {
    icon: '🎯',
    title: 'Objetivos Claros',
    description: 'Define y alcanza metas concretas en cada milestone de tu proyecto'
  }
]
</script>
