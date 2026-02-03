<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="isOpen"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="closeModal"
      >
        <!-- Backdrop -->
        <div class="absolute inset-0 bg-black/90 backdrop-blur-md"></div>
        
        <!-- Modal -->
        <div class="relative bg-dark-925/95 backdrop-blur-xl rounded-2xl shadow-2xl border border-dark-850 w-full max-w-5xl max-h-[85vh] overflow-hidden flex flex-col">
          <!-- Header -->
          <div class="flex items-center justify-between px-6 py-4 border-b border-dark-850 flex-shrink-0">
            <h2 class="text-xl font-bold text-white flex items-center gap-2">
              <svg class="w-6 h-6 text-accent-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
              </svg>
              Настройки
            </h2>
            <button
              @click="closeModal"
              class="w-8 h-8 flex items-center justify-center text-dark-500 hover:text-white hover:bg-dark-800 rounded-lg transition-all duration-200"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>

          <!-- Content with Sidebar -->
          <div class="flex flex-1 overflow-hidden">
            <!-- Sidebar -->
            <div class="w-56 border-r border-dark-850 bg-dark-950/30 overflow-y-auto flex-shrink-0">
              <nav class="p-3 space-y-1">
                <button
                  v-for="category in categories"
                  :key="category.id"
                  @click="activeCategory = category.id"
                  :class="[
                    'w-full flex items-center gap-3 px-3 py-2.5 rounded-lg transition-all duration-200 text-left border',
                    activeCategory === category.id
                      ? 'bg-accent-600/10 text-accent-500 border-accent-600/20'
                      : 'text-dark-500 hover:bg-dark-800 hover:text-white border-transparent'
                  ]"
                >
                >
                  <component :is="category.icon" class="w-4 h-4 flex-shrink-0" />
                  <span class="text-sm font-medium">{{ category.label }}</span>
                </button>
              </nav>
            </div>

            <!-- Content Area -->
            <div class="flex-1 overflow-y-auto p-6">
              <component :is="currentComponent" />
            </div>
          </div>

          <!-- Footer -->
          <div class="flex justify-end gap-3 px-6 py-4 border-t border-dark-850 bg-dark-900/30 flex-shrink-0">
            <button @click="resetSettings" class="btn-secondary text-sm px-4 py-2">
              Сбросить
            </button>
            <button @click="saveSettings" class="btn-primary text-sm px-4 py-2" :disabled="saving">
              {{ saving ? 'Сохранение...' : 'Сохранить настройки' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, computed, watch, h } from 'vue'
import { useAppStore } from '@/stores/app'
import TorrServerSettings from './settings/TorrServerSettings.vue'
import JackettSettings from './settings/JackettSettings.vue'
import SupabaseSettings from './settings/SupabaseSettings.vue'
import AppSettings from './settings/AppSettings.vue'

const appStore = useAppStore()

const props = defineProps({
  isOpen: Boolean
})

const emit = defineEmits(['close'])

const activeCategory = ref('torrserver')
const saving = ref(false)

// Icons
const ServerIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01'
      })
    ])
  }
}

const SearchIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z'
      })
    ])
  }
}

const CloudIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z'
      })
    ])
  }
}

const SettingsIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z'
      })
    ])
  }
}

const categories = [
  { id: 'torrserver', label: 'TorrServer', icon: ServerIcon, component: TorrServerSettings },
  { id: 'jackett', label: 'Jackett', icon: SearchIcon, component: JackettSettings },
  { id: 'supabase', label: 'Supabase', icon: CloudIcon, component: SupabaseSettings },
  { id: 'app', label: 'Приложение', icon: SettingsIcon, component: AppSettings }
]

const currentComponent = computed(() => {
  const category = categories.find(c => c.id === activeCategory.value)
  return category?.component || TorrServerSettings
})

const closeModal = () => {
  emit('close')
}

const saveSettings = async () => {
  saving.value = true
  
  try {
    await appStore.saveSettings(appStore.settings)
    console.log('Settings saved from modal')
    setTimeout(() => {
      saving.value = false
      closeModal()
    }, 300)
  } catch (error) {
    console.error('Failed to save settings:', error)
    saving.value = false
  }
}

const resetSettings = () => {
  // Logic to reset settings
}

// Close on Escape
watch(() => props.isOpen, (newVal) => {
  if (newVal) {
    const handleEscape = (e) => {
      if (e.key === 'Escape') {
        closeModal()
      }
    }
    document.addEventListener('keydown', handleEscape)
    return () => document.removeEventListener('keydown', handleEscape)
  }
})
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-active > div:last-child,
.modal-leave-active > div:last-child {
  transition: transform 0.2s ease, opacity 0.2s ease;
}

.modal-enter-from > div:last-child,
.modal-leave-to > div:last-child {
  transform: scale(0.95);
  opacity: 0;
}
</style>
