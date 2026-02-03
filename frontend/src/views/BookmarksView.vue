<template>
  <div class="p-6">
    <h1 class="text-2xl font-bold text-white mb-6">Закладки</h1>

    <LoadingSpinner v-if="loading" text="Загрузка закладок..." />

    <div v-else-if="bookmarks.length === 0" class="flex flex-col items-center justify-center h-64 text-center">
      <svg class="w-16 h-16 text-dark-600 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/>
      </svg>
      <h3 class="text-lg font-medium text-dark-400 mb-2">Нет закладок</h3>
      <p class="text-dark-500 text-sm">Добавляйте фильмы и сериалы в закладки для быстрого доступа</p>
    </div>

    <div v-else class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
      <div
        v-for="item in bookmarks"
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
          
          <!-- Remove button -->
          <button
            @click.stop="removeFromBookmarks(item)"
            class="absolute top-2 left-2 p-2 bg-dark-900/80 backdrop-blur-sm rounded-full opacity-0 group-hover:opacity-100 transition-opacity hover:bg-red-600"
          >
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
          
          <!-- Bookmark indicator -->
          <div class="absolute top-2 right-2">
            <svg class="w-6 h-6 text-accent-500" fill="currentColor" viewBox="0 0 24 24">
              <path d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/>
            </svg>
          </div>
        </div>
        <div class="p-3">
          <h3 class="font-medium text-white text-sm line-clamp-2 mb-1">{{ item.title }}</h3>
          <p class="text-dark-400 text-xs">{{ item.year }}</p>
          <p v-if="item.note" class="text-dark-500 text-xs mt-1 line-clamp-1">{{ item.note }}</p>
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

const bookmarks = ref([])
const loading = ref(true)

const getImageURL = (path) => {
  return mediaStore.getImageURL(path, 'w342')
}

const openDetails = (item) => {
  const routeName = item.media_type === 'tv' ? 'TVDetails' : 'MovieDetails'
  router.push({ name: routeName, params: { id: item.tmdb_id } })
}

const removeFromBookmarks = async (item) => {
  await mediaStore.removeFromBookmarks(item.tmdb_id, item.media_type)
  bookmarks.value = bookmarks.value.filter(
    b => !(b.tmdb_id === item.tmdb_id && b.media_type === item.media_type)
  )
}

onMounted(async () => {
  loading.value = true
  await mediaStore.loadBookmarks()
  bookmarks.value = mediaStore.bookmarks
  loading.value = false
})
</script>
