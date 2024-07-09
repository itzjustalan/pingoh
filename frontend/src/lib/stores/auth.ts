import { browser } from '$app/environment';
import type { UserModel } from '$lib/models/db/user.model';
import { get, writable } from 'svelte/store';

export interface AuthedUser extends UserModel {
	access_token: string;
	refresh_token: string;
}

const storeKey = 'user';

function createStore() {
	const stored = browser ? localStorage.getItem(storeKey) : null;
	const store = writable<AuthedUser | undefined>(stored === null ? undefined : JSON.parse(stored));

	return {
    get user() {
      return get(store);
    },
    get isLoggedIn() {
      return get(store) !== undefined;
    },
		subscribe: store.subscribe,
		_set: (u: AuthedUser | undefined) => store.set(u),
		_updateTokens: (tokens: Partial<AuthedUser>) =>
			store.update((v) => {
				if (!v) return;
				v.access_token = tokens.access_token ?? '';
				v.refresh_token = tokens.refresh_token ?? '';
				return v;
			})
	};
}

export const auth = createStore();
auth.subscribe((v) => {
	if (!browser) return;
	if (!v) localStorage.removeItem(storeKey);
	else localStorage.setItem(storeKey, JSON.stringify(v));
});
