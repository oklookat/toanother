<script lang="ts">
    import { onMount } from "svelte";
    import type { base } from "../../../wailsjs/go/models";
    import Playlist from "../../components/playlist.svelte";
    import YmTools from "../../components/ym_tools.svelte";
    import YmView from "../../components/ym_view.svelte";
    import YandexMusic from "../../api/yandex_music";

    let playlists: base.Playlist[];

    onMount(async () => {
        await get();
    });

    async function get() {
        playlists = await YandexMusic.getPlaylists();
    }

    async function download() {
        playlists = await YandexMusic.downloadPlaylists();
    }
</script>

<YmView>
    <div class="base--list">
        <YmTools on:fetch={download} />
        {#if playlists}
            <div class="list">
                {#each playlists as playlist}
                    <Playlist {playlist} />
                {/each}
            </div>
        {/if}
    </div>
</YmView>
