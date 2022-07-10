<script lang="ts">
    import type { base } from "../../wailsjs/go/models";
    import YandexMusic from "../api/yandex_music";
    import Overlay from "./overlay.svelte";

    export let playlist: base.Playlist;
    let tracks: base.Track[];

    let showTracks = false;
    async function setShowTracks(val: boolean) {
        if (val) {
            await get();
        } else {
            tracks = undefined;
        }
        showTracks = val;
    }

    async function get() {
        tracks = await YandexMusic.getTracks(playlist.id);
    }
</script>

<div class="playlist" on:click={async () => await setShowTracks(!showTracks)}>
    <div class="title">{playlist.title}</div>
    <div>треков: {playlist.trackCount}</div>
</div>

{#if showTracks && tracks}
    <Overlay onClose={async () => await setShowTracks(false)}>
        <div class="tracks">
            <div class="list">
                {#each tracks as track}
                    <div>{track.artist} - {track.title}</div>
                {/each}
            </div>
        </div>
    </Overlay>
{/if}

<style lang="scss">
    .tracks {
        background-color: var(--color-level-1);
        width: 50%;
        height: 50%;
        margin: auto;
        padding: 12px;
        overflow: auto;
        &,
        .list {
            display: flex;
            flex-direction: column;
            gap: 8px;
        }
    }
    .playlist {
        user-select: none;
        cursor: pointer;
        border: 1px solid var(--color-border);
        width: 100%;
        min-height: 32px;
        background-color: var(--color-level-1);
        padding: 12px;
        display: flex;
        flex-direction: column;
        gap: 12px;
        &:hover {
            background-color: var(--color-hover);
        }
        .title {
            font-weight: bold;
            font-size: 1.4rem;
        }
    }
</style>
