<template>
  <div class="p-6 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold text-white mb-6">Настройки</h1>

    <div class="space-y-6">
      <!-- TorrServer Settings -->
      <div class="card p-6">
        <h2 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
          <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"/>
          </svg>
          TorrServer
        </h2>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">URL TorrServer</label>
            <div class="flex gap-2">
              <input
                v-model="settings.torrServerURL"
                type="text"
                class="input-field flex-1"
                placeholder="http://localhost:8090"
              />
              <button
                @click="testTorrServer"
                class="btn-secondary flex items-center gap-2"
                :disabled="testingTorrServer"
              >
                <svg v-if="testingTorrServer" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
                </svg>
                Проверить
              </button>
            </div>
            <p v-if="torrServerStatus !== null" :class="['text-sm mt-1', torrServerStatus ? 'text-green-400' : 'text-red-400']">
              {{ torrServerStatus ? '✓ Подключено' : '✗ Не удалось подключиться' }}
            </p>
          </div>
        </div>
      </div>

      <!-- Jackett Settings -->
      <div class="card p-6">
        <h2 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
          <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
          </svg>
          Jackett
        </h2>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">URL Jackett</label>
            <input
              v-model="settings.jackettURL"
              type="text"
              class="input-field"
              placeholder="http://localhost:9117"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">API ключ</label>
            <div class="flex gap-2">
              <input
                v-model="settings.jackettAPIKey"
                :type="showApiKey ? 'text' : 'password'"
                class="input-field flex-1"
                placeholder="Введите API ключ Jackett"
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
                @click="testJackett"
                class="btn-secondary flex items-center gap-2"
                :disabled="testingJackett"
              >
                Проверить
              </button>
            </div>
            <p v-if="jackettStatus !== null" :class="['text-sm mt-1', jackettStatus ? 'text-green-400' : 'text-red-400']">
              {{ jackettStatus ? '✓ Подключено' : '✗ Не удалось подключиться' }}
            </p>
          </div>
        </div>
      </div>

      <!-- Supabase Settings -->
      <div class="card p-6">
        <h2 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
          <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z"/>
          </svg>
          Supabase (Синхронизация)
        </h2>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">URL проекта Supabase</label>
            <input
              v-model="settings.supabaseURL"
              type="text"
              class="input-field"
              placeholder="https://your-project.supabase.co"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">API ключ (anon)</label>
            <input
              v-model="settings.supabaseKey"
              type="password"
              class="input-field"
              placeholder="Введите anon ключ Supabase"
            />
          </div>
          
          <div class="flex items-center gap-3">
            <input
              v-model="settings.autoSync"
              type="checkbox"
              id="autoSync"
              class="w-4 h-4 rounded border-dark-600 bg-dark-700 text-accent-600 focus:ring-accent-500"
            />
            <label for="autoSync" class="text-sm text-dark-300">Автоматическая синхронизация</label>
          </div>
        </div>
      </div>

      <!-- App Settings -->
      <div class="card p-6">
        <h2 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
          <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
          </svg>
          Приложение
        </h2>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">Язык</label>
            <select v-model="settings.language" class="input-field">
              <option value="ru">Русский</option>
              <option value="en">English</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-dark-300 mb-2">Тема</label>
            <select v-model="settings.theme" class="input-field">
              <option value="dark">Тёмная</option>
              <option value="light">Светлая</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Save Button -->
      <div class="flex justify-end gap-3">
        <button @click="resetSettings" class="btn-secondary">
          Сбросить
        </button>
        <button @click="saveSettings" class="btn-primary" :disabled="saving">
          {{ saving ? 'Сохранение...' : 'Сохранить настройки' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useAppStore } from '../stores/app'
import { useTorrentStore } from '../stores/torrent'

const appStore = useAppStore()
const torrentStore = useTorrentStore()

const settings = reactive({
  torrServerURL: 'http://localhost:8090',
  jackettURL: 'http://localhost:9117',
  jackettAPIKey: '',
  supabaseURL: '',
  supabaseKey: '',
  theme: 'dark',
  language: 'ru',
  autoSync: false
})

const saving = ref(false)
const showApiKey = ref(false)
const testingTorrServer = ref(false)
const testingJackett = ref(false)
const torrServerStatus = ref(null)
const jackettStatus = ref(null)

const testTorrServer = async () => {
  testingTorrServer.value = true
  torrServerStatus.value = null
  
  try {
    torrServerStatus.value = await torrentStore.checkTorrServer()
  } catch (error) {
    torrServerStatus.value = false
  } finally {
    testingTorrServer.value = false
  }
}

const testJackett = async () => {
  testingJackett.value = true
  jackettStatus.value = null
  
  try {
    jackettStatus.value = await torrentStore.testJackettConnection()
  } catch (error) {
    jackettStatus.value = false
  } finally {
    testingJackett.value = false
  }
}

const saveSettings = async () => {
  saving.value = true
  try {
    await appStore.saveSettings(settings)
  } finally {
    saving.value = false
  }
}

const resetSettings = () => {
  Object.assign(settings, {
    torrServerURL: 'http://localhost:8090',
    jackettURL: 'http://localhost:9117',
    jackettAPIKey: '',
    supabaseURL: '',
    supabaseKey: '',
    theme: 'dark',
    language: 'ru',
    autoSync: false
  })
}

onMounted(() => {
  Object.assign(settings, appStore.settings)
})
</script>
