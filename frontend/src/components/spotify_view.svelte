<script lang="ts">
    import BaseView from "./base_view.svelte";
    import Link from "./link.svelte";

    export let isReady = true;
</script>

{#if isReady}
    <BaseView>
        <div slot="header">
            <Link path="/spotify/settings">Настройки</Link>
            <Link path="/spotify/import">Импорт</Link>
        </div>
        <slot slot="main" />
    </BaseView>
{/if}
