<script lang="ts">
	import { goto } from '$app/navigation';
	import { authNetwork } from '$lib/network/auth.network';
	import { createMutation } from '@tanstack/svelte-query';

	let email: string;
	let passw: string;

	const signin = createMutation({
		mutationKey: ['signin'],
		mutationFn: authNetwork.signin,
		onSuccess(_data, _variables, _context) {
			goto('/');
		}
	});
	const handleSignin = (e: Event) => {
		e.preventDefault();
		$signin.mutate({ email, passw });
	};
</script>

<!-- <pre>{JSON.stringify(data)}</pre>
<pre>{JSON.stringify(form)}</pre> -->

<div class="page">
	<div class="title">Signin</div>

	{#if $signin.isPending}
		loading...
	{:else if $signin.isError}
		error...
		<pre>{JSON.stringify($signin.error)}</pre>
	{/if}

	<br />
	<input type="text" name="username" bind:value={email} required /> usernamE <br />
	<input type="password" name="password" bind:value={passw} required /> passworD <br />
	<button disabled={$signin.isPending} on:click={handleSignin}>Submit</button> this is - signIN
	<br />
	no account yet? <a href="/auth/signup">signUP</a>
</div>

<style>
	.page {
		margin: 0.4rem;
		/* width: 100vw;
		padding: 0 40%;
		background-color: red; */
	}
	.title {
		font-size: 4rem;
	}
	br {
		margin: 0.2rem;
	}
</style>
