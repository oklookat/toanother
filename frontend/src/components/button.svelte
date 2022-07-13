<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher<{ click: MouseEvent }>();
</script>

<div class="button" on:click={(e) => dispatch("click", e)}>
    <div>
        <slot />
    </div>
</div>

<style lang="scss">
    .button {
        user-select: none;
        min-height: 32px;
        min-width: 18px;
        width: max-content;
        cursor: pointer;
        background-color: var(--color-level-1);
        border: 1px solid var(--color-border);
        transition: background-color 40ms linear;
        display: flex;
        align-items: center;
        justify-content: center;
        &:hover {
            background-color: var(--color-hover);
        }
        > div {
            text-align: center;
            padding: 8px;
        }
    }
</style>
