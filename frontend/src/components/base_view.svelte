<script lang="ts">
    export let isReady = true;
</script>

<div class="content">
    <header class="menu">
        <slot name="header" />
    </header>
    <div>
        {#if isReady}
            <slot name="main" />
        {/if}
    </div>
</div>

<style lang="scss">
    .content {
        height: 100%;
        display: grid;
        grid-template-columns: 1fr;
        grid-template-rows: max-content 1fr;
        .menu {
            min-height: 15%;
            width: 100%;
            background-color: var(--color-level-1);
            border-bottom: 1px solid var(--color-border);
            > :global(div) {
                width: 100%;
                height: 100%;
                display: flex;
                flex-direction: row;
                flex-wrap: wrap;
            }
        }
        > div {
            padding: 8px;
        }
    }
</style>
