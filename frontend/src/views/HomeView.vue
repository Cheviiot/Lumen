<template>
  <div class="p-6 space-y-8">
    <!-- Hero Carousel -->
    <section v-if="featuredMovies.length > 0" class="relative h-96 rounded-2xl overflow-hidden group">
      <!-- Slides -->
      <div class="relative w-full h-full">
        <TransitionGroup name="carousel">
          <div
            v-for="(movie, index) in featuredMovies"
            :key="movie.id"
            v-show="index === currentSlide"
            class="absolute inset-0"
          >
            <img
              :src="mediaStore.getImageURL(movie.backdrop_path, 'w1280')"
              :alt="movie.title || movie.name"
              class="absolute inset-0 w-full h-full object-cover"
              draggable="false"
            />
            <div class="absolute inset-0 bg-gradient-to-r from-dark-900 via-dark-900/70 to-transparent" />
            <div class="absolute inset-0 bg-gradient-to-t from-dark-900 via-transparent to-transparent" />
            
            <div class="relative h-full flex flex-col justify-end p-8">
              <div class="flex items-center gap-3 mb-3">
                <span class="px-3 py-1 bg-accent-600 text-white text-xs font-medium rounded-full">
                  Рекомендуем
                </span>
                <span v-if="movie.vote_average" class="flex items-center gap-1 text-yellow-400 text-sm">
                  <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
                  </svg>
                  {{ movie.vote_average.toFixed(1) }}
                </span>
              </div>
              <h1 class="text-4xl font-bold text-white mb-2">{{ movie.title || movie.name }}</h1>
              <p class="text-dark-300 max-w-xl line-clamp-2 mb-12">{{ movie.overview }}</p>
            </div>
          </div>
        </TransitionGroup>
      </div>

      <!-- Watch Button - Bottom Left -->
      <button 
        @click="openDetails(featuredMovies[currentSlide], currentSlide)" 
        class="absolute left-8 bottom-6 btn-primary flex items-center gap-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
          <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        Смотреть
      </button>

      <!-- Indicators - Center -->
      <div class="absolute bottom-6 left-1/2 -translate-x-1/2 flex gap-2">
        <button
          v-for="(_, index) in featuredMovies"
          :key="index"
          @click="goToSlide(index)"
          class="w-2 h-2 rounded-full transition-all duration-300"
          :class="index === currentSlide ? 'bg-accent-500 w-6' : 'bg-white/50 hover:bg-white/80'"
        />
      </div>

      <!-- Navigation Arrows - Right Bottom -->
      <div class="absolute right-4 bottom-6 flex gap-2">
        <button
          @click="prevSlide"
          class="p-2.5 bg-dark-900/70 backdrop-blur-sm rounded-lg text-white hover:bg-dark-800 transition-colors duration-200"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
        </button>
        <button
          @click="nextSlide"
          class="p-2.5 bg-dark-900/70 backdrop-blur-sm rounded-lg text-white hover:bg-dark-800 transition-colors duration-200"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
        </button>
      </div>

      <!-- Progress Bar -->
      <div class="absolute bottom-0 left-0 right-0 h-1 bg-dark-800/50">
        <div
          class="h-full bg-accent-600 transition-all duration-100"
          :style="{ width: `${progressWidth}%` }"
        />
      </div>
    </section>

    <!-- Loading -->
    <LoadingSpinner v-if="loading" text="Загрузка контента..." />

    <!-- Content -->
    <template v-else>
      <!-- Trending Movies -->
      <MediaRow
        title="Популярные фильмы"
        :items="trendingMovies"
        type="movie"
      />

      <!-- Trending TV -->
      <MediaRow
        title="Популярные сериалы"
        :items="trendingTV"
        type="tv"
      />

      <!-- Animation -->
      <MediaRow
        title="Мультфильмы"
        :items="animationMovies"
        type="movie"
      />

      <!-- Continue Watching (History) -->
      <MediaRow
        v-if="recentHistory.length > 0"
        title="Продолжить просмотр"
        :items="recentHistory"
        type="movie"
      />
    </template>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useMediaStore } from '../stores/media'
