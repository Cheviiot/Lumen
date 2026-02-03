<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold text-white">Сериалы</h1>
      
      <div class="flex gap-3">
        <select
          v-model="sortBy"
          class="input-field w-auto text-sm"
        >
          <option value="popularity">По популярности</option>
          <option value="first_air_date">По дате выхода</option>
          <option value="vote_average">По рейтингу</option>
        </select>
      </div>
    </div>

    <LoadingSpinner v-if="loading && series.length === 0" text="Загрузка сериалов..." />

    <MediaGrid
      v-else
      :items="series"
      type="tv"
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

const series = ref([])
const loading = ref(true)
const loadingMore = ref(false)
const page = ref(1)
const totalPages = ref(1)
const sortBy = ref('popularity')

const hasMore = ref(true)

const fetchSeries = async (pageNum = 1) => {
  const result = await mediaStore.fetchPopularTV(pageNum)
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
    const newSeries = await fetchSeries(page.value)
    series.value = [...series.value, ...newSeries]
  } finally {
    loadingMore.value = false
  }
}

watch(sortBy, async () => {
  loading.value = true
  page.value = 1
  series.value = await fetchSeries(1)
  loading.value = false
})

onMounted(async () => {
  loading.value = true
  series.value = await fetchSeries(1)
  loading.value = false
})
</script>
