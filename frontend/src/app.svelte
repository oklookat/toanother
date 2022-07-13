<script lang="ts">
  import "./style.scss";
  import Router from "svelte-spa-router";
  import Sidebar from "./components/sidebar.svelte";
  import routes from "./router";
  import Overlay from "./components/overlay.svelte";
  import { onDestroy, onMount } from "svelte";
  import { Hooks } from "./api/hooks";
  import { active, sign } from "./utils/store";

  onMount(() => {
    Hooks.Init();
  });

  onDestroy(() => {
    Hooks.Destroy();
  });
</script>

{#if $active}
  <Overlay
    closable={false}
    onClose={() => {
      return;
    }}
  >
    <div class="status">
      {@html $sign}
    </div>
  </Overlay>
{/if}

<main class="main">
  <div class="sidebar">
    <Sidebar />
  </div>

  <div class="router">
    <Router {routes} />
  </div>
</main>

<style lang="scss">
  .status {
    background-color: var(--color-level-1);
    width: 50%;
    height: 50%;
    margin: auto;
    border-radius: 6px;
    padding: 12px;
    font-size: 1.3rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
  }

  .main {
    height: 100%;
    display: grid;
    grid-template-columns: max-content 1fr;
    grid-template-rows: 1fr;

    .sidebar {
      width: 148px;
    }

    .router {
      width: 100%;
      overflow: auto;
    }
  }
</style>
