import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/HomeView.vue')
  },
  {
    path: '/movies',
    name: 'Movies',
    component: () => import('../views/MoviesView.vue')
  },
  {
    path: '/series',
    name: 'Series',
    component: () => import('../views/SeriesView.vue')
  },
  {
    path: '/cartoons',
    name: 'Cartoons',
    component: () => import('../views/CartoonsView.vue')
  },
  {
    path: '/favorites',
    name: 'Favorites',
    component: () => import('../views/FavoritesView.vue')
  },
  {
    path: '/history',
    name: 'History',
    component: () => import('../views/HistoryView.vue')
  },
  {
    path: '/bookmarks',
    name: 'Bookmarks',
    component: () => import('../views/BookmarksView.vue')
  },
  {
    path: '/movie/:id',
    name: 'MovieDetails',
    component: () => import('../views/MovieDetailsView.vue')
  },
  {
    path: '/tv/:id',
    name: 'TVDetails',
    component: () => import('../views/TVDetailsView.vue')
  },
  {
    path: '/search',
    name: 'Search',
    component: () => import('../views/SearchView.vue')
  },
  {
    path: '/settings',
    name: 'Settings',
    component: () => import('../views/SettingsView.vue')
  },
  {
    path: '/player',
    name: 'Player',
    component: () => import('../views/PlayerView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
