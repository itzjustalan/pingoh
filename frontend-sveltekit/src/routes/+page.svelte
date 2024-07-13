<script lang="ts">
	import backendApi from '$lib/network/apis/backend';
	import { authNetwork } from '$lib/network/auth.network';
	import { auth } from '$lib/stores/auth';

	const getVersion = async () => {
		const res = await backendApi.get('/hc');
		if (res.status !== 200) {
			throw `Error while fetching data from {url} (${res.status} ${res.statusText}).`;
		}
		return res.data;
	};
</script>

<h1>Home</h1>
Welcome {$auth?.name},<br />
{#await getVersion()}
	loading...
{:then version}
	message from Server: {version}
{:catch err}
	{err}
{/await}

<button on:click={() => authNetwork.signout()}>clear</button>
<button on:click={() => authNetwork.refresh()}>refresh</button>
