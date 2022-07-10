<script lang="ts">
	import { createBodyScrollToggler } from "../utils";
	import { fade } from "svelte/transition";
	import { onDestroy, onMount } from "svelte";
	import Portal from "@oklookat/svelte-portal";

	/** can user close overlay? */
	export let closable = true;
	export let onClose: (e: MouseEvent) => void;

	let toggleScroll: () => void;
	onMount(() => {
		toggleScroll = createBodyScrollToggler();
		toggleScroll();
	});
	onDestroy(() => {
		toggleScroll();
	});
</script>

<Portal target="body">
	<div
		class="overlay"
		style={!closable ? 'z-index: 9999;' : ''}
		transition:fade={{ duration: 120 }}
		on:click|stopPropagation|self={(e) => {
			if (closable) {
				onClose(e);
			}
		}}
	>
		<slot />
	</div>
</Portal>

<style lang="scss">
	.overlay {
		background-color: rgba(0, 0, 0, 0.4);
		z-index: 10;
		max-width: 100vw;
		width: 100%;
		height: 100%;
		position: fixed;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		display: grid;
	}
</style>
