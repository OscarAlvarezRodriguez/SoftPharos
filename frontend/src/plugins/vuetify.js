import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors: {
          primary: '#3498db',
          secondary: '#9b59b6',
          accent: '#2980b9',
          error: '#e74c3c',
          info: '#3498db',
          success: '#27ae60',
          warning: '#f39c12'
        }
      },
      dark: {
        colors: {
          primary: '#5dade2',
          secondary: '#bb8fce',
          accent: '#3498db',
          error: '#e74c3c',
          info: '#5dade2',
          success: '#2ecc71',
          warning: '#f39c12'
        }
      }
    }
  }
})

export default vuetify
