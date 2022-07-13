import { MessageError, MessageInfo } from ".";
import { ApplySettings, GetSettings, WebAuth, Ping, ImportLikedTracks } from "../../wailsjs/go/main/SpotifyApp";
import type { base } from "../../wailsjs/go/models";

export default class Spotify {
    private static onError = async (err: any) => {
        await MessageError(err);
    }

    public static async getSettings(): Promise<base.SpotifySettings> {
        try {
            const result = await GetSettings();
            if (!result) {
                return {
                    id: '',
                    secret: '',
                }
            }
            return result
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async applySettings(cfg: base.SpotifySettings): Promise<void> {
        try {
            await ApplySettings(cfg);
            await MessageInfo("Готово.");
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async webAuth(): Promise<void> {
        try {
            await WebAuth()
            await MessageInfo("Вход выполнен.");
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async ping(): Promise<void> {
        try {
            await Ping()
            await MessageInfo("Всё нормально.");
        } catch (err) {
            await this.onError(err)
        }
    }

    public static async importLikedTracks(tracks: base.Track[]): Promise<void> {
        try {
            await ImportLikedTracks(tracks)
            await MessageInfo("Готово.");
        } catch (err) {
            await this.onError(err)
        }
    }
}