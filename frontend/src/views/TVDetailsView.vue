<template>
  <div v-if="loading" class="h-full">
    <LoadingSpinner text="Загрузка информации о сериале..." />
  </div>

  <div v-else-if="tvShow" class="min-h-full">
    <!-- Backdrop -->
    <div class="relative h-96">
      <img
        :src="mediaStore.getImageURL(tvShow.backdrop_path, 'w1280')"
        :alt="tvShow.name"
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
            :src="mediaStore.getImageURL(tvShow.poster_path, 'w342')"
            :alt="tvShow.name"
            class="w-48 h-72 object-cover rounded-xl shadow-2xl"
            draggable="false"
          />
        </div>

        <!-- Info -->
        <div class="flex-1 pt-8">
          <h1 class="text-3xl font-bold text-white mb-2">{{ tvShow.name }}</h1>
          
          <div class="flex flex-wrap items-center gap-3 text-sm text-dark-400 mb-4">
            <span v-if="tvShow.first_air_date">{{ tvShow.first_air_date.split('-')[0] }}</span>
            <span>{{ tvShow.number_of_seasons }} сезон(ов)</span>
            <span>{{ tvShow.number_of_episodes }} эпизод(ов)</span>
            <div class="flex items-center gap-1 text-yellow-400">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <span class="text-white font-medium">{{ tvShow.vote_average?.toFixed(1) }}</span>
            </div>
          </div>

          <!-- Genres -->
          <div class="flex flex-wrap gap-2 mb-4">
            <span
              v-for="genre in tvShow.genres"
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
          <p v-if="tvShow.tagline" class="text-accent-400 italic mb-4">{{ tvShow.tagline }}</p>

          <!-- Overview -->
          <div class="mb-6">
            <h3 class="text-lg font-semibold text-white mb-2">Описание</h3>
            <p class="text-dark-300 leading-relaxed">{{ tvShow.overview || 'Описание отсутствует' }}</p>
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

      <!-- Seasons -->
      <div class="mt-8">
        <h3 class="text-xl font-semibold text-white mb-4">Сезоны</h3>
        
        <!-- Season Cards Grid -->
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div
            v-for="season in tvShow.seasons.filter(s => s.season_number > 0)"
            :key="season.id"
            class="group cursor-pointer"
            @click="openSeasonModal(season)"
          >
            <div class="relative aspect-[2/3] rounded-xl overflow-hidden bg-dark-800">
              <img
                v-if="season.poster_path"
                :src="mediaStore.getImageURL(season.poster_path, 'w342')"
                :alt="season.name"
                class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
                draggable="false"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-dark-800">
                <svg class="w-12 h-12 text-dark-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v16M17 4v16M3 8h4m10 0h4M3 12h18M3 16h4m10 0h4M4 20h16a1 1 0 001-1V5a1 1 0 00-1-1H4a1 1 0 00-1 1v14a1 1 0 001 1z"/>
                </svg>
              </div>
              
              <!-- Overlay on hover -->
              <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                <div class="w-12 h-12 bg-accent-600 rounded-full flex items-center justify-center">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>
                  </svg>
                </div>
              </div>
              
              <!-- Episode count badge -->
              <div class="absolute top-2 right-2 px-2 py-1 bg-dark-900/80 backdrop-blur-sm rounded-md text-xs font-medium text-white">
                {{ season.episode_count }} эп.
              </div>
            </div>
            
            <div class="mt-2">
              <h4 class="text-white font-medium text-sm line-clamp-1">{{ season.name }}</h4>
              <p v-if="season.air_date" class="text-dark-500 text-xs">{{ season.air_date?.split('-')[0] }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Similar TV Shows Section -->
      <div v-if="similarShows.length > 0" class="mt-8">
        <MediaRow title="Похожие сериалы" :items="similarShows" type="tv" />
      </div>
    </div>

    <!-- Season Episodes Modal -->
    <div 
      v-if="showSeasonModal" 
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-sm p-4"
      @click.self="showSeasonModal = false"
    >
      <div class="bg-dark-900 rounded-2xl w-full max-w-4xl max-h-[90vh] overflow-hidden flex flex-col">
        <!-- Modal Header -->
        <div class="flex items-center gap-3 px-4 py-3 border-b border-dark-800">
          <button
            @click="showSeasonModal = false"
            class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-dark-800 rounded-lg transition-colors text-dark-400 hover:text-white text-sm"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
            Назад
          </button>
          <h2 class="text-base font-semibold text-white flex-1 text-center">{{ selectedSeasonData?.name }}</h2>
          <span class="text-dark-500 text-sm w-[76px] text-right">{{ episodes.length }} эп.</span>
        </div>
        
        <!-- Episodes List -->
        <div class="flex-1 overflow-y-auto p-4">
          <div v-if="loadingEpisodes" class="py-8">
            <LoadingSpinner text="Загрузка эпизодов..." />
          </div>
          <div v-else-if="episodes.length > 0" class="space-y-3">
            <div
              v-for="episode in episodes"
              :key="episode.id"
              class="bg-dark-800 rounded-xl p-4 flex gap-4 hover:bg-dark-700/80 transition-colors cursor-pointer"
              @click="searchEpisodeTorrents(episode)"
            >
              <div class="relative flex-shrink-0">
                <img
                  v-if="episode.still_path"
                  :src="mediaStore.getImageURL(episode.still_path, 'w300')"
                  :alt="episode.name"
                  class="w-40 h-24 object-cover rounded-lg"
                  draggable="false"
                />
                <div v-else class="w-40 h-24 bg-dark-700 rounded-lg flex items-center justify-center">
                  <svg class="w-8 h-8 text-dark-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                  </svg>
                </div>
                <!-- Play overlay -->
                <div class="absolute inset-0 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
                  <div class="w-10 h-10 bg-accent-600 rounded-full flex items-center justify-center">
                      <svg class="w-4 h-4 text-white ml-0.5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
                        <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                    </svg>
                  </div>
                </div>
              </div>
              <div class="flex-1 min-w-0">
                <div class="flex items-start justify-between gap-3">
                  <div class="flex-1 min-w-0">
                    <h4 class="font-medium text-white">
                      <span class="text-accent-500">E{{ episode.episode_number }}</span> · {{ episode.name }}
                    </h4>
                    <p class="text-dark-400 text-sm mt-1 line-clamp-2">{{ episode.overview || 'Описание отсутствует' }}</p>
                    <div class="flex items-center gap-3 mt-2 text-xs text-dark-500">
                      <span v-if="episode.air_date">{{ formatDate(episode.air_date) }}</span>
                      <span v-if="episode.runtime">{{ episode.runtime }} мин</span>
                      <span v-if="episode.vote_average" class="flex items-center gap-1 text-yellow-400">
                        <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
                        </svg>
                        {{ episode.vote_average.toFixed(1) }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div v-else class="py-8 text-center text-dark-400">
            Эпизоды не найдены
          </div>
        </div>
      </div>
    </div>

    <!-- Torrent Search Modal -->
    <Modal :visible="showTorrentModal" title="Выберите торрент" @close="showTorrentModal = false">
      <div v-if="searchingTorrents" class="py-8">
        <LoadingSpinner text="Поиск торрентов..." />
      </div>
      <div v-else-if="torrentResults.length === 0" class="py-8 text-center text-dark-400">
        Торренты не найдены.
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

const tvShow = ref(null)
const loading = ref(true)
const isFavorite = ref(false)
const isBookmarked = ref(false)
const selectedSeason = ref(1)
const selectedSeasonData = ref(null)
const episodes = ref([])
const loadingEpisodes = ref(false)
const showTorrentModal = ref(false)
const showSeasonModal = ref(false)
const searchingTorrents = ref(false)
const torrentResults = ref([])
const currentEpisode = ref(null)
const trailers = ref([])
const showTrailerModal = ref(false)
const currentTrailer = ref(null)
const similarShows = ref([])

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleDateString('ru-RU', { day: 'numeric', month: 'long', year: 'numeric' })
}

const openSeasonModal = async (season) => {
  selectedSeasonData.value = season
  selectedSeason.value = season.season_number
  showSeasonModal.value = true
  loadingEpisodes.value = true
  
  try {
    const seasonDetails = await mediaStore.getSeasonDetails(tvShow.value.id, season.season_number)
    episodes.value = seasonDetails?.episodes || []
  } catch (error) {
    console.error('Failed to load episodes:', error)
  } finally {
    loadingEpisodes.value = false
  }
}

const playTrailer = (video) => {
  currentTrailer.value = video
  showTrailerModal.value = true
}

const toggleFavorite = async () => {
  if (isFavorite.value) {
    await mediaStore.removeFromFavorites(tvShow.value.id, 'tv')
  } else {
    await mediaStore.addToFavorites({
      tmdb_id: tvShow.value.id,
      media_type: 'tv',
      title: tvShow.value.name,
      poster: tvShow.value.poster_path,
      year: tvShow.value.first_air_date?.split('-')[0] || '',
      rating: tvShow.value.vote_average || 0
    })
  }
  isFavorite.value = !isFavorite.value
}

const toggleBookmark = async () => {
  if (isBookmarked.value) {
    await mediaStore.removeFromBookmarks(tvShow.value.id, 'tv')
  } else {
    await mediaStore.addToBookmarks({
      tmdb_id: tvShow.value.id,
      media_type: 'tv',
      title: tvShow.value.name,
      poster: tvShow.value.poster_path,
      year: tvShow.value.first_air_date?.split('-')[0] || ''
    })
  }
  isBookmarked.value = !isBookmarked.value
}

const searchTorrents = async () => {
  showTorrentModal.value = true
  searchingTorrents.value = true
  
  try {
    const searchQuery = `${tvShow.value.name} ${tvShow.value.first_air_date?.split('-')[0] || ''}`
    torrentResults.value = await torrentStore.searchJackett(searchQuery, 'tv')
  } catch (error) {
    console.error('Failed to search torrents:', error)
  } finally {
    searchingTorrents.value = false
  }
}

const searchEpisodeTorrents = async (episode) => {
  currentEpisode.value = episode
  showTorrentModal.value = true
  searchingTorrents.value = true
  
  try {
    const searchQuery = `${tvShow.value.name} S${String(selectedSeason.value).padStart(2, '0')}E${String(episode.episode_number).padStart(2, '0')}`
    torrentResults.value = await torrentStore.searchJackett(searchQuery, 'tv')
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
    
    const title = currentEpisode.value 
      ? `${tvShow.value.name} - ${currentEpisode.value.name}`
      : tvShow.value.name
    
    const added = await torrentStore.addTorrent(
      magnetUri,
      title,
      mediaStore.getImageURL(tvShow.value.poster_path)
    )
    
    console.log('Torrent added:', added)
    
    if (added) {
      const historyData = {
        tmdb_id: tvShow.value.id,
        media_type: 'tv',
        title: tvShow.value.name,
        poster: tvShow.value.poster_path,
        progress: 0,
        duration: 2700 // 45 минут по умолчанию
      }
      
      if (currentEpisode.value) {
        historyData.episode = `S${selectedSeason.value}E${currentEpisode.value.episode_number} - ${currentEpisode.value.name}`
        historyData.duration = (currentEpisode.value.runtime || 45) * 60
      }
      
      await mediaStore.addToHistory(historyData)
      
      // Navigate to player
      const playerTitle = currentEpisode.value 
        ? `${tvShow.value.name} - ${currentEpisode.value.name}`
        : tvShow.value.name
      
      router.push({
        name: 'Player',
        query: {
          hash: added.hash,
          title: playerTitle
        }
      })
    }
  } catch (error) {
    console.error('Failed to play torrent:', error)
  }
}

const loadTVData = async () => {
  loading.value = true
  const id = parseInt(route.params.id)
  const today = new Date().toISOString().split('T')[0]
  
  try {
    tvShow.value = await mediaStore.getTVDetails(id)
    isFavorite.value = await mediaStore.isFavorite(id, 'tv')
    
    // Load trailers
    if (window.go?.main?.TMDB?.GetTVVideos) {
      const videosResult = await window.go.main.TMDB.GetTVVideos(id)
      if (videosResult?.results) {
        trailers.value = videosResult.results
          .filter(v => v.site === 'YouTube' && ['Trailer', 'Teaser', 'Clip'].includes(v.type))
          .slice(0, 6)
      }
    }

    // Load similar TV shows
    if (window.go?.main?.TMDB?.GetTVSimilar) {
      const similarResult = await window.go.main.TMDB.GetTVSimilar(id)
      if (similarResult?.results) {
        similarShows.value = similarResult.results
          .filter(s => s.first_air_date && s.first_air_date <= today)
          .slice(0, 20)
      }
    }
  } catch (error) {
    console.error('Failed to load TV details:', error)
  } finally {
    loading.value = false
  }
}

onMounted(loadTVData)

watch(() => route.params.id, () => {
  if (route.name === 'TVDetails') {
    loadTVData()
  }
})
</script>
