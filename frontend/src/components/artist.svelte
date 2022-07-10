<script lang="ts">
    import type { base } from "../../wailsjs/go/models";

    export let artist: base.Artist;
</script>

<div class="artist">
    <div class="name">{artist.name}</div>
</div>

<style lang="scss">
    .artist {
        background-color: var(--color-level-1);
        min-height: 32px;
        padding: 12px;
    }
</style>