import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMediaStore = defineStore('media', () => {
  const trendingMovies = ref([])
  const trendingTV = ref([])
  const popularMovies = ref([])
  const popularTV = ref([])
  const animationMovies = ref([])
  const animationTV = ref([])
  const searchResults = ref([])
  const favorites = ref([])
  const history = ref([])
  const bookmarks = ref([])
  const isLoading = ref(false)

  async function fetchTrendingMovies(page = 1) {
    console.log('fetchTrendingMovies called, window.go:', !!window.go)
    console.log('TMDB available:', !!window.go?.main?.TMDB)
    console.log('GetTrendingMovies available:', !!window.go?.main?.TMDB?.GetTrendingMovies)
    
    if (window.go?.main?.TMDB?.GetTrendingMovies) {
      try {
        const result = await window.go.main.TMDB.GetTrendingMovies('week', page)
        console.log('GetTrendingMovies result:', result)
        trendingMovies.value = result?.results || []
        return result
      } catch (err) {
        console.error('GetTrendingMovies error:', err)
        return null
      }
    }
    console.log('GetTrendingMovies NOT available!')
    return null
  }

  async function fetchTrendingTV(page = 1) {
    if (window.go?.main?.TMDB?.GetTrendingTV) {
      try {
        const result = await window.go.main.TMDB.GetTrendingTV('week', page)
        console.log('GetTrendingTV result:', result)
        trendingTV.value = result?.results || []
        return result
      } catch (err) {
        console.error('GetTrendingTV error:', err)
        return null
      }
    }
    return null
  }

  async function fetchPopularMovies(page = 1) {
    if (window.go?.main?.TMDB?.GetPopularMovies) {
      const result = await window.go.main.TMDB.GetPopularMovies(page)
      popularMovies.value = result?.results || []
      return result
    }
    return null
  }

  async function fetchPopularTV(page = 1) {
    if (window.go?.main?.TMDB?.GetPopularTV) {
      const result = await window.go.main.TMDB.GetPopularTV(page)
      popularTV.value = result?.results || []
      return result
    }
    return null
  }

  async function fetchAnimationMovies(page = 1) {
    if (window.go?.main?.TMDB?.GetAnimationMovies) {
      const result = await window.go.main.TMDB.GetAnimationMovies(page)
      animationMovies.value = result?.results || []
      return result
    }
    return null
  }

  async function fetchAnimationTV(page = 1) {
    if (window.go?.main?.TMDB?.GetAnimationTV) {
      const result = await window.go.main.TMDB.GetAnimationTV(page)
      animationTV.value = result?.results || []
      return result
    }
    return null
  }

  async function searchMovies(query, page = 1) {
    if (window.go?.main?.TMDB?.SearchMovies) {
      const result = await window.go.main.TMDB.SearchMovies(query, page)
      return result
    }
    return null
  }

  async function searchTV(query, page = 1) {
    if (window.go?.main?.TMDB?.SearchTV) {
      const result = await window.go.main.TMDB.SearchTV(query, page)
      return result
    }
    return null
  }

  async function getMovieDetails(movieId) {
    if (window.go?.main?.TMDB?.GetMovieDetails) {
      return await window.go.main.TMDB.GetMovieDetails(movieId)
    }
    return null
  }

  async function getTVDetails(tvId) {
    if (window.go?.main?.TMDB?.GetTVDetails) {
      return await window.go.main.TMDB.GetTVDetails(tvId)
    }
    return null
  }

  async function getSeasonDetails(tvId, seasonNumber) {
    if (window.go?.main?.TMDB?.GetSeasonDetails) {
      return await window.go.main.TMDB.GetSeasonDetails(tvId, seasonNumber)
    }
    return null
  }

  // Favorites
  async function loadFavorites() {
    if (window.go?.main?.Database?.GetFavorites) {
      favorites.value = await window.go.main.Database.GetFavorites() || []
    }
  }

  async function addToFavorites(item) {
    if (window.go?.main?.Database?.AddFavorite) {
      await window.go.main.Database.AddFavorite(item)
      await loadFavorites()
    }
  }

  async function removeFromFavorites(tmdbId, mediaType) {
    if (window.go?.main?.Database?.RemoveFavorite) {
      await window.go.main.Database.RemoveFavorite(tmdbId, mediaType)
      await loadFavorites()
    }
  }

  async function isFavorite(tmdbId, mediaType) {
    if (window.go?.main?.Database?.IsFavorite) {
      return await window.go.main.Database.IsFavorite(tmdbId, mediaType)
    }
    return false
  }

  // History
  async function loadHistory() {
    if (window.go?.main?.Database?.GetHistory) {
      history.value = await window.go.main.Database.GetHistory() || []
    }
  }

  async function addToHistory(item) {
    if (window.go?.main?.Database?.AddHistory) {
      await window.go.main.Database.AddHistory(item)
      await loadHistory()
    }
  }

  async function clearHistory() {
    if (window.go?.main?.Database?.ClearHistory) {
      await window.go.main.Database.ClearHistory()
      history.value = []
    }
  }

  // Bookmarks
  async function loadBookmarks() {
    if (window.go?.main?.Database?.GetBookmarks) {
      bookmarks.value = await window.go.main.Database.GetBookmarks() || []
    }
  }

  async function addToBookmarks(item) {
    if (window.go?.main?.Database?.AddBookmark) {
      await window.go.main.Database.AddBookmark(item)
      await loadBookmarks()
    }
  }

  async function removeFromBookmarks(tmdbId, mediaType) {
    if (window.go?.main?.Database?.RemoveBookmark) {
      await window.go.main.Database.RemoveBookmark(tmdbId, mediaType)
      await loadBookmarks()
    }
  }

  function getImageURL(path, size = 'w500') {
    if (!path) return '/placeholder.png'
    
    // По умолчанию прокси ВКЛЮЧЕН (если значение не 'false')
    const proxyEnabled = localStorage.getItem('tmdb_proxy_enabled') !== 'false'
    const proxyImages = localStorage.getItem('tmdb_proxy_images') !== 'false'
    
    if (proxyEnabled && proxyImages) {
      return `https://tmdbimg.rootu.top/t/p/${size}${path}`
    }
    
    return `https://image.tmdb.org/t/p/${size}${path}`
  }

  function getTmdbApiUrl(endpoint) {
    // По умолчанию прокси ВКЛЮЧЕН (если значение не 'false')
    const proxyEnabled = localStorage.getItem('tmdb_proxy_enabled') !== 'false'
    const proxyApi = localStorage.getItem('tmdb_proxy_api') !== 'false'
    
    if (proxyEnabled && proxyApi) {
      return `https://tmdbapi.rootu.top/3/${endpoint}`
    }
    
    return `https://api.themoviedb.org/3/${endpoint}`
  }

  return {
    trendingMovies,
    trendingTV,
    popularMovies,
    popularTV,
    animationMovies,
    animationTV,
    searchResults,
    favorites,
    history,
    bookmarks,
    isLoading,
    fetchTrendingMovies,
    fetchTrendingTV,
    fetchPopularMovies,
    fetchPopularTV,
    fetchAnimationMovies,
    fetchAnimationTV,
    searchMovies,
    searchTV,
    getMovieDetails,
    getTVDetails,
    getSeasonDetails,
    loadFavorites,
    addToFavorites,
    removeFromFavorites,
    isFavorite,
    loadHistory,
    addToHistory,
    clearHistory,
    loadBookmarks,
    addToBookmarks,
    removeFromBookmarks,
    getImageURL,
    getTmdbApiUrl
  }
})
