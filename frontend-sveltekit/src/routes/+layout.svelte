<script lang="ts">
	import { browser } from '$app/environment';
	import { beforeNavigate } from '$app/navigation';
	import { auth } from '$lib/stores/auth';
	import { uacController } from '$lib/user.access.controller';
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	import { appTheme } from '$lib/stores/theme';
	import '../css/theme.css';
	import '../css/reset.css';
	import '../css/app.css';

	// authenticate client side routing
	beforeNavigate((navigation) => {
		if (navigation.willUnload) return;
		const error = uacController.authorize(auth.user, navigation.to?.url.pathname ?? '', 'get');
		if (browser && error) {
			navigation.cancel();
			alert(error.message);
		}
	});

	const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
	<nav>
		{#if $auth}
			<a href="/">Home</a>
			<a href="/about">About</a>
    {:else}
			<a href="/auth/signin">Signin</a>
		{/if}
		<a href="/dbg/col">colors</a>
		<button on:click={() => appTheme.toggle()}>toggle</button>
	</nav>
	<slot />
</QueryClientProvider>
