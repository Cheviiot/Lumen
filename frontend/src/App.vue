<template>
  <div class="h-screen flex flex-col overflow-hidden">
    <!-- Custom Titlebar -->
    <Titlebar @search="handleSearch" />
    
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <Sidebar />
      
      <!-- Main Content -->
      <main class="flex-1 overflow-y-auto bg-dark-900">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </main>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import Titlebar from './components/Titlebar.vue'
import Sidebar from './components/Sidebar.vue'
import { useAppStore } from './stores/app'

const appStore = useAppStore()

const handleSearch = (query) => {
  appStore.setSearchQuery(query)
}

onMounted(async () => {
  await appStore.init()
})
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
