<script lang="ts">
	import { copyToClipboard } from '$lib/utils';

	export let name = '--undefined';
	const copyName = () => copyToClipboard(`var(--col-${name.split('-').at(-1)})`);
</script>

<div class="card" style="--clr: var({name});">
	{name}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="c" on:click={copyName}></div>
</div>

<style>
	.card {
		width: 250px;
		border: 1px solid balck;
		display: flex;
		justify-content: space-between;
		font-size: large;
	}
	.c {
		width: 80px;
		margin: 2px;
		color: green;
		background-color: var(--clr);
	}
	.c:hover {
		width: 70px;
	}
</style>
