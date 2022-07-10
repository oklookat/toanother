<script lang="ts">
    import { onMount } from "svelte";

    import type { base } from "../../../wailsjs/go/models";
    import Spotify from "../../api/spotify";
    import BaseSettings from "../../components/base_settings.svelte";
    import Button from "../../components/button.svelte";
    import SpotifyView from "../../components/spotify_view.svelte";

    let settings: base.SpotifySettings;
    let isReady = false;

    onMount(async () => {
        settings = await Spotify.getSettings();
        isReady = true;
    });

    async function apply() {
        await Spotify.applySettings(settings);
    }

    async function webAuth() {
        await Spotify.webAuth();
    }

    async function ping() {
        await Spotify.ping();
    }
</script>

<SpotifyView bind:isReady>
    <BaseSettings on:apply={apply}>
        <div>
            <input
                type="text"
                placeholder="Client ID"
                bind:value={settings.id}
            />
        </div>
        <div>
            <input
                type="password"
                placeholder="Client Secret"
                bind:value={settings.secret}
            />
        </div>
        <div class="tools">
            <Button on:click={webAuth}>Вход</Button>
            <Button on:click={ping}>Пинг</Button>
        </div>
    </BaseSettings>
</SpotifyView>

<style lang="scss">
    .tools {
        display: flex;
        flex-direction: row;
        gap: 12px;
    }
</style>
