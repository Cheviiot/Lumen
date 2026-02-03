<template>
  <div class="space-y-4">
    <div>
      <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
        <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
        </svg>
        Приложение
      </h3>
      <p class="text-sm text-dark-400 mb-6">Общие настройки приложения</p>
    </div>

    <div class="card p-4">
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-dark-300 mb-2">Язык интерфейса</label>
          <select v-model="settings.language" class="input-field">
            <option value="ru">Русский</option>
            <option value="en">English</option>
            <option value="uk">Українська</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-dark-300 mb-2">Тема оформления</label>
          <select v-model="settings.theme" class="input-field">
            <option value="dark">Тёмная</option>
            <option value="light">Светлая</option>
            <option value="auto">Автоматически</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-dark-300 mb-2">Качество видео по умолчанию</label>
          <select v-model="settings.defaultQuality" class="input-field">
            <option value="auto">Автоматически</option>
            <option value="2160p">4K (2160p)</option>
            <option value="1080p">Full HD (1080p)</option>
            <option value="720p">HD (720p)</option>
            <option value="480p">SD (480p)</option>
          </select>
        </div>

        <div class="space-y-3 pt-2">
          <label class="flex items-center gap-3">
            <input
              v-model="settings.autoplay"
              type="checkbox"
              class="checkbox"
            />
            <div>
              <div class="text-sm font-medium text-white">Автовоспроизведение</div>
              <div class="text-xs text-dark-400">Автоматически начинать воспроизведение</div>
            </div>
          </label>
          
          <label class="flex items-center gap-3">
            <input
              v-model="settings.notifications"
              type="checkbox"
              class="checkbox"
            />
            <div>
              <div class="text-sm font-medium text-white">Уведомления</div>
              <div class="text-xs text-dark-400">Показывать системные уведомления</div>
            </div>
          </label>

          <label class="flex items-center gap-3">
            <input
              v-model="settings.minimizeToTray"
              type="checkbox"
              class="checkbox"
            />
            <div>
              <div class="text-sm font-medium text-white">Сворачивать в трей</div>
              <div class="text-xs text-dark-400">Сворачивать в системный трей при закрытии</div>
            </div>
          </label>
        </div>
      </div>
    </div>

    <div class="card p-4">
      <div class="space-y-4">
        <div>
          <h4 class="text-base font-semibold text-white mb-3 flex items-center gap-2">
            <svg class="w-4 h-4 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
            TMDB Прокси
          </h4>
          <p class="text-xs text-dark-500 mb-4">Использование прокси для доступа к TMDB API и изображениям через rootu.top</p>
        </div>

        <div class="space-y-3">
          <label class="flex items-center gap-3">
            <input
              v-model="settings.tmdbProxyEnabled"
              type="checkbox"
              class="checkbox"
            />
            <div>
              <div class="text-sm font-medium text-white">Включить TMDB Прокси</div>
              <div class="text-xs text-dark-400">Проксировать запросы к TMDB через rootu.top</div>
            </div>
          </label>

          <div v-if="settings.tmdbProxyEnabled" class="ml-9 space-y-3 pl-4 border-l-2 border-dark-800">
            <label class="flex items-center gap-3">
              <input
                v-model="settings.tmdbProxyApi"
                type="checkbox"
                class="checkbox"
              />
              <div>
                <div class="text-sm font-medium text-white">Проксировать API</div>
                <div class="text-xs text-dark-400">tmdbapi.rootu.top/3/</div>
              </div>
            </label>

            <label class="flex items-center gap-3">
              <input
                v-model="settings.tmdbProxyImages"
                type="checkbox"
                class="checkbox"
              />
              <div>
                <div class="text-sm font-medium text-white">Проксировать изображения</div>
                <div class="text-xs text-dark-400">tmdbimg.rootu.top/</div>
              </div>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, watch } from 'vue'

const settings = reactive({
  language: localStorage.getItem('app_language') || 'ru',
  theme: localStorage.getItem('app_theme') || 'dark',
  defaultQuality: localStorage.getItem('app_quality') || 'auto',
  autoplay: localStorage.getItem('app_autoplay') !== 'false',
  notifications: localStorage.getItem('app_notifications') !== 'false',
  minimizeToTray: localStorage.getItem('app_minimize_tray') === 'true',
  tmdbProxyEnabled: localStorage.getItem('tmdb_proxy_enabled') !== 'false', // Включено по умолчанию
  tmdbProxyApi: localStorage.getItem('tmdb_proxy_api') !== 'false',
  tmdbProxyImages: localStorage.getItem('tmdb_proxy_images') !== 'false'
})

// Сохранение настроек в localStorage при изменении
watch(() => settings.language, (val) => localStorage.setItem('app_language', val))
watch(() => settings.theme, (val) => localStorage.setItem('app_theme', val))
watch(() => settings.defaultQuality, (val) => localStorage.setItem('app_quality', val))
watch(() => settings.autoplay, (val) => localStorage.setItem('app_autoplay', val))
watch(() => settings.notifications, (val) => localStorage.setItem('app_notifications', val))
watch(() => settings.minimizeToTray, (val) => localStorage.setItem('app_minimize_tray', val))

// Обновление настроек TMDB прокси
watch(() => settings.tmdbProxyEnabled, (val) => {
  localStorage.setItem('tmdb_proxy_enabled', val)
  if (window.go?.main?.TMDB?.UpdateProxySettings) {
    window.go.main.TMDB.UpdateProxySettings(
      val && settings.tmdbProxyApi,
      val && settings.tmdbProxyImages
    )
  }
})

watch(() => settings.tmdbProxyApi, (val) => {
  localStorage.setItem('tmdb_proxy_api', val)
  if (window.go?.main?.TMDB?.UpdateProxySettings && settings.tmdbProxyEnabled) {
    window.go.main.TMDB.UpdateProxySettings(val, settings.tmdbProxyImages)
  }
})

watch(() => settings.tmdbProxyImages, (val) => {
  localStorage.setItem('tmdb_proxy_images', val)
  if (window.go?.main?.TMDB?.UpdateProxySettings && settings.tmdbProxyEnabled) {
    window.go.main.TMDB.UpdateProxySettings(settings.tmdbProxyApi, val)
  }
})
</script>
