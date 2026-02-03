<template>
  <div class="space-y-4">
    <div>
      <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
        <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"/>
        </svg>
        TorrServer
      </h3>
      <p class="text-sm text-dark-400 mb-6">Настройка подключения к серверу торрентов</p>
    </div>

    <div class="card p-4">
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
              @click="testConnection"
              class="btn-secondary flex items-center gap-2"
              :disabled="testing"
            >
              <svg v-if="testing" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
              </svg>
              Проверить
            </button>
          </div>
          <p v-if="status !== null" :class="['text-sm mt-2', status ? 'text-green-400' : 'text-red-400']">
            {{ status ? '✓ Подключено успешно' : '✗ Не удалось подключиться' }}
          </p>
        </div>

        <div>
          <label class="block text-sm font-medium text-dark-300 mb-2">Таймаут подключения (сек)</label>
          <input
            v-model.number="settings.timeout"
            type="number"
            min="5"
            max="60"
            class="input-field"
            placeholder="30"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const settings = reactive({
  torrServerURL: 'http://localhost:8090',
  timeout: 30
})

const testing = ref(false)
const status = ref(null)

const testConnection = async () => {
  testing.value = true
  status.value = null
  
  setTimeout(() => {
    status.value = true
    testing.value = false
  }, 1000)
}
</script>
