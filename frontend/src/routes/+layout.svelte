<script lang="ts">
	import { browser } from "$app/environment";
	import { beforeNavigate, goto } from "$app/navigation";
	import { authedUser } from "$lib/stores/auth";
	import { uacController } from "$lib/user.access.controller";
	import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
	import { get } from "svelte/store";

    // authenticate client side routing
    beforeNavigate((navigation) => {
        if (navigation.willUnload) return;
        const error = uacController.authorize(get(authedUser), navigation.to?.url.pathname ?? '', 'get')
		if (browser && error) {
			navigation.cancel()
			alert(error.message)
		}
    })

    const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
    <nav>
        <a href="/">Home</a>
        <a href="/about">About</a>
        <a href="/auth/signin">signin</a>
        {$authedUser?.role ?? '-'}
    </nav>
    <slot />
</QueryClientProvider>