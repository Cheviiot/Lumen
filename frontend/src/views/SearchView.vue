<template>
  <div class="p-6">
    <h1 class="text-2xl font-bold text-white mb-6">
      Результаты поиска: "{{ searchQuery }}"
    </h1>

    <!-- Tabs -->
    <div class="flex gap-2 mb-6">
      <button
        :class="[
          'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
          activeTab === 'movies' ? 'bg-accent-600 text-white' : 'bg-dark-800 text-dark-400 hover:text-white'
        ]"
        @click="activeTab = 'movies'"
      >
        Фильмы ({{ movieResults.length }})
      </button>
      <button
        :class="[
          'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
          activeTab === 'series' ? 'bg-accent-600 text-white' : 'bg-dark-800 text-dark-400 hover:text-white'
        ]"
        @click="activeTab = 'series'"
      >
        Сериалы ({{ tvResults.length }})
      </button>
    </div>

    <LoadingSpinner v-if="loading" text="Поиск..." />

    <div v-else-if="currentItems.length === 0" class="flex flex-col items-center justify-center h-64 text-center">
      <svg class="w-16 h-16 text-dark-600 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
      </svg>
      <h3 class="text-lg font-medium text-dark-400 mb-2">Ничего не найдено</h3>
      <p class="text-dark-500 text-sm">Попробуйте изменить поисковый запрос</p>
    </div>

    <MediaGrid
      v-else
      :items="currentItems"
      :type="activeTab === 'movies' ? 'movie' : 'tv'"
    />
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useMediaStore } from '../stores/media'
import MediaGrid from '../components/MediaGrid.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'

const route = useRoute()
const mediaStore = useMediaStore()

const activeTab = ref('movies')
const loading = ref(false)
const movieResults = ref([])
const tvResults = ref([])

const searchQuery = computed(() => route.query.q || '')

const currentItems = computed(() => {
  return activeTab.value === 'movies' ? movieResults.value : tvResults.value
})

const performSearch = async (query) => {
  if (!query) return
  
  loading.value = true
  
  try {
    const [movies, tv] = await Promise.all([
      mediaStore.searchMovies(query, 1),
      mediaStore.searchTV(query, 1)
    ])
    
    movieResults.value = movies?.results || []
    tvResults.value = tv?.results || []
  } catch (error) {
    console.error('Search failed:', error)
  } finally {
    loading.value = false
  }
}

watch(() => route.query.q, (newQuery) => {
  if (newQuery) {
    performSearch(newQuery)
  }
})

onMounted(() => {
  if (searchQuery.value) {
    performSearch(searchQuery.value)
  }
})
</script>
