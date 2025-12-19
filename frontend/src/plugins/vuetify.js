import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: localStorage.getItem('theme') || 'light',
    themes: {
      light: {
        dark: false,
        colors: {
          primary: '#2196F3',
          secondary: '#607D8B',
          accent: '#1976D2',
          error: '#E53935',
          info: '#2196F3',
          success: '#43A047',
          warning: '#FB8C00',
          background: '#fafafa',
          surface: '#ffffff'
        }
      },
      dark: {
        dark: true,
        colors: {
          primary: '#42A5F5',
          secondary: '#78909C',
          accent: '#2196F3',
          error: '#EF5350',
          info: '#42A5F5',
          success: '#66BB6A',
          warning: '#FFA726',
          background: '#212121',
          surface: '#2c2c2c'
        }
      }
    }
  }
})

export default vuetify
