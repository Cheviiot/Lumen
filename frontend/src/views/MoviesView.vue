<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-white">Фильмы</h1>
      
      <div class="flex gap-3">
        <select
          v-model="sortBy"
          class="input-field w-auto text-sm"
        >
          <option value="popularity">По популярности</option>
          <option value="release_date">По дате выхода</option>
          <option value="vote_average">По рейтингу</option>
        </select>
      </div>
    </div>

    <LoadingSpinner v-if="loading && movies.length === 0" text="Загрузка фильмов..." />

    <MediaGrid
      v-else
      :items="movies"
      type="movie"
      :has-more="hasMore"
      :loading="loadingMore"
      @load-more="loadMore"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useMediaStore } from '../stores/media'
import MediaGrid from '../components/MediaGrid.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'

const mediaStore = useMediaStore()

const movies = ref([])
const loading = ref(true)
const loadingMore = ref(false)
const page = ref(1)
const totalPages = ref(1)
const sortBy = ref('popularity')

const hasMore = ref(true)

const fetchMovies = async (pageNum = 1) => {
  const result = await mediaStore.fetchPopularMovies(pageNum)
  if (result) {
    totalPages.value = result.total_pages || 1
    hasMore.value = pageNum < totalPages.value
    return result.results || []
  }
  return []
}

const loadMore = async () => {
  if (loadingMore.value || !hasMore.value) return
  
  loadingMore.value = true
  page.value++
  
  try {
    const newMovies = await fetchMovies(page.value)
    movies.value = [...movies.value, ...newMovies]
  } finally {
    loadingMore.value = false
  }
}

watch(sortBy, async () => {
  loading.value = true
  page.value = 1
  movies.value = await fetchMovies(1)
  loading.value = false
})

onMounted(async () => {
  loading.value = true
  movies.value = await fetchMovies(1)
  loading.value = false
})
</script>
