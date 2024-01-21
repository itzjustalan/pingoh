<script lang="ts">
	import { goto } from '$app/navigation';
	import { authNetwork } from '$lib/network/auth.network';
	import { createMutation } from '@tanstack/svelte-query';
	import axios from 'axios';

	let email: string;
	let passw: string;
	// let password2: string;

	// const signup = createMutation<any, AxiosError>({
	const signup = createMutation({
		mutationKey: ['signup'],
		mutationFn: authNetwork.signup,
		onSuccess(data, variables, context) {
			goto("/")
        },
	});
	const handleSignup = (e: Event) => {
		e.preventDefault();
		// if (password !== password2) alert(' passwords do not match!!');
		$signup.mutate({ email, passw });
	};
</script>

<!-- <pre>{JSON.stringify(data)}</pre>
<pre>{JSON.stringify(form)}</pre> -->

<h1>signup</h1>

{#if $signup.isPending}
	loading...
{:else if $signup.isError}
	error...
	<!-- {$signup.error.response?.data} -->
	<!-- {#if $signup.error instanceof AxiosError}
        {$signup.error.response?.data}
    {/if} -->
	{#if axios.isAxiosError($signup.error)}
		{$signup.error.response?.data}
	{/if}
	<pre>{JSON.stringify($signup.error)}</pre>
{/if}

user<input type="text" name="username" bind:value={email} required />
<input type="password" name="password" bind:value={passw} required />
<button disabled={$signup.isPending} on:click={handleSignup}>summit</button>
this is - signUP
<br />
have account? <a href="/auth/signin">signIN</a>
