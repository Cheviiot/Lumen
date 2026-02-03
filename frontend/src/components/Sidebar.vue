<template>
  <aside
    :class="[
      'h-full bg-dark-925/95 backdrop-blur-xl border-r border-dark-850 flex flex-col transition-all duration-300 ease-in-out',
      isCollapsed ? 'w-11' : 'w-48'
    ]"
  >
    <!-- Navigation Items -->
    <nav :class="['flex-1 py-2 space-y-0.5 overflow-y-auto', isCollapsed ? 'px-1' : 'px-2']">
      <router-link
        v-for="item in navItems"
        :key="item.path"
        :to="item.path"
        custom
        v-slot="{ isActive, navigate }"
      >
        <button
          @click="navigate"
          :class="[
            'w-full flex items-center gap-2.5 py-2 rounded-lg transition-all duration-200',
            isCollapsed ? 'justify-center px-1' : 'px-2.5',
            isActive
              ? 'bg-accent-600/10 text-accent-500 border border-accent-600/20'
              : 'text-dark-500 hover:bg-dark-800 hover:text-white border border-transparent'
          ]"
          :title="isCollapsed ? item.label : ''"
        >
          <component :is="item.icon" class="w-4 h-4 flex-shrink-0" />
          <span
            v-if="!isCollapsed"
            class="text-sm font-medium whitespace-nowrap overflow-hidden"
          >
            {{ item.label }}
          </span>
        </button>
      </router-link>
    </nav>

    <!-- Toggle Button -->
    <div :class="['border-t border-dark-850', isCollapsed ? 'p-1' : 'p-2']">
      <button
        @click="toggleSidebar"
        :class="[
          'w-full flex items-center justify-center gap-2.5 py-2 rounded-lg text-dark-500 hover:bg-dark-800 hover:text-white transition-all duration-200 border border-transparent hover:border-dark-700',
          isCollapsed ? 'px-1' : 'px-2.5'
        ]"
        :title="isCollapsed ? 'Развернуть' : 'Свернуть'"
      >
        <svg
          class="w-4 h-4 transition-transform duration-300"
          :class="{ 'rotate-180': !isCollapsed }"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5l7 7-7 7M5 5l7 7-7 7" />
        </svg>
        <span v-if="!isCollapsed" class="text-sm font-medium">Свернуть</span>
      </button>
    </div>
  </aside>
</template>

<script setup>
import { computed, h } from 'vue'
import { useAppStore } from '../stores/app'

const appStore = useAppStore()
const isCollapsed = computed(() => appStore.sidebarCollapsed)

const toggleSidebar = () => {
  appStore.toggleSidebar()
}

// Icon components
const HomeIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6'
      })
    ])
  }
}

const MovieIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M7 4v16M17 4v16M3 8h4m10 0h4M3 12h18M3 16h4m10 0h4M4 20h16a1 1 0 001-1V5a1 1 0 00-1-1H4a1 1 0 00-1 1v14a1 1 0 001 1z'
      })
    ])
  }
}

const TVIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z'
      })
    ])
  }
}

const CartoonIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z'
      })
    ])
  }
}

const HeartIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z'
      })
    ])
  }
}

const HistoryIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
      })
    ])
  }
}

const BookmarkIcon = {
  render() {
    return h('svg', { fill: 'none', stroke: 'currentColor', viewBox: '0 0 24 24' }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        'stroke-width': '2',
        d: 'M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z'
      })
    ])
  }
}

const navItems = [
  { path: '/', label: 'Главная', icon: HomeIcon },
  { path: '/movies', label: 'Фильмы', icon: MovieIcon },
  { path: '/series', label: 'Сериалы', icon: TVIcon },
  { path: '/cartoons', label: 'Мультфильмы', icon: CartoonIcon },
  { path: '/favorites', label: 'Избранные', icon: HeartIcon },
  { path: '/history', label: 'История', icon: HistoryIcon },
  { path: '/bookmarks', label: 'Закладки', icon: BookmarkIcon }
]
</script>
