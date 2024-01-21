<script lang="ts">
	import { goto } from '$app/navigation';
	import { authNetwork } from '$lib/network/auth.network';
	import { createMutation } from '@tanstack/svelte-query';

	let email: string;
	let passw: string;

	const signin = createMutation({
		mutationKey: ['signin'],
		mutationFn: authNetwork.signin,
		onSuccess(data, variables, context) {
			goto("/")
		},
	});
	const handleSignin = (e: Event) => {
		e.preventDefault();
		$signin.mutate({ email, passw });
	};
</script>

<!-- <pre>{JSON.stringify(data)}</pre>
<pre>{JSON.stringify(form)}</pre> -->

<h1>signin {$signin.status}</h1>

{#if $signin.isPending}
	loading...
{:else if $signin.isError}
	error...
	<pre>{JSON.stringify($signin.error)}</pre>
{/if}

user<input type="text" name="username" bind:value={email} required />
<input type="password" name="password" bind:value={passw} required />
<button disabled={$signin.isPending} on:click={handleSignin}>summit</button>
this is - signIN
<br />
no account? <a href="/auth/signup">signUP</a>