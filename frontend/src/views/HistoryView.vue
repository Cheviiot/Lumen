<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-white">История просмотров</h1>
      
      <button
        v-if="history.length > 0"
        @click="clearAllHistory"
        class="btn-secondary text-sm flex items-center gap-2"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
        </svg>
        Очистить историю
      </button>
    </div>

    <LoadingSpinner v-if="loading" text="Загрузка истории..." />

    <div v-else-if="history.length === 0" class="flex flex-col items-center justify-center h-64 text-center">
      <svg class="w-16 h-16 text-dark-600 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
      </svg>
      <h3 class="text-lg font-medium text-dark-400 mb-2">История пуста</h3>
      <p class="text-dark-500 text-sm">Здесь будут отображаться просмотренные фильмы и сериалы</p>
    </div>

    <div v-else class="space-y-4">
      <div
        v-for="item in history"
        :key="`${item.media_type}-${item.tmdb_id}`"
        class="bg-dark-800 rounded-xl p-4 flex gap-4 hover:bg-dark-700/80 transition-colors cursor-pointer"
        @click="openDetails(item)"
      >
        <img
          :src="getImageURL(item.poster)"
          :alt="item.title"
          class="w-20 h-28 object-cover rounded-lg flex-shrink-0"
          draggable="false"
        />
        <div class="flex-1 min-w-0">
          <h3 class="font-medium text-white text-lg mb-1">{{ item.title }}</h3>
          <p v-if="item.episode" class="text-dark-400 text-sm mb-2">{{ item.episode }}</p>
          
          <!-- Progress bar -->
          <div class="mt-2">
            <div class="flex items-center justify-between text-xs text-dark-400 mb-1">
              <span>Прогресс</span>
              <span>{{ Math.round(item.progress) }}%</span>
            </div>
            <div class="h-1.5 bg-dark-700 rounded-full overflow-hidden">
              <div
                class="h-full bg-accent-500 rounded-full transition-all"
                :style="{ width: `${item.progress}%` }"
              />
            </div>
          </div>
          
          <p class="text-dark-500 text-xs mt-2">
            {{ formatDate(item.last_watched) }}
          </p>
        </div>
        
        <button
          @click.stop="continueWatching(item)"
          class="flex-shrink-0 self-center p-3 bg-accent-600 hover:bg-accent-700 rounded-full transition-colors"
        >
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useMediaStore } from '../stores/media'
import LoadingSpinner from '../components/LoadingSpinner.vue'

const router = useRouter()
const mediaStore = useMediaStore()

const history = ref([])
const loading = ref(true)

const getImageURL = (path) => {
  return mediaStore.getImageURL(path, 'w185')
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const openDetails = (item) => {
  const routeName = item.media_type === 'tv' ? 'TVDetails' : 'MovieDetails'
  router.push({ name: routeName, params: { id: item.tmdb_id } })
}

const continueWatching = (item) => {
  router.push({
    name: 'Player',
    query: {
      tmdb_id: item.tmdb_id,
      media_type: item.media_type,
      title: item.title
    }
  })
}

const clearAllHistory = async () => {
  await mediaStore.clearHistory()
  history.value = []
}

onMounted(async () => {
  loading.value = true
  await mediaStore.loadHistory()
  history.value = mediaStore.history
  loading.value = false
})
</script>
