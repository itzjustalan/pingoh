<script>
	// @ts-nocheck
	import { onMount } from 'svelte';
	import Card from './card.svelte';
	import { log } from '$lib/logger';

	// let cssVariables = [];
	let latteVariables = [];
	let mochaVariables = [];
	let frappeVariables = [];
	let macchiatoVariables = [];

	onMount(() => {
		const sheetId = 0;
		latteVariables = Array.from(
			Array.from(document.styleSheets).filter(
				(sheet) => sheet.href === null || sheet.href.startsWith(window.location.origin)
			)[sheetId].cssRules[0].style
		).filter((name) => name.startsWith('--latte-'));
		mochaVariables = Array.from(
			Array.from(document.styleSheets).filter(
				(sheet) => sheet.href === null || sheet.href.startsWith(window.location.origin)
			)[sheetId].cssRules[0].style
		).filter((name) => name.startsWith('--mocha-'));
		frappeVariables = Array.from(
			Array.from(document.styleSheets).filter(
				(sheet) => sheet.href === null || sheet.href.startsWith(window.location.origin)
			)[sheetId].cssRules[0].style
		).filter((name) => name.startsWith('--frappe-'));
		macchiatoVariables = Array.from(
			Array.from(document.styleSheets).filter(
				(sheet) => sheet.href === null || sheet.href.startsWith(window.location.origin)
			)[sheetId].cssRules[0].style
		).filter((name) => name.startsWith('--macchiato-'));
		log.info(latteVariables, mochaVariables);

		// const rs = getComputedStyle(document.querySelector(':root'))
		// names.forEach(name => colors[name] = rs.getPropertyValue(name))
	});
</script>

<div class="page">
	<div class="latte">
		{#each latteVariables as v (v)}
			<Card name={v} />
		{/each}
	</div>
	<div class="mocha">
		{#each mochaVariables as v (v)}
			<Card name={v} />
		{/each}
	</div>
</div>
<div class="divider"></div>
<div class="page">
	<div class="latte">
		{#each frappeVariables as v (v)}
			<Card name={v} />
		{/each}
	</div>
	<div class="mocha">
		{#each macchiatoVariables as v (v)}
			<Card name={v} />
		{/each}
	</div>
</div>

<style>
	.page {
		width: 600px;
		color: cadetblue;
		display: flex;
		justify-content: space-between;
	}
	.divider {
		width: 100vw;
		margin-top: 20px;
		margin-bottom: 20px;
		border: 1px dashed black;
	}
</style>
