<template>
  <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
    <MediaCard
      v-for="item in items"
      :key="item.id"
      :media="item"
      :type="type"
      @click="handleClick"
    />
  </div>
  
  <!-- Load More -->
  <div v-if="hasMore" class="flex justify-center mt-8">
    <button
      @click="$emit('loadMore')"
      class="btn-primary flex items-center gap-2"
      :disabled="loading"
    >
      <svg v-if="loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
      </svg>
      {{ loading ? 'Загрузка...' : 'Загрузить ещё' }}
    </button>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import MediaCard from './MediaCard.vue'

const props = defineProps({
  items: {
    type: Array,
    default: () => []
  },
  type: {
    type: String,
    default: 'movie'
  },
  hasMore: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  }
})

defineEmits(['loadMore'])

const router = useRouter()

const handleClick = (media) => {
  const routeName = props.type === 'tv' ? 'TVDetails' : 'MovieDetails'
  router.push({ name: routeName, params: { id: media.id } })
}
</script>
