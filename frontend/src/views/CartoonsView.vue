<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-white">Мультфильмы</h1>
      
      <div class="flex gap-3">
        <select
          v-model="contentType"
          class="input-field w-auto text-sm"
        >
          <option value="movies">Фильмы</option>
          <option value="series">Сериалы</option>
        </select>
        <select
          v-model="sortBy"
          class="input-field w-auto text-sm"
        >
          <option value="popularity">По популярности</option>
          <option value="date">По дате выхода</option>
          <option value="vote_average">По рейтингу</option>
        </select>
      </div>
    </div>

    <LoadingSpinner v-if="loading && items.length === 0" text="Загрузка мультфильмов..." />

    <MediaGrid
      v-else
      :items="items"
      :type="contentType === 'movies' ? 'movie' : 'tv'"
      :has-more="hasMore"
      :loading="loadingMore"
      @load-more="loadMore"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useMediaStore } from '../stores/media'
import MediaGrid from '../components/MediaGrid.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'

const mediaStore = useMediaStore()

const contentType = ref('movies')
const sortBy = ref('popularity')
const animationMovies = ref([])
const animationTV = ref([])
const loading = ref(true)
const loadingMore = ref(false)
const moviePage = ref(1)
const tvPage = ref(1)
const movieTotalPages = ref(1)
const tvTotalPages = ref(1)

const items = computed(() => {
  return contentType.value === 'movies' ? animationMovies.value : animationTV.value
})

const hasMore = computed(() => {
  if (contentType.value === 'movies') {
    return moviePage.value < movieTotalPages.value
  }
  return tvPage.value < tvTotalPages.value
})

const fetchAnimationMovies = async (pageNum = 1) => {
  const result = await mediaStore.fetchAnimationMovies(pageNum)
  if (result) {
    movieTotalPages.value = result.total_pages || 1
    return result.results || []
  }
  return []
}

const fetchAnimationTV = async (pageNum = 1) => {
  const result = await mediaStore.fetchAnimationTV(pageNum)
  if (result) {
    tvTotalPages.value = result.total_pages || 1
    return result.results || []
  }
  return []
}

const loadMore = async () => {
  if (loadingMore.value || !hasMore.value) return
  
  loadingMore.value = true
  
  try {
    if (contentType.value === 'movies') {
      moviePage.value++
      const newMovies = await fetchAnimationMovies(moviePage.value)
      animationMovies.value = [...animationMovies.value, ...newMovies]
    } else {
      tvPage.value++
      const newTV = await fetchAnimationTV(tvPage.value)
      animationTV.value = [...animationTV.value, ...newTV]
    }
  } finally {
    loadingMore.value = false
  }
}

watch(contentType, async (newType) => {
  if (newType === 'movies' && animationMovies.value.length === 0) {
    loading.value = true
    moviePage.value = 1
    animationMovies.value = await fetchAnimationMovies(1)
    loading.value = false
  } else if (newType === 'series' && animationTV.value.length === 0) {
    loading.value = true
    tvPage.value = 1
    animationTV.value = await fetchAnimationTV(1)
    loading.value = false
  }
})

watch(sortBy, async () => {
  loading.value = true
  if (contentType.value === 'movies') {
    moviePage.value = 1
    animationMovies.value = await fetchAnimationMovies(1)
  } else {
    tvPage.value = 1
    animationTV.value = await fetchAnimationTV(1)
  }
  loading.value = false
})

onMounted(async () => {
  loading.value = true
  animationMovies.value = await fetchAnimationMovies(1)
  loading.value = false
})
</script>
