<template>
  <div v-if="loading" class="h-full">
    <LoadingSpinner text="Загрузка информации о фильме..." />
  </div>

  <div v-else-if="movie" class="min-h-full">
    <!-- Backdrop -->
    <div class="relative h-96">
      <img
        :src="mediaStore.getImageURL(movie.backdrop_path, 'w1280')"
        :alt="movie.title"
        class="absolute inset-0 w-full h-full object-cover"
        draggable="false"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-dark-900 via-dark-900/60 to-dark-900/30" />
      
      <!-- Back button -->
      <button
        @click="$router.back()"
        class="absolute top-6 left-6 flex items-center gap-2 px-4 py-2 bg-dark-800/80 backdrop-blur-sm rounded-lg hover:bg-dark-700 transition-colors text-white text-sm font-medium"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
        </svg>
        Назад
      </button>
    </div>

    <!-- Content -->
    <div class="relative -mt-32 px-6 pb-6">
      <div class="flex gap-6">
        <!-- Poster -->
        <div class="flex-shrink-0">
          <img
            :src="mediaStore.getImageURL(movie.poster_path, 'w342')"
            :alt="movie.title"
            class="w-48 h-72 object-cover rounded-xl shadow-2xl"
            draggable="false"
          />
        </div>

        <!-- Info -->
        <div class="flex-1 pt-8">
          <h1 class="text-3xl font-bold text-white mb-2">{{ movie.title }}</h1>
          
          <div class="flex flex-wrap items-center gap-3 text-sm text-dark-400 mb-4">
            <span v-if="movie.release_date">{{ movie.release_date.split('-')[0] }}</span>
            <span v-if="movie.runtime">{{ formatRuntime(movie.runtime) }}</span>
            <div class="flex items-center gap-1 text-yellow-400">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <span class="text-white font-medium">{{ movie.vote_average?.toFixed(1) }}</span>
            </div>
          </div>

          <!-- Genres -->
          <div class="flex flex-wrap gap-2 mb-4">
            <span
              v-for="genre in movie.genres"
              :key="genre.id"
              class="genre-tag"
            >
              {{ genre.name }}
            </span>
          </div>

          <!-- Actions -->
          <div class="flex gap-3 mb-6">
            <button @click="searchTorrents" class="btn-primary flex items-center gap-2">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              Смотреть
            </button>
            <button @click="toggleFavorite" class="btn-secondary flex items-center gap-2">
              <svg class="w-5 h-5" :fill="isFavorite ? 'currentColor' : 'none'" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
              </svg>
              {{ isFavorite ? 'В избранном' : 'В избранное' }}
            </button>
            <button @click="toggleBookmark" class="btn-secondary flex items-center gap-2">
              <svg class="w-5 h-5" :fill="isBookmarked ? 'currentColor' : 'none'" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/>
              </svg>
              {{ isBookmarked ? 'В закладках' : 'В закладки' }}
            </button>
          </div>

          <!-- Tagline -->
          <p v-if="movie.tagline" class="text-accent-400 italic mb-4">{{ movie.tagline }}</p>

          <!-- Overview -->
          <div class="mb-6">
            <h3 class="text-lg font-semibold text-white mb-2">Описание</h3>
            <p class="text-dark-300 leading-relaxed">{{ movie.overview || 'Описание отсутствует' }}</p>
          </div>
        </div>
      </div>

      <!-- Trailers Section -->
      <div v-if="trailers.length > 0" class="mt-8">
        <h3 class="text-xl font-semibold text-white mb-4">Трейлеры и видео</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <div
            v-for="video in trailers"
            :key="video.id"
            class="relative rounded-xl overflow-hidden bg-dark-800 cursor-pointer group"
            @click="playTrailer(video)"
          >
            <img
              :src="`https://img.youtube.com/vi/${video.key}/mqdefault.jpg`"
              :alt="video.name"
              class="w-full aspect-video object-cover"
              draggable="false"
            />
            <div class="absolute inset-0 bg-black/40 flex items-center justify-center group-hover:bg-black/20 transition-colors">
              <div class="w-14 h-14 bg-accent-600 rounded-full flex items-center justify-center group-hover:scale-110 transition-transform">
                  <svg class="w-6 h-6 text-white ml-1" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
              </div>
            </div>
            <div class="absolute bottom-0 left-0 right-0 p-3 bg-gradient-to-t from-black/80 to-transparent">
              <p class="text-white text-sm font-medium line-clamp-1">{{ video.name }}</p>
              <p class="text-dark-400 text-xs">{{ video.type }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Recommendations Section -->
      <div v-if="similarMovies.length > 0" class="mt-8">
        <MediaRow title="Похожие фильмы" :items="similarMovies" type="movie" />
      </div>

      <!-- Collection Section -->
      <div v-if="collection.length > 0" class="mt-8">
        <MediaRow title="Коллекция" :items="collection" type="movie" />
      </div>
    </div>

    <!-- Trailer Modal -->
    <div 
      v-if="showTrailerModal" 
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/90"
      @click.self="showTrailerModal = false"
    >
      <button
        @click="showTrailerModal = false"
        class="absolute top-4 right-4 p-2 bg-dark-800/80 rounded-full hover:bg-dark-700 transition-colors z-10"
      >
        <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
        </svg>
      </button>
      <a
        v-if="currentTrailer"
        :href="`https://www.youtube.com/watch?v=${currentTrailer.key}`"
        target="_blank"
        class="absolute top-4 left-4 flex items-center gap-2 px-4 py-2 bg-red-600 hover:bg-red-700 rounded-lg text-white text-sm transition-colors z-10"
      >
        <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
          <path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/>
        </svg>
        Открыть на YouTube
      </a>
      <div class="w-full max-w-5xl aspect-video">
        <iframe
          v-if="currentTrailer"
          :src="`https://www.youtube-nocookie.com/embed/${currentTrailer.key}?autoplay=1&rel=0&modestbranding=1&origin=${encodeURIComponent(window.location.origin)}`"
          class="w-full h-full rounded-xl"
          frameborder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowfullscreen
        ></iframe>
      </div>
    </div>

    <!-- Torrent Search Modal -->
    <Modal :visible="showTorrentModal" title="Выберите торрент" @close="showTorrentModal = false">
      <div v-if="searchingTorrents" class="py-8">
        <LoadingSpinner text="Поиск торрентов..." />
      </div>
      <div v-else-if="torrentResults.length === 0" class="py-8 text-center text-dark-400">
        Торренты не найдены. Попробуйте позже или проверьте настройки Jackett.
      </div>
      <div v-else class="space-y-3 max-h-96 overflow-y-auto">
        <TorrentCard
          v-for="torrent in torrentResults"
          :key="torrent.Guid"
          :torrent="torrent"
          @play="playTorrent"
        />
      </div>
    </Modal>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMediaStore } from '../stores/media'
