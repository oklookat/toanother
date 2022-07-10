import { MessageError, MessageInfo } from ".";
import { ApplySettings, DownloadAlbums, DownloadArtists, DownloadPlaylists, GetAlbums, GetArtists, GetPlaylists, GetSettings, GetTracks } from "../../wailsjs/go/main/YandexMusicApp";
import type { base } from "../../wailsjs/go/models";

export default class YandexMusic {

    private static onError = async (err: any) => {
        await MessageError(err);
    }

    public static async getSettings(): Promise<base.YandexMusicSettings> {
        try {
            const result = await GetSettings();
            if(!result) {
                return {
                    login: ''
                }
            }
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async applySettings(cfg: base.YandexMusicSettings): Promise<void> {
        try {
            await ApplySettings(cfg);
            await MessageInfo("Ok");
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async getArtists(): Promise<base.Artist[]> {
        try {
            const result = (await GetArtists()) as base.Artist[]
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async downloadArtists(): Promise<base.Artist[]> {
        try {
            const result = (await DownloadArtists()) as base.Artist[];
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async getPlaylists(): Promise<base.Playlist[]> {
        try {
            const result = (await GetPlaylists()) as base.Playlist[];
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async downloadPlaylists(): Promise<base.Playlist[]> {
        try {
            const result = (await DownloadPlaylists()) as base.Playlist[];
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async getTracks(playlistID: number): Promise<base.Track[]> {
        try {
            const result = (await GetTracks(
                playlistID
            )) as base.Track[];
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async getAlbums(): Promise<base.Album[]> {
        try {
            const result = (await GetAlbums()) as base.Album[];
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async downloadAlbums(): Promise<base.Album[]> {
        try {
            const result = (await DownloadAlbums()) as base.Album[];
            return result
        } catch (err) {
            await this.onError(err)
        }
    }
}