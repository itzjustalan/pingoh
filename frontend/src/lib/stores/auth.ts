import { browser } from "$app/environment";
import type { UserModel } from "$lib/models/db/user.model";
import { writable } from "svelte/store";

export interface AuthedUser extends UserModel {
	access_token: string;
	refresh_token: string;
}

const storeKey = 'user';

function createStore() {
  const stored = browser ? localStorage.getItem(storeKey) : null
  const { subscribe, set, update } = writable<AuthedUser | undefined>(stored === null ? undefined : JSON.parse(stored))

  return {
    subscribe,
    set: (u: AuthedUser) => set(u),
    clear: () => update(() => undefined),
    updateTokens: (tokens: Partial<AuthedUser>) => update(v => {
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
