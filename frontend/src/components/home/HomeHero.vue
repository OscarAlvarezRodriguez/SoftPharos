<template>
  <div class="hero-background py-12 py-md-16">
    <v-container>
      <v-row justify="center">
        <v-col cols="12" md="10" lg="8" class="text-center">
          <h1 class="text-h3 text-md-h2 font-weight-bold mb-6">
            La ingeniería de software <span class="gradient-text">NO</span> es programar, es diseñar soluciones.
          </h1>

          <p class="text-body-1 text-md-h6 mb-10 hero-subtitle">
            SoftPharos es una plataforma educativa diseñada especifica para ingeniería de software 1
            en la Universidad Nacional de Colombia que permite a los estudiantes visualizar el avance y documentar el
            proceso de desarrollo durante el curso en sus proyectos de software.
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
    <v-container>
      <v-row justify="center" class="mb-10">
        <v-col cols="12" md="10" lg="8" class="text-center">
          <h2 class="text-h4 text-md-h3 font-weight-bold mb-4 section-title">
            ¿Qué hace especial a SoftPharos?
          </h2>
          <p class="text-body-1 section-subtitle">
            Una plataforma diseñada específicamente para el seguimiento y aprendizaje de mis estudiantes
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
              elevation="0"
              class="feature-card h-100 pa-2"
          >
            <v-card-text class="pa-6">
              <div class="text-h3 mb-4">{{ feature.icon }}</div>
              <h3 class="text-h6 font-weight-bold mb-3 feature-title">
                {{ feature.title }}
              </h3>
              <p class="text-body-2 feature-text">
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

<style scoped>
.hero-subtitle {
  color: var(--color-text-secondary);
  max-width: 700px;
  margin: 0 auto;
  line-height: 1.7;
  font-weight: 400;
}
</style>
