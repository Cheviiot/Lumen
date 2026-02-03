<template>
  <div class="p-6">
    <h1 class="text-2xl font-bold text-white mb-6">Избранные</h1>

    <LoadingSpinner v-if="loading" text="Загрузка избранного..." />

    <div v-else-if="favorites.length === 0" class="flex flex-col items-center justify-center h-64 text-center">
      <svg class="w-16 h-16 text-dark-600 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
      </svg>
      <h3 class="text-lg font-medium text-dark-400 mb-2">Нет избранных</h3>
      <p class="text-dark-500 text-sm">Добавляйте фильмы и сериалы в избранное, чтобы они отображались здесь</p>
    </div>

    <div v-else class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
      <div
        v-for="item in favorites"
        :key="`${item.media_type}-${item.tmdb_id}`"
        class="card-hover group cursor-pointer"
        @click="openDetails(item)"
      >
        <div class="media-poster">
          <img
            :src="getImageURL(item.poster)"
            :alt="item.title"
            class="transition-transform duration-300 group-hover:scale-110"
            draggable="false"
          />
          <div class="rating-badge" v-if="item.rating > 0">
            <svg class="w-3.5 h-3.5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
            </svg>
            <span class="text-white">{{ item.rating.toFixed(1) }}</span>
          </div>
          
          <!-- Remove button -->
          <button
            @click.stop="removeFromFavorites(item)"
            class="absolute top-2 left-2 p-2 bg-red-500/80 backdrop-blur-sm rounded-full opacity-0 group-hover:opacity-100 transition-opacity hover:bg-red-600"
          >
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
        <div class="p-3">
          <h3 class="font-medium text-white text-sm line-clamp-2 mb-1">{{ item.title }}</h3>
          <p class="text-dark-400 text-xs">{{ item.year }}</p>
        </div>
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

const favorites = ref([])
const loading = ref(true)

const getImageURL = (path) => {
  return mediaStore.getImageURL(path, 'w342')
}

const openDetails = (item) => {
  const routeName = item.media_type === 'tv' ? 'TVDetails' : 'MovieDetails'
  router.push({ name: routeName, params: { id: item.tmdb_id } })
}

const removeFromFavorites = async (item) => {
  await mediaStore.removeFromFavorites(item.tmdb_id, item.media_type)
  favorites.value = favorites.value.filter(
    f => !(f.tmdb_id === item.tmdb_id && f.media_type === item.media_type)
  )
}

onMounted(async () => {
  loading.value = true
  await mediaStore.loadFavorites()
  favorites.value = mediaStore.favorites
  loading.value = false
})
</script>
