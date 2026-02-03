<template>
  <header class="h-9 bg-dark-925/95 backdrop-blur-xl border-b border-dark-850 relative" style="--wails-draggable:drag">
    <div class="h-full flex items-center px-2">
      <!-- Left: Logo (draggable) -->
      <div class="flex items-center gap-2 flex-shrink-0">
        <img src="/lumen-icon.png" alt="Lumen" class="w-5 h-5" draggable="false" />
        <span class="font-semibold text-sm text-white hidden md:inline">Lumen</span>
      </div>

      <!-- Center spacer (draggable) -->
      <div class="flex-1"></div>

      <!-- Center: Search Bar (absolute, non-draggable) -->
      <div class="absolute left-1/2 -translate-x-1/2 w-80 pointer-events-none">
        <div class="relative pointer-events-auto" style="--wails-draggable:no-drag">
          <input
            v-model="searchInput"
            type="text"
            placeholder="Поиск фильмов, сериалов..."
            autocomplete="off"
            autocorrect="off"
            autocapitalize="off"
            spellcheck="false"
            class="w-full h-7 pl-8 pr-4 bg-dark-850/90 border border-dark-700 rounded-lg
                   text-sm text-white placeholder-dark-500 focus:outline-none focus:border-accent-600
                   focus:ring-1 focus:ring-accent-600/20 transition-all duration-200"
            @keyup.enter="handleSearch"
          />
          <svg
            class="absolute left-2.5 top-1/2 -translate-y-1/2 w-4 h-4 text-dark-500 pointer-events-none"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            />
          </svg>
          <button
            v-if="searchInput"
            @click="clearSearch"
            class="absolute right-3 top-1/2 -translate-y-1/2 text-dark-400 hover:text-white"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Right spacer (draggable) -->
      <div class="flex-1"></div>

      <!-- Right: Window Controls (non-draggable) -->
      <div class="flex items-center gap-1 flex-shrink-0" style="--wails-draggable:no-drag">
        <!-- Settings Button -->
        <button
          @click="openSettings"
          class="w-8 h-8 flex items-center justify-center text-dark-500 hover:text-white hover:bg-dark-800 rounded-md transition-all duration-200"
          title="Настройки"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
          </svg>
        </button>
        
        <!-- Divider -->
        <div class="w-px h-4 bg-dark-800 mx-1"></div>
        
        <button
          @click="minimize"
          class="w-8 h-8 flex items-center justify-center text-dark-500 hover:text-white hover:bg-dark-800 rounded-md transition-all duration-200"
          title="Свернуть"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
          </svg>
        </button>
        <button
          @click="maximize"
          class="w-8 h-8 flex items-center justify-center text-dark-500 hover:text-white hover:bg-dark-800 rounded-md transition-all duration-200"
          title="Развернуть"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <rect x="4" y="4" width="16" height="16" rx="1" stroke-width="2" fill="none" />
          </svg>
        </button>
        <button
          @click="close"
          class="w-8 h-8 flex items-center justify-center text-dark-500 hover:text-white hover:bg-error-600 rounded-md transition-all duration-200"
          title="Закрыть"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </div>
    
    <!-- Settings Modal -->
    <SettingsModal :isOpen="showSettings" @close="showSettings = false" />
  </header>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import SettingsModal from './SettingsModal.vue'

const emit = defineEmits(['search'])
const router = useRouter()
const searchInput = ref('')
const showSettings = ref(false)

const openSettings = () => {
  showSettings.value = true
}

const handleSearch = () => {
  if (searchInput.value.trim()) {
    emit('search', searchInput.value.trim())
    router.push({ name: 'Search', query: { q: searchInput.value.trim() } })
  }
}

const clearSearch = () => {
  searchInput.value = ''
}

const minimize = () => {
  if (window.go?.main?.App?.WindowMinimize) {
    window.go.main.App.WindowMinimize()
  }
}

const maximize = () => {
  if (window.go?.main?.App?.WindowMaximize) {
    window.go.main.App.WindowMaximize()
  }
}

const close = () => {
  if (window.go?.main?.App?.WindowClose) {
    window.go.main.App.WindowClose()
  }
}
</script>
