export namespace main {
	
	export class Bookmark {
	    id: number;
	    tmdb_id: number;
	    media_type: string;
	    title: string;
	    poster: string;
	    year: string;
	    note?: string;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Bookmark(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.tmdb_id = source["tmdb_id"];
	        this.media_type = source["media_type"];
	        this.title = source["title"];
	        this.poster = source["poster"];
	        this.year = source["year"];
	        this.note = source["note"];
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Collection {
	    id: number;
	    name: string;
	    poster_path: string;
	    backdrop_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Collection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	    }
	}
	export class Movie {
	    id: number;
	    title: string;
	    original_title: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    release_date: string;
	    vote_average: number;
	    vote_count: number;
	    popularity: number;
	    genre_ids: number[];
	    adult: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Movie(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.original_title = source["original_title"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.release_date = source["release_date"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.popularity = source["popularity"];
	        this.genre_ids = source["genre_ids"];
	        this.adult = source["adult"];
	    }
	}
	export class CollectionDetails {
	    id: number;
	    name: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    parts: Movie[];
	
	    static createFrom(source: any = {}) {
	        return new CollectionDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.parts = this.convertValues(source["parts"], Movie);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Company {
	    id: number;
	    name: string;
	    logo_path: string;
	    origin_country: string;
	
	    static createFrom(source: any = {}) {
	        return new Company(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.logo_path = source["logo_path"];
	        this.origin_country = source["origin_country"];
	    }
	}
	export class Episode {
	    id: number;
	    episode_number: number;
	    season_number: number;
	    name: string;
	    overview: string;
	    still_path: string;
	    air_date: string;
	    vote_average: number;
	    runtime: number;
	
	    static createFrom(source: any = {}) {
	        return new Episode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.episode_number = source["episode_number"];
	        this.season_number = source["season_number"];
	        this.name = source["name"];
	        this.overview = source["overview"];
	        this.still_path = source["still_path"];
	        this.air_date = source["air_date"];
	        this.vote_average = source["vote_average"];
	        this.runtime = source["runtime"];
	    }
	}
	export class Favorite {
	    id: number;
	    tmdb_id: number;
	    media_type: string;
	    title: string;
	    poster: string;
	    year: string;
	    rating: number;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Favorite(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.tmdb_id = source["tmdb_id"];
	        this.media_type = source["media_type"];
	        this.title = source["title"];
	        this.poster = source["poster"];
	        this.year = source["year"];
	        this.rating = source["rating"];
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Genre {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Genre(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class History {
	    id: number;
	    tmdb_id: number;
	    media_type: string;
	    title: string;
	    poster: string;
	    episode?: string;
	    progress: number;
	    duration: number;
	    // Go type: time
	    watched_at: any;
	    // Go type: time
	    last_watched: any;
	
	    static createFrom(source: any = {}) {
	        return new History(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.tmdb_id = source["tmdb_id"];
	        this.media_type = source["media_type"];
	        this.title = source["title"];
	        this.poster = source["poster"];
	        this.episode = source["episode"];
	        this.progress = source["progress"];
	        this.duration = source["duration"];
	        this.watched_at = this.convertValues(source["watched_at"], null);
	        this.last_watched = this.convertValues(source["last_watched"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IndexerConfig {
	    id: string;
	    name: string;
	    type: string;
	    configured: boolean;
	    language: string;
	
	    static createFrom(source: any = {}) {
	        return new IndexerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.configured = source["configured"];
	        this.language = source["language"];
	    }
	}
	export class JackettIndexer {
	    id: string;
	    name: string;
	    status: number;
	    results: number;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new JackettIndexer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.status = source["status"];
	        this.results = source["results"];
	        this.error = source["error"];
	    }
	}
	export class JackettResult {
	    Title: string;
	    Tracker: string;
	    TrackerId: string;
	    CategoryDesc: string;
	    PublishDate: string;
	    Size: number;
	    Seeders: number;
	    Peers: number;
	    Grabs: number;
	    MagnetUri: string;
	    Link: string;
	    Guid: string;
	    Details: string;
	    InfoHash: string;
	    DownloadVolumeFactor: number;
	    UploadVolumeFactor: number;
	
	    static createFrom(source: any = {}) {
	        return new JackettResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Tracker = source["Tracker"];
	        this.TrackerId = source["TrackerId"];
	        this.CategoryDesc = source["CategoryDesc"];
	        this.PublishDate = source["PublishDate"];
	        this.Size = source["Size"];
	        this.Seeders = source["Seeders"];
	        this.Peers = source["Peers"];
	        this.Grabs = source["Grabs"];
	        this.MagnetUri = source["MagnetUri"];
	        this.Link = source["Link"];
	        this.Guid = source["Guid"];
	        this.Details = source["Details"];
	        this.InfoHash = source["InfoHash"];
	        this.DownloadVolumeFactor = source["DownloadVolumeFactor"];
	        this.UploadVolumeFactor = source["UploadVolumeFactor"];
	    }
	}
	export class JackettResponse {
	    Results: JackettResult[];
	    Indexers: JackettIndexer[];
	
	    static createFrom(source: any = {}) {
	        return new JackettResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Results = this.convertValues(source["Results"], JackettResult);
	        this.Indexers = this.convertValues(source["Indexers"], JackettIndexer);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class MovieDetails {
	    id: number;
	    title: string;
	    original_title: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    release_date: string;
	    vote_average: number;
	    vote_count: number;
	    popularity: number;
	    genre_ids: number[];
	    adult: boolean;
	    runtime: number;
	    genres: Genre[];
	    tagline: string;
	    status: string;
	    budget: number;
	    revenue: number;
	    imdb_id: string;
	    homepage: string;
	    production_companies: Company[];
	    belongs_to_collection?: Collection;
	
	    static createFrom(source: any = {}) {
	        return new MovieDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.original_title = source["original_title"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.release_date = source["release_date"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.popularity = source["popularity"];
	        this.genre_ids = source["genre_ids"];
	        this.adult = source["adult"];
	        this.runtime = source["runtime"];
	        this.genres = this.convertValues(source["genres"], Genre);
	        this.tagline = source["tagline"];
	        this.status = source["status"];
	        this.budget = source["budget"];
	        this.revenue = source["revenue"];
	        this.imdb_id = source["imdb_id"];
	        this.homepage = source["homepage"];
	        this.production_companies = this.convertValues(source["production_companies"], Company);
	        this.belongs_to_collection = this.convertValues(source["belongs_to_collection"], Collection);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SearchResult {
	    page: number;
	    total_pages: number;
	    total_results: number;
	    results: Movie[];
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.total_pages = source["total_pages"];
	        this.total_results = source["total_results"];
	        this.results = this.convertValues(source["results"], Movie);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Season {
	    id: number;
	    season_number: number;
	    name: string;
	    overview: string;
	    poster_path: string;
	    air_date: string;
	    episode_count: number;
	
	    static createFrom(source: any = {}) {
	        return new Season(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.season_number = source["season_number"];
	        this.name = source["name"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.air_date = source["air_date"];
	        this.episode_count = source["episode_count"];
	    }
	}
	export class SeasonDetails {
	    id: number;
	    season_number: number;
	    name: string;
	    overview: string;
	    poster_path: string;
	    air_date: string;
	    episodes: Episode[];
	
	    static createFrom(source: any = {}) {
	        return new SeasonDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.season_number = source["season_number"];
	        this.name = source["name"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.air_date = source["air_date"];
	        this.episodes = this.convertValues(source["episodes"], Episode);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Settings {
	    ID: number;
	    torr_server_url: string;
	    jackett_url: string;
	    jackett_api_key: string;
	    jackett_public_parser: string;
	    supabase_url: string;
	    supabase_key: string;
	    theme: string;
	    language: string;
	    auto_sync: boolean;
	    download_path: string;
	    player_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.torr_server_url = source["torr_server_url"];
	        this.jackett_url = source["jackett_url"];
	        this.jackett_api_key = source["jackett_api_key"];
	        this.jackett_public_parser = source["jackett_public_parser"];
	        this.supabase_url = source["supabase_url"];
	        this.supabase_key = source["supabase_key"];
	        this.theme = source["theme"];
	        this.language = source["language"];
	        this.auto_sync = source["auto_sync"];
	        this.download_path = source["download_path"];
	        this.player_path = source["player_path"];
	    }
	}
	export class TVShow {
	    id: number;
	    name: string;
	    original_name: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    first_air_date: string;
	    vote_average: number;
	    vote_count: number;
	    popularity: number;
	    genre_ids: number[];
	
	    static createFrom(source: any = {}) {
	        return new TVShow(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.original_name = source["original_name"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.first_air_date = source["first_air_date"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.popularity = source["popularity"];
	        this.genre_ids = source["genre_ids"];
	    }
	}
	export class TVSearchResult {
	    page: number;
	    total_pages: number;
	    total_results: number;
	    results: TVShow[];
	
	    static createFrom(source: any = {}) {
	        return new TVSearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.total_pages = source["total_pages"];
	        this.total_results = source["total_results"];
	        this.results = this.convertValues(source["results"], TVShow);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class TVShowDetails {
	    id: number;
	    name: string;
	    original_name: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    first_air_date: string;
	    vote_average: number;
	    vote_count: number;
	    popularity: number;
	    genre_ids: number[];
	    number_of_seasons: number;
	    number_of_episodes: number;
	    genres: Genre[];
	    tagline: string;
	    status: string;
	    episode_run_time: number[];
	    seasons: Season[];
	    homepage: string;
	    production_companies: Company[];
	
	    static createFrom(source: any = {}) {
	        return new TVShowDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.original_name = source["original_name"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.first_air_date = source["first_air_date"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.popularity = source["popularity"];
	        this.genre_ids = source["genre_ids"];
	        this.number_of_seasons = source["number_of_seasons"];
	        this.number_of_episodes = source["number_of_episodes"];
	        this.genres = this.convertValues(source["genres"], Genre);
	        this.tagline = source["tagline"];
	        this.status = source["status"];
	        this.episode_run_time = source["episode_run_time"];
	        this.seasons = this.convertValues(source["seasons"], Season);
	        this.homepage = source["homepage"];
	        this.production_companies = this.convertValues(source["production_companies"], Company);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Torrent {
	    title: string;
	    hash: string;
	    poster: string;
	    data: string;
	    timestamp: number;
	    size: number;
	    category: string;
	    name: string;
	    stat: number;
	
	    static createFrom(source: any = {}) {
	        return new Torrent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.hash = source["hash"];
	        this.poster = source["poster"];
	        this.data = source["data"];
	        this.timestamp = source["timestamp"];
	        this.size = source["size"];
	        this.category = source["category"];
	        this.name = source["name"];
	        this.stat = source["stat"];
	    }
	}
	export class TorrentFile {
	    id: number;
	    path: string;
	    length: number;
	
	    static createFrom(source: any = {}) {
	        return new TorrentFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.path = source["path"];
	        this.length = source["length"];
	    }
	}
	export class TorrentStat {
	    hash: string;
	    title: string;
	    torrent_size: number;
	    download_speed: number;
	    upload_speed: number;
	    active_peers: number;
	    total_peers: number;
	    connected_seeders: number;
	    file_stats: TorrentFile[];
	
	    static createFrom(source: any = {}) {
	        return new TorrentStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hash = source["hash"];
	        this.title = source["title"];
	        this.torrent_size = source["torrent_size"];
	        this.download_speed = source["download_speed"];
	        this.upload_speed = source["upload_speed"];
	        this.active_peers = source["active_peers"];
	        this.total_peers = source["total_peers"];
	        this.connected_seeders = source["connected_seeders"];
	        this.file_stats = this.convertValues(source["file_stats"], TorrentFile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Video {
	    id: string;
	    key: string;
	    name: string;
	    site: string;
	    size: number;
	    type: string;
	    official: boolean;
	    published_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.key = source["key"];
	        this.name = source["name"];
	        this.site = source["site"];
	        this.size = source["size"];
	        this.type = source["type"];
	        this.official = source["official"];
	        this.published_at = source["published_at"];
	    }
	}
	export class VideosResult {
	    id: number;
	    results: Video[];
	
	    static createFrom(source: any = {}) {
	        return new VideosResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.results = this.convertValues(source["results"], Video);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

