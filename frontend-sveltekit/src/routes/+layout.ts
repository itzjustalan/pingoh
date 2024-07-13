import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { auth } from '$lib/stores/auth';
import { uacController } from '$lib/user.access.controller';

export const csr = true;
export const ssr = false;

// this is needed to give us force prerendering of all pages
// This can be false if you're using a fallback (i.e. SPA mode)
export const prerender = true;

export const trailingSlash = 'always';

export async function load(input) {
	// handle direct links
	if (browser) {
		const error = uacController.authorize(auth.user, input.route.id ?? '', 'get');
		if (error) {
			if (auth.isLoggedIn) {
				goto('/');
			} else {
				goto('/auth/signin');
			}
		}
	}
	return {};
}
