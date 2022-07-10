<script lang="ts">
    import type { base } from "../../wailsjs/go/models";

    export let album: base.Album;
</script>

<div class="album">
    <div class="name">
        {album.artist} - {album.title}
    </div>
</div>

<style lang="scss">
    .album {
        background-color: var(--color-level-1);
        min-height: 32px;
        padding: 12px;
    }
</style>