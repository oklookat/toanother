<script lang="ts">
    import { onMount } from "svelte";
    import type { base } from "../../../wailsjs/go/models";
    import Album from "../../components/album.svelte";
    import YmTools from "../../components/ym_tools.svelte";
    import YmView from "../../components/ym_view.svelte";
    import YandexMusic from "../../api/yandex_music";

    let albums: base.Album[];

    onMount(async () => {
        await get();
    });

    async function get() {
        albums = await YandexMusic.getAlbums();
    }

    async function download() {
        albums = await YandexMusic.downloadAlbums();
    }
</script>

<YmView>
    <div class="base--list">
        <YmTools on:fetch={download} />
        {#if albums}
            <div class="list">
                {#each albums as album}
                    <Album {album} />
                {/each}
            </div>
        {/if}
    </div>
</YmView>