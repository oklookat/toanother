<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import Button from "./button.svelte";

    const dispatch = createEventDispatcher<{ apply: MouseEvent }>();
</script>

<div class="settings">
    <div class="content">
        <slot />
    </div>
    <footer>
        <Button on:click={(e) => dispatch("apply", e.detail)}>Применить</Button>
    </footer>
</div>

<style lang="scss">
    .settings {
        height: 100%;
        display: grid;
        grid-template-columns: 1fr;
        grid-template-rows: 1fr max-content;
        .content {
            display: flex;
            flex-direction: column;
            gap: 16px;
        }
        footer {
            width: 100%;
            display: flex;
            justify-content: flex-end;
        }
    }
</style>
