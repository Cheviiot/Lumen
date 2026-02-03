<template>
  <div class="relative">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-semibold text-white">{{ title }}</h2>
      <div class="flex gap-2">
        <button
          @click="scrollLeft"
          class="p-2 bg-dark-800 hover:bg-dark-700 rounded-full transition-colors"
          :disabled="!canScrollLeft"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
        </button>
        <button
          @click="scrollRight"
          class="p-2 bg-dark-800 hover:bg-dark-700 rounded-full transition-colors"
          :disabled="!canScrollRight"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
        </button>
      </div>
    </div>
    
    <div
      ref="scrollContainer"
      class="flex gap-4 overflow-x-auto scroll-smooth py-6 -my-6 hide-scrollbar"
      @scroll="updateScrollState"
    >
      <div
        v-for="item in items"
        :key="item.id"
        class="flex-shrink-0 w-44 relative z-0 hover:z-10"
      >
        <MediaCard
          :media="item"
          :type="type"
          @click="handleClick"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import MediaCard from './MediaCard.vue'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  items: {
    type: Array,
    default: () => []
  },
  type: {
    type: String,
    default: 'movie'
  }
})

const router = useRouter()
const scrollContainer = ref(null)
const canScrollLeft = ref(false)
const canScrollRight = ref(true)

const scrollLeft = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollBy({ left: -600, behavior: 'smooth' })
  }
}

const scrollRight = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollBy({ left: 600, behavior: 'smooth' })
  }
}

const updateScrollState = () => {
  if (scrollContainer.value) {
    const { scrollLeft, scrollWidth, clientWidth } = scrollContainer.value
    canScrollLeft.value = scrollLeft > 0
    canScrollRight.value = scrollLeft + clientWidth < scrollWidth - 10
  }
}

const handleClick = (media) => {
  const routeName = props.type === 'tv' ? 'TVDetails' : 'MovieDetails'
  router.push({ name: routeName, params: { id: media.id } })
}

onMounted(() => {
  updateScrollState()
})
</script>

<style scoped>
.hide-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.hide-scrollbar::-webkit-scrollbar {
  display: none;
}
</style>
