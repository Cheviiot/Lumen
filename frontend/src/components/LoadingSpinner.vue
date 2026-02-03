<template>
  <div class="flex items-center justify-center h-full">
    <div class="flex flex-col items-center gap-4">
      <div class="relative">
        <div class="w-12 h-12 border-2 border-dark-800 rounded-full"></div>
        <div class="absolute top-0 left-0 w-12 h-12 border-2 border-accent-500 rounded-full border-t-transparent animate-spin"></div>
      </div>
      <p class="text-dark-500 text-sm">{{ text }}</p>
    </div>
  </div>
</template>

<script setup>
defineProps({
  text: {
    type: String,
    default: 'Загрузка...'
  }
})
</script>
