export namespace base {
	
	export class Artist {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Artist(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class Playlist {
	    id: number;
	    isLikedTracks: boolean;
	    title: string;
	    trackCount: number;
	
	    static createFrom(source: any = {}) {
	        return new Playlist(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.isLikedTracks = source["isLikedTracks"];
	        this.title = source["title"];
	        this.trackCount = source["trackCount"];
	    }
	}
	export class Track {
	    id: number;
	    playlistID: number;
	    title: string;
	    albumTitle?: string;
	    durationMs: number;
	    artist: string[];
	
	    static createFrom(source: any = {}) {
	        return new Track(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.playlistID = source["playlistID"];
	        this.title = source["title"];
	        this.albumTitle = source["albumTitle"];
	        this.durationMs = source["durationMs"];
	        this.artist = source["artist"];
	    }
	}
	export class SpotifySettings {
	    id: string;
	    secret: string;
	
	    static createFrom(source: any = {}) {
	        return new SpotifySettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.secret = source["secret"];
	    }
	}
	export class YandexMusicSettings {
	    login: string;
	
	    static createFrom(source: any = {}) {
	        return new YandexMusicSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.login = source["login"];
	    }
	}
	export class Album {
	    id: number;
	    title: string;
	    artist: string[];
	    releaseDate: number;
	    trackCount: number;
	    year: number;
	
	    static createFrom(source: any = {}) {
	        return new Album(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.artist = source["artist"];
	        this.releaseDate = source["releaseDate"];
	        this.trackCount = source["trackCount"];
	        this.year = source["year"];
	    }
	}

}

