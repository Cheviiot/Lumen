<template>
  <div
    v-if="visible"
    class="fixed inset-0 z-50 flex items-center justify-center p-4"
    @click.self="close"
  >
    <!-- Backdrop -->
    <div class="absolute inset-0 bg-black/80 backdrop-blur-sm" />
    
    <!-- Modal Content -->
    <div class="relative bg-dark-800 rounded-xl max-w-3xl w-full max-h-[90vh] overflow-hidden shadow-2xl animate-fade-in">
      <!-- Header -->
      <div class="flex items-center gap-3 px-4 py-2.5 border-b border-dark-700">
        <button
          @click="close"
          class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-dark-700 rounded-lg transition-colors text-dark-400 hover:text-white text-sm"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
          Назад
        </button>
        <h3 class="text-base font-semibold text-white flex-1 text-center">{{ title }}</h3>
        <div class="w-[76px]"></div>
      </div>
      
      <!-- Body -->
      <div class="p-4 overflow-y-auto max-h-[calc(90vh-100px)]">
        <slot />
      </div>
      
      <!-- Footer -->
      <div v-if="$slots.footer" class="p-4 border-t border-dark-700">
        <slot name="footer" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { watch, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['close'])

const close = () => {
  emit('close')
}

const handleEscape = (e) => {
  if (e.key === 'Escape' && props.visible) {
    close()
  }
}

watch(() => props.visible, (val) => {
  if (val) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
})

onMounted(() => {
  document.addEventListener('keydown', handleEscape)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscape)
  document.body.style.overflow = ''
})
</script>
