import { browser } from "$app/environment";
import type { UserModel } from "$lib/models/db/user.model";
import { get, writable } from "svelte/store";

export interface AuthedUser extends UserModel {
	access_token: string;
	refresh_token: string;
}

const storeKey = 'user';

function createStore() {
  const stored = browser ? localStorage.getItem(storeKey) : null
  const store = writable<AuthedUser | undefined>(stored === null ? undefined : JSON.parse(stored))

  return {
    get: () => get(store),
    subscribe: store.subscribe,
    set: (u: AuthedUser) => store.set(u),
    clear: () => store.update(() => undefined),
    authorized: () => get(store) !== undefined,
    updateTokens: (tokens: Partial<AuthedUser>) => store.update(v => {
      if (!v) return
      v.access_token = tokens.access_token ?? ''
      v.refresh_token = tokens.refresh_token ?? ''
      return v
    })
  };
}

export const authedUser = createStore()
authedUser.subscribe(v => {
  if (!browser) return
  if (!v) localStorage.removeItem(storeKey)
  else localStorage.setItem(storeKey, JSON.stringify(v))
});
