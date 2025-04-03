import './assets/index.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { PiniaColada } from '@pinia/colada'

const app = createApp(App)

app.use(createPinia())
app.use(PiniaColada)
app.use(router)

app.mount('#app')
