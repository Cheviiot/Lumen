import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useTorrentStore = defineStore('torrent', () => {
  const torrents = ref([])
  const searchResults = ref([])
  const isSearching = ref(false)
  const currentTorrent = ref(null)

  // TorrServer functions
  async function checkTorrServer() {
    if (window.go?.main?.TorrServer?.Echo) {
      return await window.go.main.TorrServer.Echo()
    }
    return false
  }

  async function addTorrent(magnetLink, title, poster) {
    if (window.go?.main?.TorrServer?.AddTorrent) {
      return await window.go.main.TorrServer.AddTorrent(magnetLink, title, poster, true)
    }
    return null
  }

  async function getTorrents() {
    if (window.go?.main?.TorrServer?.GetTorrents) {
      torrents.value = await window.go.main.TorrServer.GetTorrents() || []
      return torrents.value
    }
    return []
  }

  async function removeTorrent(hash) {
    if (window.go?.main?.TorrServer?.RemoveTorrent) {
      await window.go.main.TorrServer.RemoveTorrent(hash)
      await getTorrents()
    }
  }

  async function getStreamURL(hash, fileIndex) {
    if (window.go?.main?.TorrServer?.GetStreamURL) {
      return await window.go.main.TorrServer.GetStreamURL(hash, fileIndex)
    }
    return ''
  }

  async function preloadTorrent(hash, fileIndex) {
    if (window.go?.main?.TorrServer?.PreloadTorrent) {
      await window.go.main.TorrServer.PreloadTorrent(hash, fileIndex)
    }
  }

  // Jackett functions
  async function searchJackett(query, type = 'all') {
    isSearching.value = true
    try {
      let result = null
      
      if (type === 'movie' && window.go?.main?.Jackett?.SearchMovies) {
        result = await window.go.main.Jackett.SearchMovies(query)
      } else if (type === 'tv' && window.go?.main?.Jackett?.SearchTV) {
        result = await window.go.main.Jackett.SearchTV(query)
      } else if (type === 'anime' && window.go?.main?.Jackett?.SearchAnime) {
        result = await window.go.main.Jackett.SearchAnime(query)
      } else if (window.go?.main?.Jackett?.Search) {
        result = await window.go.main.Jackett.Search(query, [])
      }

      searchResults.value = result?.Results || []
      return searchResults.value
    } finally {
      isSearching.value = false
    }
  }

  async function getIndexers() {
    if (window.go?.main?.Jackett?.GetIndexers) {
      return await window.go.main.Jackett.GetIndexers()
    }
    return []
  }

  async function testJackettConnection() {
    if (window.go?.main?.Jackett?.TestConnection) {
      return await window.go.main.Jackett.TestConnection()
    }
    return false
  }

  function formatSize(bytes) {
    const units = ['B', 'KB', 'MB', 'GB', 'TB']
    let size = bytes
    let unitIndex = 0
    
    while (size >= 1024 && unitIndex < units.length - 1) {
      size /= 1024
      unitIndex++
    }
    
    return `${size.toFixed(1)} ${units[unitIndex]}`
  }

  return {
    torrents,
    searchResults,
    isSearching,
    currentTorrent,
    checkTorrServer,
    addTorrent,
    getTorrents,
    removeTorrent,
    getStreamURL,
    preloadTorrent,
    searchJackett,
    getIndexers,
    testJackettConnection,
    formatSize
  }
})
