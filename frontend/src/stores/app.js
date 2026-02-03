import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const isLoading = ref(false)
  const searchQuery = ref('')
  const sidebarCollapsed = ref(true) // По умолчанию свёрнут
  const settings = ref({
    torrServerURL: 'http://localhost:8090',
    jackettURL: 'http://localhost:9117',
    jackettAPIKey: '',
    jackettPublicParser: '',
    supabaseURL: '',
    supabaseKey: '',
    theme: 'dark',
    language: 'ru',
    autoSync: false
  })

  async function init() {
    isLoading.value = true
    try {
      // Initialize database
      if (window.go?.main?.Database?.Init) {
        await window.go.main.Database.Init()
        console.log('Database initialized')
      }
      
      // Load settings
      if (window.go?.main?.Database?.GetSettings) {
        const loadedSettings = await window.go.main.Database.GetSettings()
        console.log('Loaded settings:', JSON.stringify(loadedSettings, null, 2))
        if (loadedSettings) {
          settings.value = { ...settings.value, ...loadedSettings }
          console.log('Applied settings:', JSON.stringify(settings.value, null, 2))
        }
      }
    } catch (error) {
      console.error('Failed to initialize app:', error)
    } finally {
      isLoading.value = false
    }
  }

  function setSearchQuery(query) {
    searchQuery.value = query
  }

  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  async function saveSettings(newSettings) {
    settings.value = { ...settings.value, ...newSettings }
    console.log('Saving settings:', JSON.stringify(settings.value, null, 2))
    if (window.go?.main?.Database?.SaveSettings) {
      await window.go.main.Database.SaveSettings(settings.value)
      console.log('Settings saved successfully')
    } else {
      console.error('SaveSettings function not available')
    }
  }

  return {
    isLoading,
    searchQuery,
    sidebarCollapsed,
    settings,
    init,
    setSearchQuery,
    toggleSidebar,
    saveSettings
  }
})
