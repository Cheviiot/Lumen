<template>
  <div class="bg-dark-800 rounded-lg p-4 hover:bg-dark-700/80 transition-colors cursor-pointer" @click="$emit('select', torrent)">
    <div class="flex items-start justify-between gap-4">
      <div class="flex-1 min-w-0">
        <h4 class="text-white font-medium text-sm line-clamp-2 mb-2">{{ torrent.Title }}</h4>
        <div class="flex flex-wrap items-center gap-3 text-xs text-dark-400">
          <span class="flex items-center gap-1">
            <svg class="w-4 h-4 text-accent-500" fill="currentColor" viewBox="0 0 20 20">
              <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"/>
              <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd"/>
            </svg>
            {{ torrent.Tracker }}
          </span>
          <span class="flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
            </svg>
            {{ formattedSize }}
          </span>
          <span class="flex items-center gap-1 text-green-400">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
            </svg>
            {{ torrent.Seeders }}
          </span>
          <span class="flex items-center gap-1 text-red-400">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M14.707 12.293a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 111.414-1.414L9 14.586V3a1 1 0 012 0v11.586l2.293-2.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
            </svg>
            {{ torrent.Peers }}
          </span>
        </div>
      </div>
      <button
        @click.stop="$emit('play', torrent)"
        class="flex-shrink-0 p-3 bg-accent-600 hover:bg-accent-700 rounded-full transition-colors"
      >
        <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
          <path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useTorrentStore } from '../stores/torrent'

const props = defineProps({
  torrent: {
    type: Object,
    required: true
  }
})

defineEmits(['select', 'play'])

const torrentStore = useTorrentStore()

const formattedSize = computed(() => {
  return torrentStore.formatSize(props.torrent.Size)
})
</script>
