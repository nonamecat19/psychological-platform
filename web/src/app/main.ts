import { createApp } from 'vue'
import '@app/style.css'
import { App, appRouter } from '@app'
import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'
import { createPinia } from 'pinia'
import ToastService from 'primevue/toastservice'

const app = createApp(App)
const pinia = createPinia()

app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
})
app.use(appRouter)
app.use(pinia)
app.use(ToastService)

app.mount('#app')
