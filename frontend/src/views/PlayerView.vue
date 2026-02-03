<template>
  <div class="h-full flex flex-col bg-black">
    <!-- Player Header -->
    <div 
      class="absolute top-0 left-0 right-0 z-10 p-4 bg-gradient-to-b from-black/80 to-transparent transition-opacity duration-300"
      :class="{ 'opacity-0': !showControls }"
    >
      <div class="flex items-center justify-between">
        <button
          @click="goBack"
          class="flex items-center gap-2 text-white hover:text-accent-400 transition-colors"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
          Назад
        </button>
        <h2 class="text-white font-medium">{{ title }}</h2>
        <div class="w-20"></div>
      </div>
    </div>

    <!-- Video Player -->
    <div 
      class="flex-1 flex items-center justify-center relative"
      @mousemove="handleMouseMove"
      @click="togglePlay"
    >
      <video
        ref="videoRef"
        class="max-w-full max-h-full"
        :src="streamUrl"
        @timeupdate="handleTimeUpdate"
        @loadedmetadata="handleLoadedMetadata"
        @ended="handleEnded"
        @waiting="isBuffering = true"
        @playing="isBuffering = false"
        @error="handleError"
      />

      <!-- Buffering Indicator -->
      <div v-if="isBuffering" class="absolute inset-0 flex items-center justify-center">
        <div class="flex flex-col items-center gap-4">
          <div class="w-16 h-16 border-4 border-accent-500 border-t-transparent rounded-full animate-spin"></div>
          <p class="text-white">Буферизация...</p>
        </div>
      </div>

      <!-- Play/Pause Overlay -->
      <div 
        v-if="!isPlaying && !isBuffering"
        class="absolute inset-0 flex items-center justify-center bg-black/40"
      >
        <button class="p-6 bg-accent-600/80 rounded-full hover:bg-accent-600 transition-colors">
          <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </button>
      </div>
    </div>

    <!-- Player Controls -->
    <div 
      class="absolute bottom-0 left-0 right-0 z-10 p-4 bg-gradient-to-t from-black/80 to-transparent transition-opacity duration-300"
      :class="{ 'opacity-0': !showControls }"
    >
      <!-- Progress Bar -->
      <div class="mb-4">
        <input
          type="range"
          :value="currentTime"
          :max="duration"
          @input="seek"
          class="w-full h-1 bg-dark-600 rounded-full appearance-none cursor-pointer
                 [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:w-4 [&::-webkit-slider-thumb]:h-4
                 [&::-webkit-slider-thumb]:bg-accent-500 [&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:cursor-pointer"
        />
        <div class="flex justify-between text-xs text-dark-400 mt-1">
          <span>{{ formatTime(currentTime) }}</span>
          <span>{{ formatTime(duration) }}</span>
        </div>
      </div>

      <!-- Control Buttons -->
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <!-- Play/Pause -->
          <button @click.stop="togglePlay" class="text-white hover:text-accent-400 transition-colors">
            <svg v-if="isPlaying" class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zM7 8a1 1 0 012 0v4a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v4a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd"/>
            </svg>
            <svg v-else class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
            </svg>
          </button>

          <!-- Skip buttons -->
          <button @click.stop="skip(-10)" class="text-white hover:text-accent-400 transition-colors">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.066 11.2a1 1 0 000 1.6l5.334 4A1 1 0 0019 16V8a1 1 0 00-1.6-.8l-5.333 4zM4.066 11.2a1 1 0 000 1.6l5.334 4A1 1 0 0011 16V8a1 1 0 00-1.6-.8l-5.334 4z"/>
            </svg>
          </button>
          <button @click.stop="skip(10)" class="text-white hover:text-accent-400 transition-colors">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.933 12.8a1 1 0 000-1.6L6.6 7.2A1 1 0 005 8v8a1 1 0 001.6.8l5.333-4zM19.933 12.8a1 1 0 000-1.6l-5.333-4A1 1 0 0013 8v8a1 1 0 001.6.8l5.333-4z"/>
            </svg>
          </button>

          <!-- Volume -->
          <div class="flex items-center gap-2">
            <button @click.stop="toggleMute" class="text-white hover:text-accent-400 transition-colors">
              <svg v-if="isMuted || volume === 0" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2"/>
              </svg>
              <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z"/>
              </svg>
            </button>
            <input
              type="range"
              v-model="volume"
              min="0"
              max="1"
              step="0.1"
              @input="setVolume"
              class="w-20 h-1 bg-dark-600 rounded-full appearance-none cursor-pointer"
            />
          </div>
        </div>

        <div class="flex items-center gap-4">
          <!-- Fullscreen -->
          <button @click.stop="toggleFullscreen" class="text-white hover:text-accent-400 transition-colors">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5l-5-5m5 5v-4m0 4h-4"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Error State -->
    <div v-if="error" class="absolute inset-0 flex items-center justify-center bg-dark-900">
      <div class="text-center">
        <svg class="w-16 h-16 text-red-500 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <h3 class="text-xl font-semibold text-white mb-2">Ошибка воспроизведения</h3>
        <p class="text-dark-400 mb-4">{{ error }}</p>
        <button @click="goBack" class="btn-primary">Назад</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useTorrentStore } from '../stores/torrent'

const route = useRoute()
const router = useRouter()
const torrentStore = useTorrentStore()

const videoRef = ref(null)
const isPlaying = ref(false)
const isBuffering = ref(true)
const isMuted = ref(false)
const volume = ref(1)
const currentTime = ref(0)
const duration = ref(0)
const showControls = ref(true)
const error = ref(null)
const streamUrl = ref('')

let hideControlsTimeout = null

const title = computed(() => route.query.title || 'Видео')
const hash = computed(() => route.query.hash || '')

const loadStreamUrl = async () => {
  if (hash.value) {
    try {
      streamUrl.value = await torrentStore.getStreamURL(hash.value, 0)
      console.log('Stream URL loaded:', streamUrl.value)
    } catch (err) {
      console.error('Failed to get stream URL:', err)
      error.value = 'Не удалось получить ссылку на поток'
    }
  }
}

const formatTime = (seconds) => {
  const h = Math.floor(seconds / 3600)
  const m = Math.floor((seconds % 3600) / 60)
  const s = Math.floor(seconds % 60)
  
  if (h > 0) {
    return `${h}:${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`
  }
  return `${m}:${String(s).padStart(2, '0')}`
}

const togglePlay = () => {
  if (!videoRef.value) return
  
  if (isPlaying.value) {
    videoRef.value.pause()
  } else {
    videoRef.value.play()
  }
  isPlaying.value = !isPlaying.value
}

const seek = (e) => {
  if (!videoRef.value) return
  videoRef.value.currentTime = parseFloat(e.target.value)
}

const skip = (seconds) => {
  if (!videoRef.value) return
  videoRef.value.currentTime = Math.max(0, Math.min(duration.value, currentTime.value + seconds))
}

const setVolume = () => {
  if (!videoRef.value) return
  videoRef.value.volume = volume.value
  isMuted.value = volume.value === 0
}

const toggleMute = () => {
  if (!videoRef.value) return
  isMuted.value = !isMuted.value
  videoRef.value.muted = isMuted.value
}

const toggleFullscreen = () => {
  if (document.fullscreenElement) {
    document.exitFullscreen()
  } else {
    document.documentElement.requestFullscreen()
  }
}

const handleTimeUpdate = () => {
  if (!videoRef.value) return
  currentTime.value = videoRef.value.currentTime
}

const handleLoadedMetadata = () => {
  if (!videoRef.value) return
  duration.value = videoRef.value.duration
  isBuffering.value = false
}

const handleEnded = () => {
  isPlaying.value = false
}

const handleError = (e) => {
  console.error('Video error:', e)
  error.value = 'Не удалось воспроизвести видео. Проверьте, запущен ли TorrServer.'
}

const handleMouseMove = () => {
  showControls.value = true
  
  if (hideControlsTimeout) {
    clearTimeout(hideControlsTimeout)
  }
  
  hideControlsTimeout = setTimeout(() => {
    if (isPlaying.value) {
      showControls.value = false
    }
  }, 3000)
}

const goBack = () => {
  router.back()
}

const handleKeydown = (e) => {
  switch (e.key) {
    case ' ':
    case 'k':
      e.preventDefault()
      togglePlay()
      break
    case 'ArrowLeft':
      skip(-10)
      break
    case 'ArrowRight':
      skip(10)
      break
    case 'ArrowUp':
      volume.value = Math.min(1, volume.value + 0.1)
      setVolume()
      break
    case 'ArrowDown':
      volume.value = Math.max(0, volume.value - 0.1)
      setVolume()
      break
    case 'm':
      toggleMute()
      break
    case 'f':
      toggleFullscreen()
      break
    case 'Escape':
      goBack()
      break
  }
}

onMounted(async () => {
  document.addEventListener('keydown', handleKeydown)
  await loadStreamUrl()
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
  if (hideControlsTimeout) {
    clearTimeout(hideControlsTimeout)
  }
})
</script>
