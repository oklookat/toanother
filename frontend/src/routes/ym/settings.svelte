<script lang="ts">
    import { onMount } from "svelte";
    import type { base } from "../../../wailsjs/go/models";
    import BaseSettings from "../../components/base_settings.svelte";
    import YmView from "../../components/ym_view.svelte";
    import YandexMusic from "../../api/yandex_music";

    let settings: base.YandexMusicSettings;
    let isReady = false;

    onMount(async () => {
        settings = await YandexMusic.getSettings();
        isReady = true;
    });

    async function apply() {
        await YandexMusic.applySettings(settings);
    }
</script>

<YmView bind:isReady>
    <BaseSettings on:apply={apply}>
        <div class="login">
            <input
                type="text"
                placeholder="Логин"
                bind:value={settings.login}
            />
        </div>
    </BaseSettings>
</YmView>
