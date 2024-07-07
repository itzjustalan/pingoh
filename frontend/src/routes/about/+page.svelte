<script lang="ts">
	import { authedUser } from '$lib/stores/auth';
	import { get } from 'svelte/store';

	export let data;
</script>

<h1>About</h1>
<p>{data.message}</p>
base = =<a href="/">home</a>

<div>{JSON.stringify(get(authedUser))}</div>
