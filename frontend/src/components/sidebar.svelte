<script lang="ts">
    import Link from "./link.svelte";
</script>

<aside class="sidebar">
    <Link path="/">Привет</Link>
    <Link path="/ym">Яндекс</Link>
    <Link path="/spotify">Спотифай</Link>
</aside>

<style lang="scss">
    .sidebar {
        font-weight: bold;
        width: 100%;
        background-color: var(--color-level-1);
        height: 100%;
        display: flex;
        flex-direction: column;
        border-right: 1px solid var(--color-border);
    }
</style>
