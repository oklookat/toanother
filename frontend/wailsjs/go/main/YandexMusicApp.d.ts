// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {base} from '../models';

export function GetPlaylists():Promise<Array<base.Playlist>|Error>;

export function ApplySettings(arg1:base.YandexMusicSettings):Promise<Error>;

export function DownloadPlaylists():Promise<Array<base.Playlist>|Error>;

export function GetAlbums():Promise<Array<base.Album>|Error>;

export function GetArtists():Promise<Array<base.Artist>|Error>;

export function DownloadAlbums():Promise<Array<base.Album>|Error>;

export function DownloadArtists():Promise<Array<base.Artist>|Error>;

export function GetSettings():Promise<base.YandexMusicSettings>;

export function GetTracks(arg1:number):Promise<Array<base.Track>|Error>;
