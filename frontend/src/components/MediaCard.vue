<template>
  <div
    class="card-hover group"
    @click="$emit('click', media)"
  >
    <div class="media-poster">
      <img
        :src="posterUrl"
        :alt="title"
        loading="lazy"
        draggable="false"
        class="transition-transform duration-300"
        @error="handleImageError"
      />
      
      <!-- Rating Badge -->
      <div class="rating-badge" v-if="rating > 0">
        <svg class="w-3 h-3 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
        </svg>
        <span class="text-white text-xs font-semibold">{{ rating.toFixed(1) }}</span>
      </div>

      <!-- Overlay on hover -->
      <div class="absolute inset-0 bg-gradient-to-t from-dark-900 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <div class="absolute bottom-0 left-0 right-0 p-3">
          <div class="flex gap-2">
            <button
              @click.stop="toggleFavorite"
              class="p-2 bg-dark-800/80 backdrop-blur-sm rounded-full hover:bg-accent-600 transition-colors"
              :class="{ 'bg-red-500 hover:bg-red-600': isFav }"
            >
              <svg class="w-4 h-4" :fill="isFav ? 'currentColor' : 'none'" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
              </svg>
            </button>
            <button
              @click.stop="toggleBookmark"
              class="p-2 bg-dark-800/80 backdrop-blur-sm rounded-full hover:bg-accent-600 transition-colors"
              :class="{ 'bg-accent-600': isBookmarked }"
            >
              <svg class="w-4 h-4" :fill="isBookmarked ? 'currentColor' : 'none'" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Info -->
    <div class="p-3 h-[60px] flex flex-col justify-between">
      <h3 class="font-medium text-white text-sm line-clamp-2 mb-1">{{ title }}</h3>
      <p class="text-dark-500 text-xs">{{ year }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useMediaStore } from '../stores/media'

const props = defineProps({
  media: {
    type: Object,
    required: true
  },
  type: {
    type: String,
    default: 'movie' // movie | tv
  }
})

defineEmits(['click'])

const mediaStore = useMediaStore()
const isFav = ref(false)
const isBookmarked = ref(false)

const title = computed(() => props.media.title || props.media.name || 'Без названия')

const year = computed(() => {
  const date = props.media.release_date || props.media.first_air_date
  return date ? date.split('-')[0] : ''
})

const rating = computed(() => props.media.vote_average || 0)

const posterUrl = computed(() => {
  return mediaStore.getImageURL(props.media.poster_path, 'w342')
})

const handleImageError = (e) => {
  e.target.src = 'data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" width="342" height="513" viewBox="0 0 342 513"><rect fill="%231e293b" width="342" height="513"/><text fill="%2364748b" font-family="sans-serif" font-size="24" x="50%" y="50%" text-anchor="middle" dy=".3em">Нет постера</text></svg>'
}

const toggleFavorite = async () => {
  if (isFav.value) {
    await mediaStore.removeFromFavorites(props.media.id, props.type)
  } else {
    await mediaStore.addToFavorites({
      tmdb_id: props.media.id,
      media_type: props.type,
      title: title.value,
      poster: props.media.poster_path,
      year: year.value,
      rating: rating.value
    })
  }
  isFav.value = !isFav.value
}

const toggleBookmark = async () => {
  if (isBookmarked.value) {
    await mediaStore.removeFromBookmarks(props.media.id, props.type)
  } else {
    await mediaStore.addToBookmarks({
      tmdb_id: props.media.id,
      media_type: props.type,
      title: title.value,
      poster: props.media.poster_path,
      year: year.value
    })
  }
  isBookmarked.value = !isBookmarked.value
}

onMounted(async () => {
  isFav.value = await mediaStore.isFavorite(props.media.id, props.type)
})
</script>
