<script lang="ts">
    import { link } from "svelte-spa-router";
    import active from "svelte-spa-router/active";

    export let path: string;
    let activePath = path;

    if (path && path !== "/") {
        if (path.endsWith("/")) {
            path = path.slice(path.length - 1, path.length);
        }
        activePath = `${path}|${path}/*`;
    }
</script>

<a
    class="link--route"
    href={path}
    use:link
    use:active={{ path: activePath, className: "active" }}><slot /></a
>

<style lang="scss">
    :global(.link--route.active) {
        background-color: var(--color-hover);
    }

    .link--route {
        transition: background-color 40ms linear;
        text-align: center;
        padding: 12px;
        &:hover {
            background-color: var(--color-hover);
        }
    }
</style>
