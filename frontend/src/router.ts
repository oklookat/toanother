//
import YandexMusic from './routes/ym/index.svelte'
import YandexMusicSettings from './routes/ym/settings.svelte'
import YandexMusicPlaylists from './routes/ym/playlists.svelte'
import YandexMusicAlbums from './routes/ym/albums.svelte'
import YandexMusicArtists from './routes/ym/artists.svelte'
//
import Spotify from './routes/spotify/index.svelte'
import SpotifySettings from './routes/spotify/settings.svelte'
import SpotifyImport from './routes/spotify/import.svelte'

const routes = {
    //
    '/ym': YandexMusic,
    '/ym/settings': YandexMusicSettings,
    '/ym/playlists': YandexMusicPlaylists,
    '/ym/albums': YandexMusicAlbums,
    '/ym/artists': YandexMusicArtists,
    //
    '/spotify': Spotify,
    '/spotify/settings': SpotifySettings,
    '/spotify/import': SpotifyImport,
}

export default routes