<script lang="ts">
    import { push } from "svelte-spa-router";
    import SpotifyView from "../../components/spotify_view.svelte";

    push("/spotify/settings");
</script>

<SpotifyView />
