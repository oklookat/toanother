<script lang="ts">
    import { createEventDispatcher } from "svelte";

    import Button from "./button.svelte";

    const dispatch = createEventDispatcher<{ fetch: MouseEvent }>();
</script>

<div class="tools">
    <Button on:click={(e) => dispatch("fetch", e.detail)}>Загрузить</Button>
</div>

<style lang="scss">
    .tools {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: 8px;
    }
</style>
