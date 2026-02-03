<template>
  <div class="space-y-4">
    <div>
      <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
        <svg class="w-5 h-5 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z"/>
        </svg>
        Supabase
      </h3>
      <p class="text-sm text-dark-400 mb-6">Синхронизация данных между устройствами</p>
    </div>

    <div class="card p-4">
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
        
        <div class="space-y-3 pt-2">
          <label class="flex items-center gap-3">
            <input
              v-model="settings.autoSync"
              type="checkbox"
              class="checkbox"
            />
            <div>
              <div class="text-sm font-medium text-white">Автоматическая синхронизация</div>
              <div class="text-xs text-dark-400">Синхронизировать при каждом изменении</div>
            </div>
          </label>
          
          <label class="flex items-center gap-3">
            <input
              v-model="settings.syncOnStartup"
              type="checkbox"
              class="checkbox"
            />
            <div>
              <div class="text-sm font-medium text-white">Синхронизация при запуске</div>
              <div class="text-xs text-dark-400">Загружать данные при старте приложения</div>
            </div>
          </label>
        </div>

        <div class="pt-2">
          <button class="btn-secondary w-full">
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
            </svg>
            Синхронизировать сейчас
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'

const settings = reactive({
  supabaseURL: '',
  supabaseKey: '',
  autoSync: false,
  syncOnStartup: true
})
</script>