import MediaRow from '../components/MediaRow.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'

const router = useRouter()
const mediaStore = useMediaStore()

const loading = ref(true)
const trendingMovies = ref([])
const trendingTV = ref([])
const animationMovies = ref([])

// Carousel state
const currentSlide = ref(0)
const slideInterval = ref(null)
const progressInterval = ref(null)
const progressWidth = ref(0)
const SLIDE_DURATION = 8000 // 8 секунд на слайд

// Берём первые 5-7 фильмов/сериалов для карусели
const featuredMovies = computed(() => {
  const movies = trendingMovies.value.slice(0, 4)
  const tvShows = trendingTV.value.slice(0, 3)
  // Чередуем фильмы и сериалы
  const mixed = []
  const maxLen = Math.max(movies.length, tvShows.length)
  for (let i = 0; i < maxLen; i++) {
    if (movies[i]) mixed.push({ ...movies[i], mediaType: 'movie' })
    if (tvShows[i]) mixed.push({ ...tvShows[i], mediaType: 'tv' })
  }
  return mixed.slice(0, 7)
})

const recentHistory = computed(() => {
  return mediaStore.history.slice(0, 10).map(h => ({
    id: h.tmdb_id,
    title: h.title,
    poster_path: h.poster,
    vote_average: 0
  }))
})

// Carousel functions
const startCarousel = () => {
  stopCarousel()
  progressWidth.value = 0
  
  // Progress bar update
  const progressStep = 100 / (SLIDE_DURATION / 100)
  progressInterval.value = setInterval(() => {
    progressWidth.value += progressStep
  }, 100)
  
  // Slide change
  slideInterval.value = setInterval(() => {
    nextSlide()
  }, SLIDE_DURATION)
}

const stopCarousel = () => {
  if (slideInterval.value) {
    clearInterval(slideInterval.value)
    slideInterval.value = null
  }
  if (progressInterval.value) {
    clearInterval(progressInterval.value)
    progressInterval.value = null
  }
}

const nextSlide = () => {
  currentSlide.value = (currentSlide.value + 1) % featuredMovies.value.length
  progressWidth.value = 0
}

const prevSlide = () => {
  currentSlide.value = currentSlide.value === 0 
    ? featuredMovies.value.length - 1 
    : currentSlide.value - 1
  progressWidth.value = 0
  startCarousel()
}

const goToSlide = (index) => {
  currentSlide.value = index
  progressWidth.value = 0
  startCarousel()
}

const openDetails = (movie, index) => {
  const type = featuredMovies.value[index]?.mediaType || 'movie'
  if (type === 'tv') {
    router.push({ name: 'TVDetails', params: { id: movie.id } })
  } else {
    router.push({ name: 'MovieDetails', params: { id: movie.id } })
  }
}

onMounted(async () => {
  loading.value = true
  try {
    const [movies, tv, animation] = await Promise.all([
      mediaStore.fetchTrendingMovies(),
      mediaStore.fetchTrendingTV(),
      mediaStore.fetchAnimationMovies(),
      mediaStore.loadHistory()
    ])
    
    trendingMovies.value = movies?.results || []
    trendingTV.value = tv?.results || []
    animationMovies.value = animation?.results || []
    
    // Запускаем карусель после загрузки данных
    if (featuredMovies.value.length > 0) {
      startCarousel()
    }
  } catch (error) {
    console.error('Failed to load home content:', error)
  } finally {
    loading.value = false
  }
})

onUnmounted(() => {
  stopCarousel()
})
</script>

<style scoped>
.carousel-enter-active,
.carousel-leave-active {
  transition: all 0.5s ease;
}

.carousel-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.carousel-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
</style>
