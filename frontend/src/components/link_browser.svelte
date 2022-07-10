<script lang="ts">
    import { openInBrowser } from "../api";

    export let url: string;
</script>

<div
    class="link"
    on:click={() => openInBrowser(url)}
>
    <slot />
</div>