import { useTorrentStore } from '../stores/torrent'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import Modal from '../components/Modal.vue'
import TorrentCard from '../components/TorrentCard.vue'
import MediaRow from '../components/MediaRow.vue'

const route = useRoute()
const router = useRouter()
const mediaStore = useMediaStore()
const torrentStore = useTorrentStore()

const movie = ref(null)
const loading = ref(true)
const isFavorite = ref(false)
const isBookmarked = ref(false)
const showTorrentModal = ref(false)
const searchingTorrents = ref(false)
const torrentResults = ref([])
const trailers = ref([])
const showTrailerModal = ref(false)
const currentTrailer = ref(null)
const collection = ref([])
const similarMovies = ref([])

const formatRuntime = (minutes) => {
  const hours = Math.floor(minutes / 60)
  const mins = minutes % 60
  return `${hours}ч ${mins}мин`
}

const toggleFavorite = async () => {
  if (isFavorite.value) {
    await mediaStore.removeFromFavorites(movie.value.id, 'movie')
  } else {
    await mediaStore.addToFavorites({
      tmdb_id: movie.value.id,
      media_type: 'movie',
      title: movie.value.title,
      poster: movie.value.poster_path,
      year: movie.value.release_date?.split('-')[0] || '',
      rating: movie.value.vote_average || 0
    })
  }
  isFavorite.value = !isFavorite.value
}

