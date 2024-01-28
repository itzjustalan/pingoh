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

<div class="title">
Signup
</div>

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

<br>
<input type="text" name="username" bind:value={email} required /> usernamE <br>
<input type="password" name="password" bind:value={passw} required /> passworD <br>
<button disabled={$signup.isPending} on:click={handleSignup}>Submit</button> this is a - signUP <br>
have an account? <a href="/auth/signin">signIN</a>

<style>
	.title {
		font-size: 4rem;
	}
</style>