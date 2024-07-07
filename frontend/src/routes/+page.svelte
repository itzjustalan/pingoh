<script lang="ts">
	import { authNetwork } from '$lib/network/auth.network';
	import { authedUser } from '$lib/stores/auth';

	const apiUrl = (path: string) => `http://localhost:3000/api${path}`;

	const getVersion = async () => {
		const url = apiUrl('/hc');
		const res = await fetch(url);
		if (!res.ok) {
			throw `Error while fetching data from ${url} (${res.status} ${res.statusText}).`;
		}
		return await res.text();
	};
</script>

<h1>Home</h1>
{#await getVersion()}
	loading...
{:then version}
	message from Server: {version}
{:catch err}
	{err}
{/await}

<button on:click={() => authedUser.clear()}>clear</button>
<button on:click={() => authNetwork.refresh()}>refresh</button>