const toggleBookmark = async () => {
  if (isBookmarked.value) {
    await mediaStore.removeFromBookmarks(movie.value.id, 'movie')
  } else {
    await mediaStore.addToBookmarks({
      tmdb_id: movie.value.id,
      media_type: 'movie',
      title: movie.value.title,
      poster: movie.value.poster_path,
      year: movie.value.release_date?.split('-')[0] || ''
    })
  }
  isBookmarked.value = !isBookmarked.value
}

const searchTorrents = async () => {
  showTorrentModal.value = true
  searchingTorrents.value = true
  
  try {
    const searchQuery = `${movie.value.title} ${movie.value.release_date?.split('-')[0] || ''}`
    torrentResults.value = await torrentStore.searchJackett(searchQuery, 'movie')
  } catch (error) {
    console.error('Failed to search torrents:', error)
  } finally {
    searchingTorrents.value = false
  }
}

const playTorrent = async (torrent) => {
  try {
    // Use MagnetUri if available, otherwise use Link (for public parsers)
    const magnetUri = torrent.MagnetUri || torrent.Link
    
    if (!magnetUri) {
      console.error('No MagnetUri or Link found in torrent object!')
      alert('Невозможно получить ссылку для этого торрента')
      return
    }
    
    const added = await torrentStore.addTorrent(
      magnetUri,
      movie.value.title,
      mediaStore.getImageURL(movie.value.poster_path)
    )
    
    console.log('Torrent added:', added)
    
    if (added) {
      // Add to history
      await mediaStore.addToHistory({
        tmdb_id: movie.value.id,
        media_type: 'movie',
        title: movie.value.title,
        poster: movie.value.poster_path,
        progress: 0,
        duration: movie.value.runtime * 60
      })
      
      // Navigate to player
      router.push({
        name: 'Player',
        query: {
          hash: added.hash,
          title: movie.value.title
        }
      })
    }
  } catch (error) {
    console.error('Failed to play torrent:', error)
  }
}

const loadMovieData = async () => {
  loading.value = true
  const id = parseInt(route.params.id)
  const today = new Date().toISOString().split('T')[0]
  
  try {
    movie.value = await mediaStore.getMovieDetails(id)
    isFavorite.value = await mediaStore.isFavorite(id, 'movie')
    
    // Load trailers
    if (window.go?.main?.TMDB?.GetMovieVideos) {
      const videosResult = await window.go.main.TMDB.GetMovieVideos(id)
      if (videosResult?.results) {
        // Filter for YouTube trailers and teasers
        trailers.value = videosResult.results
          .filter(v => v.site === 'YouTube' && ['Trailer', 'Teaser', 'Clip'].includes(v.type))
          .slice(0, 6)
      }
    }

    // Load similar movies
    if (window.go?.main?.TMDB?.GetMovieSimilar) {
      const similarResult = await window.go.main.TMDB.GetMovieSimilar(id)
      if (similarResult?.results) {
        similarMovies.value = similarResult.results
          .filter(m => m.release_date && m.release_date <= today)
          .slice(0, 20)
      }
    }

    // Load collection if movie belongs to one
    if (movie.value.belongs_to_collection && window.go?.main?.TMDB?.GetCollection) {
      const collectionResult = await window.go.main.TMDB.GetCollection(movie.value.belongs_to_collection.id)
      if (collectionResult?.parts) {
        collection.value = collectionResult.parts
          .filter(m => m.id !== id && m.release_date && m.release_date <= today)
          .slice(0, 20)
      }
    }
  } catch (error) {
    console.error('Failed to load movie details:', error)
  } finally {
    loading.value = false
  }
}

onMounted(loadMovieData)

watch(() => route.params.id, () => {
  if (route.name === 'MovieDetails') {
    loadMovieData()
  }
})

const playTrailer = (video) => {
  currentTrailer.value = video
  showTrailerModal.value = true
}
</script>
