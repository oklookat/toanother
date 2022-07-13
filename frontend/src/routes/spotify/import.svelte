<script lang="ts">
    import { onDestroy } from "svelte";

    import type { base } from "../../../wailsjs/go/models";
    import { MessageError } from "../../api";
    import Spotify from "../../api/spotify";

    import YandexMusic from "../../api/yandex_music";

    import Button from "../../components/button.svelte";
    import SpotifyView from "../../components/spotify_view.svelte";
    import { args } from "../../utils/store";

    let notFoundsMsg: string[] = [];

    async function importLikedTracks() {
        let notFound: base.Track[];
        notFoundsMsg = [];
        const unsub = args.subscribe((v) => {
            if (!v || !v[2]) {
                return;
            }
            notFound = v[2];
            if (!(notFound instanceof Array)) {
                return;
            }
            for (const track of notFound) {
                if (!track.artist || !track.title) {
                    return;
                }
                notFoundsMsg.push(
                    `${track.artist.join(", ")} - ${track.title}`
                );
            }
        });

        let tracks: base.Track[];
        const playlists = await YandexMusic.getPlaylists();
        if (!playlists) {
            return;
        }
        for (const playlist of playlists) {
            if (!playlist.isLikedTracks) {
                continue;
            }
            tracks = await YandexMusic.getTracks(playlist.id);
        }
        if (!tracks || tracks.length < 1) {
            await MessageError(
                "В ЯМ нет треков. Может они не загружены?"
            );
            return;
        }
        await Spotify.importLikedTracks(tracks);

        //
        unsub();
    }
</script>

<SpotifyView>
    <div class="content">
        <div class="yandex">
            <div class="title">Яндекс.Музыка</div>
            <Button on:click={async () => await importLikedTracks()}
                >Любимые треки</Button
            >
            {#if notFoundsMsg && notFoundsMsg.length > 0}
                <h2>Не найдены:</h2>
                {#each notFoundsMsg as msg}
                    <div>{msg}</div>
                {/each}
            {/if}
        </div>
    </div>
</SpotifyView>

<style lang="scss">
    .title {
        font-size: 2rem;
    }
    .yandex {
        display: flex;
        flex-direction: column;
        gap: 14px;
    }
</style>
