<script lang="ts">
    import BaseView from "./base_view.svelte";
    import Link from "./link.svelte";

    export let isReady = true;
</script>

{#if isReady}
    <BaseView>
        <div slot="header">
            <Link path="/ym/settings">Настройки</Link>
            <Link path="/ym/playlists">Плейлисты</Link>
            <Link path="/ym/albums">Альбомы</Link>
            <Link path="/ym/artists">Артисты</Link>
        </div>
        <slot slot="main" />
    </BaseView>
{/if}
