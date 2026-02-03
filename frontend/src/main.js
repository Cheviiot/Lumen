import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './styles/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')

// Применить настройки TMDB прокси при запуске
if (window.go?.main?.TMDB?.UpdateProxySettings) {
  const proxyEnabled = localStorage.getItem('tmdb_proxy_enabled') !== 'false' // По умолчанию включено
  const proxyAPI = localStorage.getItem('tmdb_proxy_api') !== 'false'
  const proxyImages = localStorage.getItem('tmdb_proxy_images') !== 'false'
  
  if (proxyEnabled) {
    window.go.main.TMDB.UpdateProxySettings(proxyAPI, proxyImages)
    console.log('TMDB Proxy enabled:', { api: proxyAPI, images: proxyImages })
  }
}
