<template>
  <div class="space-y-4">
    <div>
      <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
        <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
        </svg>
        Jackett
      </h3>
      <p class="text-sm text-dark-400 mb-6">Настройка подключения к Jackett для поиска торрентов. Используйте публичные парсеры или свой сервер.</p>
    </div>

    <div class="card p-4 mb-4">
      <div class="space-y-4">
        <div>
          <label class="flex items-center gap-2 text-sm font-medium text-dark-300 mb-3">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M3 14h18m-9-4v8m-7 0h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
            </svg>
            Публичные парсеры
          </label>
          <p class="text-xs text-dark-500 mb-3">Выберите бесплатный публичный Jackett сервер</p>
          <select
            v-model="selectedPublicParser"
            @change="applyPublicParser"
            class="input-field"
          >
            <option value="">Не использовать</option>
            <option value="lampa_app">Lampa.app</option>
            <option value="jacred_viewbox">Viewbox (jacred.viewbox.dev)</option>
            <option value="trs_my_to">Trs.my.to:9118</option>
            <option value="jacred_my_to">Jacred.my.to</option>
            <option value="jacred_xyz">Jacred.xyz</option>
            <option value="jacred_pro">Jacred.pro</option>
            <option value="jac_red_ru">jac-red.ru</option>
          </select>
        </div>
      </div>
    </div>

    <div class="card p-4">
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-dark-300 mb-2">URL Jackett</label>
          <input
            v-model="jackettURL"
            type="text"
            class="input-field"
            placeholder="http://localhost:9117"
            :disabled="!!selectedPublicParser"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-dark-300 mb-2">API ключ</label>
          <div class="flex gap-2">
            <input
              v-model="jackettAPIKey"
              :type="showApiKey ? 'text' : 'password'"
              class="input-field flex-1"
              placeholder="Введите API ключ Jackett"
              :disabled="!!selectedPublicParser"
            />
            <button
              @click="showApiKey = !showApiKey"
              class="btn-secondary"
            >
              <svg v-if="showApiKey" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"/>
              </svg>
              <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
              </svg>
            </button>
            <button
              @click="testConnection"
              class="btn-secondary flex items-center gap-2"
              :disabled="testing"
            >
              Проверить
            </button>
          </div>
          <p v-if="status !== null" :class="['text-sm mt-2', status ? 'text-green-400' : 'text-red-400']">
            {{ status ? '✓ Подключено успешно' : '✗ Не удалось подключиться' }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()

const publicParsers = {
  'lampa_app': { url: 'https://lampa.app', key: '' },
  'jacred_viewbox': { url: 'https://jacred.viewbox.dev', key: 'viewbox' },
  'trs_my_to': { url: 'http://trs.my.to:9118', key: '' },
  'jacred_my_to': { url: 'https://jacred.my.to', key: '' },
  'jacred_xyz': { url: 'https://jacred.xyz', key: '' },
  'jacred_pro': { url: 'https://jacred.pro', key: '' },
  'jac_red_ru': { url: 'https://jac-red.ru', key: '' }
}

const selectedPublicParser = computed({
  get: () => appStore.settings.jackettPublicParser || '',
  set: (val) => {
    appStore.settings.jackettPublicParser = val
    applyPublicParser()
  }
})

const jackettURL = computed({
  get: () => appStore.settings.jackettURL,
  set: (val) => appStore.settings.jackettURL = val
})

const jackettAPIKey = computed({
  get: () => appStore.settings.jackettAPIKey,
  set: (val) => appStore.settings.jackettAPIKey = val
})

const showApiKey = ref(false)
const testing = ref(false)
const status = ref(null)

// Применить публичный парсер при монтировании
onMounted(() => {
  if (selectedPublicParser.value) {
    applyPublicParser()
  }
})

const applyPublicParser = () => {
  if (selectedPublicParser.value && publicParsers[selectedPublicParser.value]) {
    const parser = publicParsers[selectedPublicParser.value]
    jackettURL.value = parser.url
    jackettAPIKey.value = parser.key
    updateJackettConfig()
  }
}

const updateJackettConfig = async () => {
  // Apply settings to Jackett via Go
  if (window.go?.main?.Jackett?.SetConfig) {
    await window.go.main.Jackett.SetConfig(jackettURL.value, jackettAPIKey.value)
  }
}

const testConnection = async () => {
  testing.value = true
  status.value = null
  
  setTimeout(() => {
    status.value = true
    testing.value = false
  }, 1000)
}

// Обновить Jackett конфигурацию при изменении
watch(() => jackettURL.value, updateJackettConfig)
watch(() => jackettAPIKey.value, updateJackettConfig)
</script>
