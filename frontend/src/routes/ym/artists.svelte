<script lang="ts">
    import { onMount } from "svelte";
    import type { base } from "../../../wailsjs/go/models";
    import Artist from "../../components/artist.svelte";
    import YmTools from "../../components/ym_tools.svelte";

    import YmView from "../../components/ym_view.svelte";
    import YandexMusic from "../../api/yandex_music";

    let artists: base.Artist[];

    onMount(async () => {
        await get();
    });

    async function get() {
        artists = await YandexMusic.getArtists();
    }

    async function download() {
        artists = await YandexMusic.downloadArtists();
    }
</script>

<YmView>
    <div class="base--list">
        <YmTools on:fetch={download} />
        {#if artists}
            <div class="list">
                {#each artists as artist}
                    <Artist {artist} />
                {/each}
            </div>
        {/if}
    </div>
</YmView>
