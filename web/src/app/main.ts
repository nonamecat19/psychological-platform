import { createApp } from 'vue'
import '@app/style.css'
import { App, appRouter } from '@app'
import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'

const app = createApp(App)

app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
})
app.use(appRouter)

app.mount('#app')
